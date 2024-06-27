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

resource "digitalocean_project" "project" {
  name        = var.project_name
  description = "Project created by Terraform"
  purpose     = "Web Application"
  environment = "Development"
}

resource "digitalocean_droplet" "web" {
  name   = var.droplet_name
  region = var.region
  size   = var.size
  image  = var.image

  tags = ["web"]

  # Provisioning the droplet with a startup script
  user_data = <<-EOF
              #!/bin/bash
              apt-get update
              apt-get install -y nginx
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
