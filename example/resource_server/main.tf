provider "auth0" {}

resource "auth0_resource_server" "my_resource_server" {
  name        = "Example Resource Server (Managed by Terraform)"
  identifier  = "https://api.example.com"
  signing_alg = "RS256"

  scopes {
    value       = "create:foo"
    description = "Create foos"
  }

  scopes {
    value       = "create:bar"
    description = "Create bars"
  }

  allow_offline_access                            = true
  token_lifetime                                  = 8600
  skip_consent_for_verifiable_first_party_clients = true
}
