provider "auth0" {}

resource "auth0_tenant" "tenant" {
  change_password {
    enabled = true
    html = "${file("./password_reset.html")}"
  }

  guardian_mfa_page {
    enabled = true
    html = "${file("./guardian_multifactor.html")}"
  }

  default_audience  = "<client_id>"
  default_directory = "Connection-Name"

  error_page {
      html          = "${file("./error.html")}"
      show_log_link = true
      url           = "http://mysite/errors"
  }

  friendly_name       = "Tenant Name"
  picture_url         = "http://mysite/logo.png"
  support_email       = "support@mysite"  
  support_url         = "http://mysite/support"
  allowed_logout_urls = [
      "http://mysite/logout"
  ]
  session_lifetime    = 46000
  sandbox_version     = "8"  
}
