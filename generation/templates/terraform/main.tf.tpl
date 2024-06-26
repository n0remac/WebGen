provider "github" {
  token = var.github_token
  owner = var.github_owner
}

resource "github_repository" "repo" {
  name        = var.repo_name
  description = "Repository created by Terraform"
  visibility  = "public"
}
