variable "digitalocean_token" {
  description = "DigitalOcean API token"
  type        = string
  sensitive   = true
}

variable "project_name" {
  description = "Name of the DigitalOcean project"
  type        = string
}

variable "droplet_name" {
  description = "Name of the DigitalOcean droplet"
  type        = string
}

variable "region" {
  description = "Region where the droplet will be created"
  type        = string
  default     = "nyc3"
}

variable "size" {
  description = "Size of the droplet"
  type        = string
  default     = "s-1vcpu-1gb"
}

variable "image" {
  description = "Image to use for the droplet"
  type        = string
  default     = "ubuntu-20-04-x64"
}
