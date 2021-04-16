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
	session_lifetime = 720
	sandbox_version = "12"
	idle_session_lifetime = 72
	enabled_locales = ["de", "fr"]
	flags {
		universal_login = true
		enable_public_signup_user_exists_error = true
		disable_clickjack_protection_headers = false # <---- disable and test
		use_scope_descriptions_for_consent = false   #
	}
	universal_login {
		colors {
			primary = "#0059d6"
			page_background = "#000000"
		}
	}
	default_redirection_uri = "https://example.com/login"
}