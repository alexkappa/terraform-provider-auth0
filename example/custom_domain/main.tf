provider "auth0" {}

resource "auth0_custom_domain" "my_custom_domain" {
  domain = "auth.example.com"
  type = "auth0_managed_certs"
  verification_method = "txt"
}
