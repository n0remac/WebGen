variable "github_token" {
  description = "GitHub token with repo permissions"
  type        = string
  sensitive   = true
}

variable "github_owner" {
  description = "GitHub username or organization name"
  type        = string
}

variable "repo_name" {
  description = "Name of the repository to create"
  type        = string
}
