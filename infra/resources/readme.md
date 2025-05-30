## infrastructure

### Cloud Resources

This directory contains Terraform code for deploying and managing infrastructure on Hetzner Cloud. It provisions a network with a jump host, a database server, and a web API server.

## Architecture

The infrastructure consists of the following components:

- **Jump Host:** A publicly accessible server with SSH port (22) open, used for accessing the private network.
- **Database Server:** A server residing within the private network, accessible only from other servers within the same network.
- **Web API Server:** A server residing within the private network, accessible only from other servers within the same network.
- **Private Network:** A network isolating the database and web API servers, enhancing security.

## Directory Structure

```
├── modules
│   │
│   ├── instance        # Module for creating servers
│   │   ├── main.tf
│   │   ├── outputs.tf
│   │   └── variables.tf
│   │
│   └── network         # Module for creating networks
│       ├── main.tf
│       ├── outputs.tf
│       └── variables.tf
│
├── outputs.tf
├── providers.tf         # Defines the Hetzner Cloud
├── terraform.tfvars     # Has secret values (gitignored)
├── variables.tf         # Defines all global variables
└── main.tf              # Main entrypoint for resources
```

## Usage

1.  **Prerequisites:**

    - Terraform installed ([https://www.terraform.io/downloads](https://www.terraform.io/downloads))
    - Hetzner Cloud API token
    - SSH key pair for access

2.  **Configuration:**

    - Update the `ssh_public_key_path` variable to point to your public SSH key.
    - Create a `terraform.tfvars` file and set your Hetzner Cloud API token:

    ```terraform
    hcloud_token = "YOUR_HCLOUD_TOKEN"
    ```

    - Generate SSH key pairs for each system under the systems directory.

    ```shell
    # run this command in each system directory
    ssh-keygen -t rsa -b 4096 -C "admin@provar.se" -f id_rsa
    ```

    - Prepere each system as described on readme

3.  **Deployment:**

    - Initialize Terraform: `terraform init`
    - Plan the deployment: `terraform plan`
    - Apply the configuration: `terraform apply`

4.  **Access:**

    - Connect to the jump host using SSH: `ssh <user>@<jump_host_public_ip>`
    - From the jump host, you can then connect to the database and web API servers using their private IPs.

5.  **Destruction:**
    - Destroy the infrastructure: `terraform destroy`

## Modules

- **`modules/instance`:** This module creates a Hetzner Cloud instance.
- **`modules/network`:** This module creates a Hetzner Cloud network.

## Outputs

The `outputs.tf` file defines outputs for important information:

- `jump_host_public_ip`: Public IP address of the jump host.
- `jump_host_private_ip`: Private IP address of the jump host.
- `database_host_private_ip`: Private IP address of the database server.
- `webapi_host_private_ip`: Private IP address of the web API server.

## Debugging

You can use an example ssh config similar to what is given below to ssh into systems using the jump host.

```
Host sshjump.provar.se
       # FIXME: use correct ip
       HostName 1.1.1.1
       Port 22
       User root
       TCPKeepAlive yes
       ServerAliveInterval 10
       IdentityFile /path/to/sshjump/id_rsa

Host postgres.provar.se
       HostName 10.0.0.11
       Port 22
       User root
       TCPKeepAlive yes
       ServerAliveInterval 10
       IdentityFile /path/to/postgres/id_rsa
       ProxyJump sshjump.provar.se
```
