output "sshjump_public_ip" {
  value = hcloud_server.sshjump.ipv4_address
}

output "webapi_public_ip" {
  value = hcloud_server.webapi.ipv4_address
}
