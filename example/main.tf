provider "auth0" {
  # "domain" = <doman>
  # "client_id" = <client-id>
  # "client_secret" = <client-secret>
}

resource "auth0_client" "my_app_client" {
  name = "My Application (Managed by Terraform)"
  description = "My Applications Long Description"
  app_type = "non_interactive"
  is_first_party = false
  oidc_conformant = false
  callbacks = [ "https://example.com/callback" ]
  allowed_origins = [ "https://example.com" ]
  web_origins = [ "https://example.com" ]
  jwt_configuration = {
    lifetime_in_seconds = 300
    secret_encoded = true
    alg = "RS256"
  }
}

resource "auth0_resource_server" "my_resource_server" {
  name = "Test Resource Server (Managed by Terraform)"
  identifier = "https://api.example.com"
  signing_alg = "RS256"
  scopes = {
    value = "create:foo"
    description = "Create foos"
  }
  scopes = {
    value = "create:bar"
    description = "Create bars"
  }
  allow_offline_access = true
  token_lifetime = 8600
  skip_consent_for_verifiable_first_party_clients = true
}
