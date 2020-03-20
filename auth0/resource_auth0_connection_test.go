package auth0

import (
	"log"
	"strings"
	"testing"

	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/terraform-providers/terraform-provider-auth0/auth0/internal/debug"
	"github.com/terraform-providers/terraform-provider-auth0/auth0/internal/random"
	"gopkg.in/auth0.v3/management"
)

func init() {
	resource.AddTestSweepers("auth0_connection", &resource.Sweeper{
		Name: "auth0_connection",
		F: func(_ string) error {
			api, err := Auth0()
			if err != nil {
				return err
			}
			var page int
			for {
				l, err := api.Connection.List(
					management.WithFields("id", "name"),
					management.Page(page))
				if err != nil {
					return err
				}
				for _, connection := range l.Connections {
					if strings.Contains(connection.GetName(), "Acceptance-Test") {
						log.Printf("[DEBUG] Deleting connection %v\n", connection.GetName())
						if e := api.Connection.Delete(connection.GetID()); e != nil {
							multierror.Append(err, e)
						}
					}
				}
				if err != nil {
					return err
				}
				if !l.HasNext() {
					break
				}
				page++
			}
			return nil
		},
	})
}

func TestAccConnection(t *testing.T) {

	rand := random.String(6)

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"auth0": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: random.Template(testAccConnectionConfig, rand),
				Check: resource.ComposeTestCheckFunc(
					random.TestCheckResourceAttr("auth0_connection.my_connection", "name", "Acceptance-Test-Connection-{{.random}}", rand),
					resource.TestCheckResourceAttr("auth0_connection.my_connection", "is_domain_connection", "true"),
					resource.TestCheckResourceAttr("auth0_connection.my_connection", "strategy", "auth0"),
					resource.TestCheckResourceAttr("auth0_connection.my_connection", "options.0.password_policy", "fair"),
					resource.TestCheckResourceAttr("auth0_connection.my_connection", "options.0.password_no_personal_info.0.enable", "true"),
					resource.TestCheckResourceAttr("auth0_connection.my_connection", "options.0.password_dictionary.0.enable", "true"),
					resource.TestCheckResourceAttr("auth0_connection.my_connection", "options.0.password_complexity_options.0.min_length", "6"),
					resource.TestCheckResourceAttr("auth0_connection.my_connection", "options.0.enabled_database_customization", "false"),
					resource.TestCheckResourceAttr("auth0_connection.my_connection", "options.0.brute_force_protection", "true"),
					resource.TestCheckResourceAttr("auth0_connection.my_connection", "options.0.import_mode", "false"),
					resource.TestCheckResourceAttr("auth0_connection.my_connection", "options.0.disable_signup", "false"),
					resource.TestCheckResourceAttr("auth0_connection.my_connection", "options.0.requires_username", "true"),
					resource.TestCheckResourceAttr("auth0_connection.my_connection", "options.0.custom_scripts.get_user", "myFunction"),
					resource.TestCheckResourceAttrSet("auth0_connection.my_connection", "options.0.configuration.foo"),
				),
			},
			{
				Config: random.Template(testAccConnectionConfigUpdate, rand),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("auth0_connection.my_connection", "options.0.brute_force_protection", "false"),
				),
			},
		},
	})
}

const testAccConnectionConfig = `

resource "auth0_connection" "my_connection" {
	name = "Acceptance-Test-Connection-{{.random}}"
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
		password_dictionary {
			enable = true
			dictionary = [ "password", "admin", "1234" ]
		}
		password_complexity_options {
			min_length = 6
		}
		enabled_database_customization = false
		brute_force_protection = true
		import_mode = false
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

const testAccConnectionConfigUpdate = `

