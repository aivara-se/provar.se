output "public_ip" {
  value = hcloud_server.instance.ipv4_address
}
output "private_ip" {
  value = hcloud_server.instance.network[0].ip
}
output "server_id" {
  value = hcloud_server.instance.id
}
