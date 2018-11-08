provider "auth0" {}

resource "auth0_client" "my_app_client" {
  name            = "Example Application (Managed by Terraform)"
  description     = "Example Application Loooooong Description"
  app_type        = "regular_web"
  is_first_party  = true
  oidc_conformant = false
  callbacks       = ["https://example.com/callback"]
  allowed_origins = ["https://example.com"]
  web_origins     = ["https://example.com"]
  grant_types     = ["authorization_code", "http://auth0.com/oauth/grant-type/password-realm", "implicit", "password", "refresh_token"]

  jwt_configuration = {
    lifetime_in_seconds = 120
    secret_encoded      = true
    alg                 = "RS256"
  }

  custom_login_page_on = "true"

  addons = {
    firebase = {
      client_email = "wer"
      lifetime_in_seconds = 1
      private_key = "wer"
      private_key_id = "qwreerwerwe"
    }

    samlp = {
      audience = "https://example.com/saml",
      mappings = {
        email = "http://schemas.xmlsoap.org/ws/2005/05/identity/claims/emailaddress",
        name = "http://schemas.xmlsoap.org/ws/2005/05/identity/claims/name"
      },
      create_upn_claim = false,
      passthrough_claims_with_no_mapping = true,
      map_unknown_claims_as_is = false,
      map_identities = false,
      name_identifier_format = "urn:oasis:names:tc:SAML:2.0:nameid-format:persistent",
      name_identifier_probes = [
        "http://schemas.xmlsoap.org/ws/2005/05/identity/claims/emailaddress"
      ]
    }
  }
}