resource "auth0_connection" "my_connection" {
	name = "Acceptance-Test-Connection-{{.random}}"
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
		brute_force_protection = false
		import_mode = false
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

func TestAccConnectionAD(t *testing.T) {

	rand := random.String(6)

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"auth0": Provider(),
		},
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: random.Template(testAccConnectionADConfig, rand),
				Check: resource.ComposeTestCheckFunc(
					random.TestCheckResourceAttr("auth0_connection.ad", "name", "Acceptance-Test-AD-{{.random}}", rand),
					resource.TestCheckResourceAttr("auth0_connection.ad", "strategy", "ad"),
					resource.TestCheckResourceAttr("auth0_connection.ad", "options.0.domain_aliases.#", "2"),
					resource.TestCheckResourceAttr("auth0_connection.ad", "options.0.tenant_domain", "example.com"),
					resource.TestCheckResourceAttr("auth0_connection.ad", "options.0.use_kerberos", "false"),
					resource.TestCheckResourceAttr("auth0_connection.ad", "options.0.ips.3011009788", "192.168.1.2"),
					resource.TestCheckResourceAttr("auth0_connection.ad", "options.0.ips.2555711295", "192.168.1.1"),
					resource.TestCheckResourceAttr("auth0_connection.ad", "options.0.domain_aliases.3506632655", "example.com"),
					resource.TestCheckResourceAttr("auth0_connection.ad", "options.0.domain_aliases.3154807651", "api.example.com"),
				),
			},
		},
	})
}

const testAccConnectionADConfig = `

resource "auth0_connection" "ad" {
	name = "Acceptance-Test-AD-{{.random}}"
	strategy = "ad"
	options {
		tenant_domain = "example.com"
		domain_aliases = [
			"example.com",
			"api.example.com"
		]
		ips = [ "192.168.1.1", "192.168.1.2" ]
	}
}
`

func TestAccConnectionAzureAD(t *testing.T) {

	rand := random.String(6)

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"auth0": Provider(),
		},
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: random.Template(testAccConnectionAzureADConfig, rand),
				Check: resource.ComposeTestCheckFunc(
					random.TestCheckResourceAttr("auth0_connection.azure_ad", "name", "Acceptance-Test-Azure-AD-{{.random}}", rand),
					resource.TestCheckResourceAttr("auth0_connection.azure_ad", "strategy", "waad"),
					resource.TestCheckResourceAttr("auth0_connection.azure_ad", "options.0.client_id", "123456"),
					resource.TestCheckResourceAttr("auth0_connection.azure_ad", "options.0.client_secret", "123456"),
					resource.TestCheckResourceAttr("auth0_connection.azure_ad", "options.0.tenant_domain", "example.onmicrosoft.com"),
					resource.TestCheckResourceAttr("auth0_connection.azure_ad", "options.0.domain", "example.onmicrosoft.com"),
					resource.TestCheckResourceAttr("auth0_connection.azure_ad", "options.0.domain_aliases.#", "2"),
					resource.TestCheckResourceAttr("auth0_connection.azure_ad", "options.0.domain_aliases.3506632655", "example.com"),
					resource.TestCheckResourceAttr("auth0_connection.azure_ad", "options.0.domain_aliases.3154807651", "api.example.com"),
					resource.TestCheckResourceAttr("auth0_connection.azure_ad", "options.0.scopes.#", "3"),
					resource.TestCheckResourceAttr("auth0_connection.azure_ad", "options.0.scopes.370042894", "basic_profile"),
					resource.TestCheckResourceAttr("auth0_connection.azure_ad", "options.0.scopes.1268340351", "ext_profile"),
					resource.TestCheckResourceAttr("auth0_connection.azure_ad", "options.0.scopes.541325467", "ext_groups"),
				),
			},
		},
	})
}

const testAccConnectionAzureADConfig = `

resource "auth0_connection" "azure_ad" {
	name     = "Acceptance-Test-Azure-AD-{{.random}}"
	strategy = "waad"
	options {
		client_id     = "123456"
		client_secret = "123456"
		tenant_domain = "example.onmicrosoft.com"
		domain        = "example.onmicrosoft.com"
		domain_aliases = [
			"example.com",
			"api.example.com"
		]
		use_wsfed            = false
		waad_protocol        = "openid-connect"
		waad_common_endpoint = false
		api_enable_users     = true
		scopes               = [
			"basic_profile",
			"ext_groups",
			"ext_profile"
		]
	}
}
`

