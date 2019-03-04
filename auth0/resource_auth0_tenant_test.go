package auth0

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccTenant(t *testing.T) {

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"auth0": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: testAccTenantConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("auth0_tenant.my_tenant", "change_password.0.enabled", "true"),
					resource.TestCheckResourceAttr("auth0_tenant.my_tenant", "change_password.0.html", "<html>Change Password</html>"),
					resource.TestCheckResourceAttr("auth0_tenant.my_tenant", "guardian_mfa_page.0.enabled", "true"),
					resource.TestCheckResourceAttr("auth0_tenant.my_tenant", "guardian_mfa_page.0.html", "<html>MFA</html>"),
					resource.TestCheckResourceAttr("auth0_tenant.my_tenant", "default_audience", ""),
					resource.TestCheckResourceAttr("auth0_tenant.my_tenant", "default_directory", ""),
					resource.TestCheckResourceAttr("auth0_tenant.my_tenant", "error_page.0.html", "<html>Error Page</html>"),
					resource.TestCheckResourceAttr("auth0_tenant.my_tenant", "error_page.0.show_log_link", "false"),
					resource.TestCheckResourceAttr("auth0_tenant.my_tenant", "error_page.0.url", "https://mycompany.org/error"),
					resource.TestCheckResourceAttr("auth0_tenant.my_tenant", "friendly_name", "My Test Tenant"),
					resource.TestCheckResourceAttr("auth0_tenant.my_tenant", "picture_url", "https://mycompany.org/logo.png"),
					resource.TestCheckResourceAttr("auth0_tenant.my_tenant", "support_email", "support@mycompany.org"),
					resource.TestCheckResourceAttr("auth0_tenant.my_tenant", "support_url", "https://mycompany.org/support"),
					resource.TestCheckResourceAttr("auth0_tenant.my_tenant", "allowed_logout_urls.0", "https://mycompany.org/logoutCallback"),
					resource.TestCheckResourceAttr("auth0_tenant.my_tenant", "session_lifetime", "168"),
					resource.TestCheckResourceAttr("auth0_tenant.my_tenant", "sandbox_version", "8"),
				),
			},
		},
	})
}

const testAccTenantConfig = `
provider "auth0" {}

resource "auth0_tenant" "my_tenant" {
	change_password {
	  enabled = true
	  html = "<html>Change Password</html>"
	}
	guardian_mfa_page {
	  enabled = true
	  html = "<html>MFA</html>"
	}
	default_audience = ""
	default_directory = ""
	error_page {
	  html = "<html>Error Page</html>",
	  show_log_link = false,
	  url = "https://mycompany.org/error"
	}
	friendly_name = "My Test Tenant"
	picture_url = "https://mycompany.org/logo.png"
	support_email = "support@mycompany.org"
	support_url = "https://mycompany.org/support"
	allowed_logout_urls = [
	  "https://mycompany.org/logoutCallback"
	]
	session_lifetime = 168,
	sandbox_version = "8"
  }
`
