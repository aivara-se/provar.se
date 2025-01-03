terraform {
  required_providers {
    hcloud = {
      source  = "registry.terraform.io/hetznercloud/hcloud"
      version = "~> 1.47"
    }
  }
}

provider "hcloud" {
  token = var.hcloud_token
}
