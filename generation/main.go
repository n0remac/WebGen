package main

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"golang.org/x/crypto/ssh"

	"generation/internal/generator"
)

var projectName string = "MyApp"

func main() {
	// Define variables here
	githubOwner := "n0remac"
	gitDirectory := "../git-terraform"
	digitalOceansDirectory := "../digitalocean-terraform"

	generateApp()
	setupGitTerraform(gitDirectory, githubOwner)
	setupDigitalOceanTerraform(digitalOceansDirectory)
	createKeys()
	runDigitalOceanTerraform(digitalOceansDirectory)
	runGitTerraform(gitDirectory, githubOwner)
}

func generateApp() {
	appPath := "../MyApp"

	copyDirectory("../app", appPath)
	replaceWordInDirectory(appPath, "CodeGen", projectName)
	replaceWordInDirectory(appPath, "codegen", strings.ToLower(projectName))

	err := generator.GenerateAll("schema.yaml", appPath, projectName)
	if err != nil {
		log.Fatalf("Generation failed: %v", err)
	}
}

func setupGitTerraform(gitDirectory string, githubOwner string) {

	// Create the Terraform variables
	vars := generator.TerraformVars{
		GithubOwner: githubOwner,
		RepoName:    projectName,
	}

	createDirectory(gitDirectory)

	// Generate Terraform files
	err := generator.GenerateTerraform(vars, gitDirectory)
	if err != nil {
		log.Fatalf("Error generating Terraform files: %v", err)
	}
}

func runGitTerraform(gitDirectory string, githubOwner string) {
	// Run Terraform
	runTerraform(gitDirectory)

	// Run Git commands to push code
	runGitCommands(githubOwner, projectName)
}

func setupDigitalOceanTerraform(directory string) {
	// DigitalOcean credentials and variables
	digitaloceanToken := os.Getenv("DIGITAL_OCEANS_API")
	if digitaloceanToken == "" {
		log.Fatal("DIGITAL_OCEANS_API environment variable not set")
	}
	dropletName := projectName + "-droplet"
	region := "nyc3"
	size := "s-1vcpu-1gb"
	image := "ubuntu-20-04-x64"

	// Create the DigitalOcean Terraform variables
	vars := generator.DigitalOceanVars{
		DigitalOceanToken: digitaloceanToken,
		ProjectName:       projectName,
		DropletName:       dropletName,
		Region:            region,
		Size:              size,
		Image:             image,
	}

	createDirectory(directory)

	// Generate DigitalOcean Terraform files
	err := generator.GenerateDigitalOceanTerraform(vars, directory)
	if err != nil {
		log.Fatalf("Error generating DigitalOcean Terraform files: %v", err)
	}
}

func runDigitalOceanTerraform(directory string) {
	// Run Terraform for DigitalOcean
	runTerraform(directory)

	// Get the droplet IP address
	dropletIP, err := getTerraformOutput(directory, "droplet_ip")
	if err != nil {
		log.Fatalf("Error getting droplet IP: %v", err)
	}
	//  save droplet IP to a file
	err = os.WriteFile("../git-terraform/digitalocean_host", []byte(dropletIP), 0644)
	if err != nil {
		log.Fatalf("Error saving droplet IP: %v", err)
	}
}

func runTerraform(appPath string) {
	cmds := [][]string{
		{"terraform", "init"},
		{"terraform", "apply", "-auto-approve"},
	}

	for _, cmd := range cmds {
		runCommand(appPath, cmd[0], cmd[1:]...)
	}
}

func getTerraformOutput(appPath, outputName string) (string, error) {
	cmd := exec.Command("terraform", "output", "-raw", outputName)
	cmd.Dir = appPath
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("command terraform output -raw %s failed with %w", outputName, err)
	}
	return strings.TrimSpace(out.String()), nil
}

func runGitCommands(owner, repo string) {
	commands := [][]string{
		{"sh", "-c", `echo "# ` + repo + `" >> README.md`},
		{"git", "init"},
		{"git", "add", "."},
		{"git", "commit", "-m", "first commit"},
		{"git", "branch", "-M", "main"},
		{"git", "remote", "add", "origin", "git@github.com:" + owner + "/" + repo + ".git"},
		{"git", "push", "-u", "origin", "main"},
	}

	for _, cmd := range commands {
		runCommand("../"+projectName, cmd[0], cmd[1:]...)
	}
}

func runCommand(dir, name string, args ...string) {
	cmd := exec.Command(name, args...)
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatalf("Command %s %v failed with %v", name, args, err)
	}
}

func replaceWordInDirectory(directoryPath, oldWord, newWord string) error {
	err := filepath.Walk(directoryPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			read, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			content := string(read)
			newContent := strings.ReplaceAll(content, oldWord, newWord)
			if content != newContent {
				err = os.WriteFile(path, []byte(newContent), info.Mode())
				if err != nil {
					return err
				}
				fmt.Printf("Replaced in file: %s\n", path)
			}
		}
		return nil
	})

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return err
	}
	return nil
}

func copyDirectory(source, destination string) error {
	// Construct the cp command
	cmd := exec.Command("cp", "-r", source, destination)

	// Run the command
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return err
	}

	// Print the output of the command
	fmt.Printf("Output: %s\n", string(output))

	return nil
}

func createDirectory(appPath string) error {
	err := os.Mkdir(appPath, 0755)
	if err != nil {
		return fmt.Errorf("creating directory: %w", err)
	}

	return nil
}

func createKeys() {
	privateKey, publicKey, err := generateSSHKeyPair(2048)
	if err != nil {
		fmt.Printf("Error generating SSH key pair: %v\n", err)
		os.Exit(1)
	}

	err = saveKeyToFile(privateKey, "../git-terraform/id_rsa")
	if err != nil {
		fmt.Printf("Error saving private key: %v\n", err)
		os.Exit(1)
	}

	err = saveKeyToFile(publicKey, "../digitalocean-terraform/id_rsa.pub")
	if err != nil {
		fmt.Printf("Error saving public key: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("SSH key pair generated and saved to id_rsa and id_rsa.pub")
}

func generateSSHKeyPair(bits int) (privateKey []byte, publicKey []byte, err error) {
	// Generate a new RSA key pair
	privateKeyObj, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, nil, err
	}

	// Encode the private key to PEM format
	privateKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKeyObj),
	})

	// Create the public key
	pub, err := ssh.NewPublicKey(&privateKeyObj.PublicKey)
	if err != nil {
		return nil, nil, err
	}
	publicKeyBytes := ssh.MarshalAuthorizedKey(pub)

	return privateKeyPEM, publicKeyBytes, nil
}

func saveKeyToFile(key []byte, filename string) error {
	return os.WriteFile(filename, key, 0600)
}
