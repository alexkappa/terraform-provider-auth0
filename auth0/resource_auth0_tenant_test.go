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
				Config: Fixture("fixtures/tenant/create.tf"),
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
					resource.TestCheckResourceAttr("auth0_tenant.my_tenant", "session_lifetime", "720"),
					resource.TestCheckResourceAttr("auth0_tenant.my_tenant", "sandbox_version", "12"),
					resource.TestCheckResourceAttr("auth0_tenant.my_tenant", "idle_session_lifetime", "72"),
					resource.TestCheckResourceAttr("auth0_tenant.my_tenant", "enabled_locales.0", "en"),
					resource.TestCheckResourceAttr("auth0_tenant.my_tenant", "enabled_locales.1", "de"),
					resource.TestCheckResourceAttr("auth0_tenant.my_tenant", "enabled_locales.2", "fr"),
					resource.TestCheckResourceAttr("auth0_tenant.my_tenant", "flags.0.universal_login", "true"),
					resource.TestCheckResourceAttr("auth0_tenant.my_tenant", "flags.0.disable_clickjack_protection_headers", "true"),
					resource.TestCheckResourceAttr("auth0_tenant.my_tenant", "flags.0.enable_public_signup_user_exists_error", "true"),
					resource.TestCheckResourceAttr("auth0_tenant.my_tenant", "flags.0.use_scope_descriptions_for_consent", "true"),
					resource.TestCheckResourceAttr("auth0_tenant.my_tenant", "universal_login.0.colors.0.primary", "#0059d6"),
					resource.TestCheckResourceAttr("auth0_tenant.my_tenant", "universal_login.0.colors.0.page_background", "#000000"),
					resource.TestCheckResourceAttr("auth0_tenant.my_tenant", "default_redirection_uri", "https://example.com/login"),
				),
			},
			{
				Config: Fixture("fixtures/tenant/update.tf"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("auth0_tenant.my_tenant", "enabled_locales.0", "de"),
					resource.TestCheckResourceAttr("auth0_tenant.my_tenant", "enabled_locales.1", "fr"),
					resource.TestCheckResourceAttr("auth0_tenant.my_tenant", "flags.0.disable_clickjack_protection_headers", "false"),
					resource.TestCheckResourceAttr("auth0_tenant.my_tenant", "flags.0.enable_public_signup_user_exists_error", "true"),
					resource.TestCheckResourceAttr("auth0_tenant.my_tenant", "flags.0.use_scope_descriptions_for_consent", "false"),
				),
			},
		},
	})
}
