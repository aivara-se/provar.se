variable "name" {
  type = string
}

variable "server_type" {
  type    = string
  default = "cx22"
}

variable "image" {
  type    = string
  default = "ubuntu-24.04"
}

variable "location" {
  type = string
}

variable "network_id" {
  type = string
}

variable "ssh_key_id" {
  type = string
}

variable "public_network" {
  type = bool
  default = false
}

variable "ip" {
  type = string
  default = null
}
