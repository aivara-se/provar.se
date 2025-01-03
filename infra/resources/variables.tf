variable "is_production" {
  type = bool
  description = "Enables deletion protection"
  default = false
}

variable "hcloud_token" {
  type        = string
  description = "Hetzner Cloud token"
  sensitive   = true
}

variable "hcloud_location" {
  type        = string
  description = "Hetzner Cloud location"
  default     = "hel1"
}

variable "hcloud_datacenter" {
  type        = string
  description = "Hetzner Cloud datacenter"
  default     = "hel1-dc2"
}

variable "hcloud_server_type" {
  type        = string
  description = "Hetzner Cloud server type"
  default     = "cx22"
}

variable "hcloud_server_image" {
  type        = string
  description = "Hetzner Cloud server image"
  default     = "debian-11"
}

variable "postgres_password" {
  type        = string
  description = "Password for the postgres user"
  sensitive   = true
}
