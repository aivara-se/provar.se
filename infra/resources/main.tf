# Network that contains all resources
resource "hcloud_network" "default_network" {
  name              = "network"
  ip_range          = "10.0.0.0/8"
  delete_protection = var.is_production
}


# Default subnet for the internal network
resource "hcloud_network_subnet" "default_subnet" {
  network_id   = hcloud_network.default_network.id
  type         = "cloud"
  network_zone = "eu-central"
  ip_range     = "10.0.0.0/24"
}


# Firewall rule to allow HTTPS from all IP addresses
resource "hcloud_firewall" "public_https" {
  name = "public_https"

  rule {
    direction   = "in"
    protocol    = "tcp"
    port        = "443"
    source_ips  = ["0.0.0.0/0", "::/0"]
  }

  rule {
    direction   = "in"
    protocol    = "udp"
    port        = "443"
    source_ips  = ["0.0.0.0/0", "::/0"]
  }

  rule {
    direction   = "in"
    protocol    = "tcp"
    port        = "80"
    source_ips  = ["0.0.0.0/0", "::/0"]
  }
}


# Firewall rule to allow SSH from all IP addresses
resource "hcloud_firewall" "public_ssh" {
  name = "public_ssh"

  rule {
    direction   = "in"
    protocol    = "tcp"
    port        = "22"
    source_ips  = ["0.0.0.0/0", "::/0"]
  }
}


# Firewall rule to allow SSH from internal network
resource "hcloud_firewall" "private_ssh" {
  name = "private_ssh"

  rule {
    direction   = "in"
    protocol    = "tcp"
    port        = "22"
    source_ips  = ["10.0.0.0/24"]
  }
}


# Firewall rule to allow DB connections from internal network
resource "hcloud_firewall" "private_postgres" {
  name = "private_postgres"

  rule {
    direction   = "in"
    protocol    = "tcp"
    port        = "5432"
    source_ips  = ["10.0.0.0/24"]
  }
}


# Public IP address for the SSH jump host
resource "hcloud_primary_ip" "sshjump_ipv4" {
  name              = "sshjump_ipv4"
  type              = "ipv4"
  assignee_type     = "server"
  datacenter        = var.hcloud_datacenter
  auto_delete       = true
  delete_protection = var.is_production
}


# Public IPv6 address for the SSH jump host
resource "hcloud_primary_ip" "sshjump_ipv6" {
  name              = "sshjump_ipv6"
  type              = "ipv6"
  assignee_type     = "server"
  datacenter        = var.hcloud_datacenter
  auto_delete       = true
  delete_protection = var.is_production
}


# SSH key for accessing the sshjump instance
resource "hcloud_ssh_key" "sshjump_key" {
  name       = "sshjump_key"
  public_key = file("./systems/sshjump/id_rsa.pub")
}


# Server instance with a public IP address
# This instance will act as an SSH jump host
resource "hcloud_server" "sshjump" {
  name               = "sshjump"
  server_type        = var.hcloud_server_type
  datacenter         = var.hcloud_datacenter
  image              = var.hcloud_server_image
  ssh_keys           = [ hcloud_ssh_key.sshjump_key.id ]
  firewall_ids       = [ hcloud_firewall.public_ssh.id ]
  delete_protection  = var.is_production
  rebuild_protection = var.is_production

  user_data = templatefile("./systems/sshjump/cloud-init.yaml", {
    // ...variables
  })

  network {
    network_id = hcloud_network.default_network.id
    ip = "10.0.0.10"
  }

  public_net {
    ipv4_enabled = true
    ipv6_enabled = true
    ipv4         = hcloud_primary_ip.sshjump_ipv4.id
    ipv6         = hcloud_primary_ip.sshjump_ipv6.id
  }
}


# Public IP address for the SSH jump host
resource "hcloud_primary_ip" "postgres_ipv4" {
  name              = "postgres_ipv4"
  type              = "ipv4"
  assignee_type     = "server"
  datacenter        = var.hcloud_datacenter
  auto_delete       = true
  delete_protection = var.is_production
}


# Public IPv6 address for the SSH jump host
resource "hcloud_primary_ip" "postgres_ipv6" {
  name              = "postgres_ipv6"
  type              = "ipv6"
  assignee_type     = "server"
  datacenter        = var.hcloud_datacenter
  auto_delete       = true
  delete_protection = var.is_production
}


# SSH key for accessing the postgres instance
resource "hcloud_ssh_key" "postgres_key" {
  name       = "postgres_key"
  public_key = file("./systems/postgres/id_rsa.pub")
}


# Server instance with the postgres database
# This instance will have a volume attached for storing data
resource "hcloud_server" "postgres" {
  name               = "postgres"
  server_type        = var.hcloud_server_type
  datacenter         = var.hcloud_datacenter
  image              = var.hcloud_server_image
  ssh_keys           = [ hcloud_ssh_key.postgres_key.id ]
  firewall_ids       = [ hcloud_firewall.private_ssh.id, hcloud_firewall.private_postgres.id ]
  delete_protection  = var.is_production
  rebuild_protection = var.is_production

  user_data = templatefile("./systems/postgres/cloud-init.yaml", {
    postgres_password = var.postgres_password,
    postgres_config   = file("./systems/postgres/postgresql.conf")
  })

  network {
    network_id = hcloud_network.default_network.id
    ip = "10.0.0.11"
  }

  public_net {
    ipv4_enabled = true
    ipv6_enabled = true
    ipv4         = hcloud_primary_ip.postgres_ipv4.id
    ipv6         = hcloud_primary_ip.postgres_ipv6.id
  }
}


# Volume for storing the postgres database data
resource "hcloud_volume" "postgres_data" {
  name               = "postgres_data"
  size               = 10
  server_id          = hcloud_server.postgres.id
  automount          = true
  format             = "ext4"
  delete_protection  = var.is_production
}


# Public IPv4 address for the SSH jump host
resource "hcloud_primary_ip" "webapi_ipv4" {
  name              = "webapi_ipv4"
  type              = "ipv4"
  assignee_type     = "server"
  datacenter        = var.hcloud_datacenter
  auto_delete       = true
  delete_protection = var.is_production
}


# Public IPv4 address for the SSH jump host
resource "hcloud_primary_ip" "webapi_ipv6" {
  name              = "webapi_ipv6"
  type              = "ipv6"
  assignee_type     = "server"
  datacenter        = var.hcloud_datacenter
  auto_delete       = true
  delete_protection = var.is_production
}


# SSH key for accessing the webapi instance
resource "hcloud_ssh_key" "webapi_key" {
  name       = "webapi_key"
  public_key = file("./systems/webapi/id_rsa.pub")
}


# Server instance with the backend application
resource "hcloud_server" "webapi" {
  name               = "webapi"
  server_type        = var.hcloud_server_type
  datacenter         = var.hcloud_datacenter
  image              = var.hcloud_server_image
  ssh_keys           = [ hcloud_ssh_key.webapi_key.id ]
  firewall_ids       = [ hcloud_firewall.public_https.id, hcloud_firewall.private_ssh.id ]
  delete_protection  = var.is_production
  rebuild_protection = var.is_production

  user_data = templatefile("./systems/webapi/cloud-init.yaml", {
    // ...variables
  })

  network {
    network_id = hcloud_network.default_network.id
    ip = "10.0.0.12"
  }

  public_net {
    ipv4_enabled = true
    ipv6_enabled = true
    ipv4         = hcloud_primary_ip.webapi_ipv4.id
    ipv6         = hcloud_primary_ip.webapi_ipv6.id
  }
}
