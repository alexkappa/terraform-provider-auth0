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

resource "auth0_email_template" "my_email_template" {
  template = "welcome_email"
  body = "<html><body><h1>Welcome!</h1></body></html>"
  from = "welcome@example.com"
  result_url = "https://example.com/welcome"
  subject = "Welcome"
  syntax = "liquid"
  url_lifetime_in_seconds = 3600
  enabled = true

  depends_on = [ "${auth0_email.my_email_provider}" ]
}
