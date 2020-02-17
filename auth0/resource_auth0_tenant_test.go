package auth0

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccTenant(t *testing.T) {

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"auth0": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: testAccTenantConfigCreate,
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
					resource.TestCheckResourceAttr("auth0_tenant.my_tenant", "session_lifetime", "1080"),
					resource.TestCheckResourceAttr("auth0_tenant.my_tenant", "sandbox_version", "8"),
					resource.TestCheckResourceAttr("auth0_tenant.my_tenant", "idle_session_lifetime", "720"),
					resource.TestCheckResourceAttr("auth0_tenant.my_tenant", "flags.0.universal_login", "true"),
					resource.TestCheckResourceAttr("auth0_tenant.my_tenant", "flags.0.disable_clickjack_protection_headers", "true"),
					resource.TestCheckResourceAttr("auth0_tenant.my_tenant", "flags.0.enable_public_signup_user_exists_error", "true"),
					resource.TestCheckResourceAttr("auth0_tenant.my_tenant", "universal_login.0.colors.0.primary", "#0059d6"),
					resource.TestCheckResourceAttr("auth0_tenant.my_tenant", "universal_login.0.colors.0.page_background", "#000000"),
					resource.TestCheckResourceAttr("auth0_tenant.my_tenant", "default_redirection_uri", "https://example.com/login"),
				),
			},
			// This test case confirms issue #160 where boolean values from a
			// Bool(Map()) don't get picked up as their value is considered zero
			// (e.g. false, "").
			//
			// See: https://github.com/alexkappa/terraform-provider-auth0/issues/160
			//
			// {
			// 	Config: testAccTenantConfigUpdate,
			// 	Check: resource.ComposeTestCheckFunc(
			// 		resource.TestCheckResourceAttr("auth0_tenant.my_tenant", "flags.0.disable_clickjack_protection_headers", "false"),
			// 		resource.TestCheckResourceAttr("auth0_tenant.my_tenant", "flags.0.enable_public_signup_user_exists_error", "true"),
			// 	),
			// },
		},
	})
}

const testAccTenantConfigCreate = `
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
		html = "<html>Error Page</html>"
		show_log_link = false
		url = "https://mycompany.org/error"
	}
	friendly_name = "My Test Tenant"
	picture_url = "https://mycompany.org/logo.png"
	support_email = "support@mycompany.org"
	support_url = "https://mycompany.org/support"
	allowed_logout_urls = [
		"https://mycompany.org/logoutCallback"
	]
	session_lifetime = 1080
	sandbox_version = "8"
	idle_session_lifetime = 720
	flags {
		universal_login = true
		disable_clickjack_protection_headers = true
		enable_public_signup_user_exists_error = true
	}
	universal_login {
		colors {
			primary = "#0059d6"
			page_background = "#000000"
		}
	}
	default_redirection_uri = "https://example.com/login"
}
`

const testAccTenantConfigUpdate = `
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
		html = "<html>Error Page</html>"
		show_log_link = false
		url = "https://mycompany.org/error"
	}
	friendly_name = "My Test Tenant"
	picture_url = "https://mycompany.org/logo.png"
	support_email = "support@mycompany.org"
	support_url = "https://mycompany.org/support"
	allowed_logout_urls = [
		"https://mycompany.org/logoutCallback"
	]
	session_lifetime = 1080
	sandbox_version = "8"
	idle_session_lifetime = 720
	flags {
		universal_login = true
		disable_clickjack_protection_headers = false # <---- disable and test
		enable_public_signup_user_exists_error = true
	}
	universal_login {
		colors {
			primary = "#0059d6"
			page_background = "#000000"
		}
	}
	default_redirection_uri = "https://example.com/login"
}
`
