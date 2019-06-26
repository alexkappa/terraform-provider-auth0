provider "auth0" {}

resource "auth0_client" "my_client" {
  name = "Example Application - Client Grant (Managed by Terraform)"
}

resource "auth0_resource_server" "my_resource_server" {
  name       = "Example Resource Server - Client Grant (Managed by Terraform)"
  identifier = "https://api.example.com/client-grant"

  scopes {
    value       = "create:foo"
    description = "Create foos"
  }

  scopes {
    value       = "create:bar"
    description = "Create bars"
  }
}

resource "auth0_client_grant" "my_client_grant" {
  client_id = "${auth0_client.my_client.id}"
  audience  = "${auth0_resource_server.my_resource_server.identifier}"
  scope     = ["create:foo"]
}