func TestAccConnectionWithEnbledClients(t *testing.T) {

	rand := random.String(6)

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"auth0": Provider(),
		},
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: random.Template(testAccConnectionWithEnabledClientsConfig, rand),
				Check: resource.ComposeTestCheckFunc(
					random.TestCheckResourceAttr("auth0_connection.my_connection", "name", "Acceptance-Test-Connection-{{.random}}", rand),
					resource.TestCheckResourceAttr("auth0_connection.my_connection", "enabled_clients.#", "4"),
				),
			},
		},
	})
}

const testAccConnectionWithEnabledClientsConfig = `

resource "auth0_client" "my_client_1" {
	name = "Application - Acceptance Test - 1 - {{.random}}"
	description = "Test Applications Long Description"
	app_type = "non_interactive"
}

resource "auth0_client" "my_client_2" {
	name = "Application - Acceptance Test - 2 - {{.random}}"
	description = "Test Applications Long Description"
	app_type = "non_interactive"
}

resource "auth0_client" "my_client_3" {
	name = "Application - Acceptance Test - 3 - {{.random}}"
	description = "Test Applications Long Description"
	app_type = "non_interactive"
}

resource "auth0_client" "my_client_4" {
	name = "Application - Acceptance Test - 4 - {{.random}}"
	description = "Test Applications Long Description"
	app_type = "non_interactive"
}

resource "auth0_connection" "my_connection" {
	name = "Acceptance-Test-Connection-{{.random}}"
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

func TestAccConnectionSMS(t *testing.T) {

	rand := random.String(6)

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"auth0": Provider(),
		},
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: random.Template(testAccConnectionSMSConfig, rand),
				Check: resource.ComposeTestCheckFunc(
					random.TestCheckResourceAttr("auth0_connection.sms", "name", "Acceptance-Test-SMS-{{.random}}", rand),
					resource.TestCheckResourceAttr("auth0_connection.sms", "strategy", "sms"),
					resource.TestCheckResourceAttr("auth0_connection.sms", "options.0.twilio_sid", "ABC123"),
					resource.TestCheckResourceAttr("auth0_connection.sms", "options.0.twilio_token", "DEF456"),
					resource.TestCheckResourceAttr("auth0_connection.sms", "options.0.totp.#", "1"),
					resource.TestCheckResourceAttr("auth0_connection.sms", "options.0.totp.0.time_step", "300"),
					resource.TestCheckResourceAttr("auth0_connection.sms", "options.0.totp.0.length", "6"),
					debug.DumpAttr("auth0_connection.sms"),
				),
			},
		},
	})
}

const testAccConnectionSMSConfig = `

resource "auth0_connection" "sms" {
	name = "Acceptance-Test-SMS-{{.random}}"
	is_domain_connection = false
	strategy = "sms"

	options {
		disable_signup = false
		name = "SMS OTP"
		twilio_sid = "ABC123"
		twilio_token = "DEF456"
		from = "+12345678"
		syntax = "md_with_macros"
		template = "@@password@@"
		messaging_service_sid = "GHI789"
		brute_force_protection = true

		totp {
			time_step = 300
			length = 6
		}
	}
}
`

func TestAccConnectionEmail(t *testing.T) {

	rand := random.String(6)

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"auth0": Provider(),
		},
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: random.Template(testAccConnectionEmailConfig, rand),
				Check: resource.ComposeTestCheckFunc(
					random.TestCheckResourceAttr("auth0_connection.email", "name", "Acceptance-Test-Email-{{.random}}", rand),
					resource.TestCheckResourceAttr("auth0_connection.email", "strategy", "email"),
					resource.TestCheckResourceAttr("auth0_connection.email", "options.0.from", "Magic Password <password@example.com>"),
					resource.TestCheckResourceAttr("auth0_connection.email", "options.0.subject", "Sign in!"),
					resource.TestCheckResourceAttr("auth0_connection.email", "options.0.totp.#", "1"),
					resource.TestCheckResourceAttr("auth0_connection.email", "options.0.totp.0.time_step", "300"),
					resource.TestCheckResourceAttr("auth0_connection.email", "options.0.totp.0.length", "6"),
					debug.DumpAttr("auth0_connection.email"),
				),
			},
		},
	})
}

const testAccConnectionEmailConfig = `

