resource "hcloud_server" "instance" {
  name         = var.name
  server_type  = var.server_type
  image        = var.image
  location     = var.location
  ssh_keys     = [ var.ssh_key_id ]
  firewall_ids = var.public_network ? [hcloud_firewall.allow_ssh.id] : []

  network {
    network_id = var.network_id
    ip = var.ip
  }
}

resource "hcloud_firewall" "allow_ssh" {
  name = "allow-ssh"

  rule {
    direction   = "in"
    protocol    = "tcp"
    port        = "22"
    source_ips  = ["0.0.0.0/0", "::/0"]
  }
}
