terraform {
  required_providers {
    digitalocean = {
      source  = "digitalocean/digitalocean"
      version = "~> 2.0"
    }
  }
}

provider "digitalocean" {
  token = var.digitalocean_token
}

resource "digitalocean_ssh_key" "default" {
  name       = "my-ssh-key"
  public_key = file("id_rsa.pub")
}

resource "digitalocean_project" "project" {
  name        = var.project_name
  description = "Project created by Terraform"
  purpose     = "Web Application"
  environment = "Development"
}

resource "digitalocean_droplet" "web" {
  name       = var.droplet_name
  region     = var.region
  size       = var.size
  image      = var.image
  ssh_keys   = [digitalocean_ssh_key.default.fingerprint]

  tags = ["web"]

  # Provisioning the droplet with a startup script
  user_data = <<-EOF
              #!/bin/bash
              set -e

              # Create a non-root user
              useradd -m -s /bin/bash app
              mkdir -p /home/app/.ssh
              cp /root/.ssh/authorized_keys /home/app/.ssh/
              chown -R app:app /home/app/.ssh
              chmod 700 /home/app/.ssh
              chmod 600 /home/app/.ssh/authorized_keys

              # Update and install necessary packages
              apt update
              apt install -y nginx

              # Install Go
              wget https://golang.org/dl/go1.21.4.linux-amd64.tar.gz
              tar -C /usr/local -xzf go1.21.4.linux-amd64.tar.gz
              echo 'export PATH=$PATH:/usr/local/go/bin' >> /etc/profile
              echo 'export PATH=$PATH:/usr/local/go/bin' >> /home/app/.profile
              chown app:app /home/app/.profile

              # Install Node.js
              curl -fsSL https://deb.nodesource.com/setup_16.x | bash -
              apt install -y nodejs

              # Verify installations
              su - app -c 'go version'
              su - app -c 'node -v'
              su - app -c 'npm -v'
              EOF
}

resource "digitalocean_project_resources" "web_project" {
  project = digitalocean_project.project.id

  resources = [
    digitalocean_droplet.web.urn
  ]
}

output "droplet_ip" {
  value = digitalocean_droplet.web.ipv4_address
}
