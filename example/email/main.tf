provider "auth0" {}

resource "auth0_email" "my_email_provider" {
  name = "ses"
  enabled = true
  default_from_address = "accounts@example.com"
  credentials {
    access_key_id = "AKIAXXXXXXXXXXXXXXXX"
    secret_access_key = "7e8c2148xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
    region = "us-east-1"
  }
}
