provider "auth0" {}

resource "auth0_client" "my_app_client" {
  name            = "Example Penis Application (Managed by Terraform)"
  description     = "Example Penis Application Loooooong Description"
  app_type        = "non_interactive"
  is_first_party  = true
  oidc_conformant = false
  callbacks       = ["https://peni.com/callback"]
  allowed_origins = ["https://peni.com"]
  web_origins     = ["https://peni.com"]

  jwt_configuration = {
    lifetime_in_seconds = 120
    secret_encoded      = true
    alg                 = "RS256"
  }
}