resource "auth0_connection" "email" {
	name = "Acceptance-Test-Email-{{.random}}"
	is_domain_connection = false
	strategy = "email"

	options {
		disable_signup = false
		name = "Email OTP"
		from = "Magic Password <password@example.com>"
		subject = "Sign in!"
		syntax = "liquid"
		template = "<html><body><h1>Here's your password!</h1></body></html>"
		
		brute_force_protection = true

		totp {
			time_step = 300
			length = 6
		}
	}
}
`

func TestAccConnectionSalesforce(t *testing.T) {

	rand := random.String(6)

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"auth0": Provider(),
		},
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: random.Template(testAccConnectionSalesforceConfig, rand),
				Check: resource.ComposeTestCheckFunc(
					random.TestCheckResourceAttr("auth0_connection.salesforce_community", "name", "Acceptance-Test-Salesforce-Connection-{{.random}}", rand),
					resource.TestCheckResourceAttr("auth0_connection.salesforce_community", "strategy", "salesforce-community"),
					resource.TestCheckResourceAttr("auth0_connection.salesforce_community", "options.0.community_base_url", "https://salesforce.example.com"),
					debug.DumpAttr("auth0_connection.salesforce_community"),
				),
			},
		},
	})
}

const testAccConnectionSalesforceConfig = `

resource "auth0_connection" "salesforce_community" {
	name = "Acceptance-Test-Salesforce-Connection-{{.random}}"
	is_domain_connection = false
	strategy = "salesforce-community"

	options {
		client_id = false
		client_secret = "sms-connection"
		community_base_url = "https://salesforce.example.com"
	}
}
`

func TestAccConnectionGoogleOAuth2(t *testing.T) {

	rand := random.String(6)

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"auth0": Provider(),
		},
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: random.Template(testAccConnectionGoogleOAuth2Config, rand),
				Check: resource.ComposeTestCheckFunc(
					random.TestCheckResourceAttr("auth0_connection.google_oauth2", "name", "Acceptance-Test-Google-OAuth2-{{.random}}", rand),
					resource.TestCheckResourceAttr("auth0_connection.google_oauth2", "strategy", "google-oauth2"),
					resource.TestCheckResourceAttr("auth0_connection.google_oauth2", "options.0.client_id", ""),
					resource.TestCheckResourceAttr("auth0_connection.google_oauth2", "options.0.client_secret", ""),
					resource.TestCheckResourceAttr("auth0_connection.google_oauth2", "options.0.allowed_audiences.#", "2"),
					resource.TestCheckResourceAttr("auth0_connection.google_oauth2", "options.0.allowed_audiences.3506632655", "example.com"),
					resource.TestCheckResourceAttr("auth0_connection.google_oauth2", "options.0.allowed_audiences.3154807651", "api.example.com"),
					resource.TestCheckResourceAttr("auth0_connection.google_oauth2", "options.0.scopes.#", "4"),
					resource.TestCheckResourceAttr("auth0_connection.google_oauth2", "options.0.scopes.881205744", "email"),
					resource.TestCheckResourceAttr("auth0_connection.google_oauth2", "options.0.scopes.4080487570", "profile"),
					// debug.DumpAttr("auth0_connection.google_oauth2"),
				),
			},
		},
	})
}

const testAccConnectionGoogleOAuth2Config = `

resource "auth0_connection" "google_oauth2" {
	name = "Acceptance-Test-Google-OAuth2-{{.random}}"
	is_domain_connection = false
	strategy = "google-oauth2"
	options {
		client_id = ""
		client_secret = ""
		allowed_audiences = [ "example.com", "api.example.com" ]
		scopes = [ "email", "profile", "gmail", "youtube" ]
	}
}
`
