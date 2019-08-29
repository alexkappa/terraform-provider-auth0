provider "auth0" {}

resource "auth0_role" "my_role" {
  name        = "My Role - (Managed by Terraform)"
  description = "Role Description..."
}
