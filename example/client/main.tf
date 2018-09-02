provider "auth0" {}

resource "auth0_client" "my_app_client" {
  name            = "Example Application (Managed by Terraform)"
  description     = "Example Application Loooooong Description"
  app_type        = "non_interactive"
  is_first_party  = true
  oidc_conformant = true
  callbacks       = ["https://example.com/callback"]
  allowed_origins = ["https://example.com"]
  web_origins     = ["https://example.com"]
  grant_types     = ["authorization_code", "http://auth0.com/oauth/grant-type/password-realm", "implicit", "password", "refresh_token"]

  jwt_configuration = {
    lifetime_in_seconds = 120
    secret_encoded      = true
    alg                 = "RS256"
  }
}
