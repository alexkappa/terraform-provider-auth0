provider "auth0" {}

resource "auth0_client" "my_app_client" {
  name            = "Example Application (Managed by Terraform)"
  description     = "Example Application Long Description"
  app_type        = "non_interactive"
  is_first_party  = false
  oidc_conformant = false
  callbacks       = ["https://example.com/callback"]
  allowed_origins = ["https://example.com"]
  web_origins     = ["https://example.com"]

  jwt_configuration = {
    lifetime_in_seconds = 300
    secret_encoded      = true
    alg                 = "RS256"
  }
}
