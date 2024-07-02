provider "github" {
  token = var.github_token
  owner = var.github_owner
}

resource "github_repository" "repo" {
  name        = var.repo_name
  description = "Repository created by Terraform"
  visibility  = "public"
}

resource "github_actions_secret" "ssh_private_key" {
  repository      = github_repository.repo.name
  secret_name     = "SSH_PRIVATE_KEY"
  plaintext_value = file("id_rsa")
}

resource "github_actions_secret" "digitalocean_host" {
  repository      = github_repository.repo.name
  secret_name     = "DIGITALOCEAN_HOST"
  plaintext_value = file("digitalocean_host")
}
