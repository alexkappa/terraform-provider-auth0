provider "auth0" {}

resource "auth0_tenant" "tenant" {
  change_password {
    enabled = true
    html    = "<html>Change Password</html>"
  }

  guardian_mfa_page {
    enabled = true
    html    = "<html>MFA</html>"
  }

  # default_audience  = "<client_id>"
  # default_directory = "Connection-Name"

  error_page {
    html          = "<html>Error Page</html>"
    show_log_link = true
    url           = "http://example.com/errors"
  }

  friendly_name = "Tenant Name"
  picture_url   = "http://example.com/logo.png"
  support_email = "support@example.com"
  support_url   = "http://example.com/support"
  allowed_logout_urls = [
    "http://example.com/logout"
  ]
  session_lifetime = 8760
  sandbox_version  = "8"
  enabled_locales = ["en"]
}
