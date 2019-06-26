package auth0

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccConnection(t *testing.T) {

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"auth0": Provider(),
		},
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccConnectionConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("auth0_connection.my_connection", "name", "Acceptance-Test-Connection"),
					resource.TestCheckResourceAttr("auth0_connection.my_connection", "is_domain_connection", "true"),
					resource.TestCheckResourceAttr("auth0_connection.my_connection", "strategy", "auth0"),
					resource.TestCheckResourceAttr("auth0_connection.my_connection", "options.0.password_policy", "fair"),
					resource.TestCheckResourceAttr("auth0_connection.my_connection", "options.0.password_no_personal_info.0.enable", "true"),
					resource.TestCheckResourceAttr("auth0_connection.my_connection", "options.0.enabled_database_customization", "false"),
					resource.TestCheckResourceAttr("auth0_connection.my_connection", "options.0.brute_force_protection", "true"),
					resource.TestCheckResourceAttr("auth0_connection.my_connection", "options.0.import_mode", "true"),
					resource.TestCheckResourceAttr("auth0_connection.my_connection", "options.0.disable_signup", "false"),
					resource.TestCheckResourceAttr("auth0_connection.my_connection", "options.0.requires_username", "true"),
					resource.TestCheckResourceAttr("auth0_connection.my_connection", "options.0.custom_scripts.get_user", "myFunction"),
					resource.TestCheckResourceAttrSet("auth0_connection.my_connection", "options.0.configuration.foo"),
				),
			},
		},
	})
}

const testAccConnectionConfig = `
provider "auth0" {}

resource "auth0_connection" "my_connection" {
	name = "Acceptance-Test-Connection"
	is_domain_connection = true
	strategy = "auth0"
	options {
		password_policy = "fair"
		password_history {
			enable = true
			size = 5
		}
		password_no_personal_info {
			enable = true
		}
		enabled_database_customization = false
		brute_force_protection = true
		import_mode = true
		disable_signup = false
		requires_username = true
		custom_scripts = {
			get_user = "myFunction"
		}
		configuration = {
			foo = "bar"
		}
	}
}
`

func TestAccConnectionAd(t *testing.T) {

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"auth0": Provider(),
		},
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccConnectionAdConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("auth0_connection.my_ad_connection", "name", "Acceptance-Test-Ad-Connection"),
					resource.TestCheckResourceAttr("auth0_connection.my_ad_connection", "strategy", "ad"),
				),
			},
		},
	})
}

const testAccConnectionAdConfig = `
provider "auth0" {}

resource "auth0_connection" "my_ad_connection" {
	name = "Acceptance-Test-Ad-Connection"
	strategy = "ad"
}
`

func TestAccConnectionWaad(t *testing.T) {

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"auth0": Provider(),
		},
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccConnectionWaadConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("auth0_connection.my_connection", "name", "Acceptance-Test-Waad-Connection"),
					resource.TestCheckResourceAttr("auth0_connection.my_connection", "strategy", "waad"),
				),
			},
		},
	})
}

const testAccConnectionWaadConfig = `
provider "auth0" {}

resource "auth0_connection" "my_connection" {
	name     = "Acceptance-Test-Waad-Connection"
	strategy = "waad"
	options {
		client_id     = "123456"
		client_secret = "123456"
		tenant_domain = "example.onmicrosoft.com"
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
`

func TestAccConnectionWithEnbledClients(t *testing.T) {

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"auth0": Provider(),
		},
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccConnectionWithEnabledClientsConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("auth0_connection.my_connection", "name", "Acceptance-Test-Connection"),
					resource.TestCheckResourceAttr("auth0_connection.my_connection", "enabled_clients.#", "4"),
				),
			},
		},
	})
}

const testAccConnectionWithEnabledClientsConfig = `
provider "auth0" {}

resource "auth0_client" "my_client_1" {
  name = "Application - Acceptance Test 1"
  description = "Test Applications Long Description"
  app_type = "non_interactive"
}

resource "auth0_client" "my_client_2" {
  name = "Application - Acceptance Test 2"
  description = "Test Applications Long Description"
  app_type = "non_interactive"
}

resource "auth0_client" "my_client_3" {
  name = "Application - Acceptance Test 3"
  description = "Test Applications Long Description"
  app_type = "non_interactive"
}

resource "auth0_client" "my_client_4" {
  name = "Application - Acceptance Test 4"
  description = "Test Applications Long Description"
  app_type = "non_interactive"
}

resource "auth0_connection" "my_connection" {
	name = "Acceptance-Test-Connection"
	is_domain_connection = true
	strategy = "auth0"
	enabled_clients = [
    "${auth0_client.my_client_1.id}",
    "${auth0_client.my_client_2.id}",
    "${auth0_client.my_client_3.id}",
    "${auth0_client.my_client_4.id}",
  ]

}
`

func testTwilioConnection(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"auth0": Provider(),
		},
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testTwilioConnectionConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("auth0_connection.sms_connection", "name", "Acceptance-Test-Connection"),
					resource.TestCheckResourceAttr("auth0_connection.sms_connection", "strategy", "sms"),
				),
			},
		},
	})
}

const testTwilioConnectionConfig = `
resource "auth0_connection" "sms_connection" {
	name = "sms-connection"
	is_domain_connection = false
	strategy = "sms"
	
	options = {
		disable_signup = false
		name = "sms-connection"
		twilio_sid = "ABC123"
		twilio_token = "DEF456"
		from = "+12345678"
		syntax = "md_with_macros"
		template = "@@password@@"
		messaging_service_sid = "GHI789"
		brute_force_protection = true
		
		totp = {
			time_step = 300
			length = 6
		}
	}
}
`
