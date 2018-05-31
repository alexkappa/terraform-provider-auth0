provider "auth0" {}

resource "auth0_connection" "my_connection" {
  name     = "my-waad-connection"
  strategy = "waad"

  options = {
    client_id     = "1234"
    client_secret = "1234"
    tenant_domain = "exmaple.com"

    domain_aliases = [
      "example.io",
    ]

    use_wsfed            = false
    waad_protocol        = "openid-connect"
    waad_common_endpoint = false

    app_domain       = "my-auth0-app.eu.auth0.com"
    api_enable_users = true
    basic_profile    = true
    ext_groups       = true
    ext_profile      = true
  }
}
