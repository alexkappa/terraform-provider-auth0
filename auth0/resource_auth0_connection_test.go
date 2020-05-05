package auth0

import (
	"log"
	"strings"
	"testing"

	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/terraform-providers/terraform-provider-auth0/auth0/internal/random"
	"gopkg.in/auth0.v4/management"
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
					if strings.Contains(connection.GetName(), "Test") {
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
			{
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
			{
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
			{
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
			{
				Config: random.Template(testAccConnectionSMSConfig, rand),
				Check: resource.ComposeTestCheckFunc(
					random.TestCheckResourceAttr("auth0_connection.sms", "name", "Acceptance-Test-SMS-{{.random}}", rand),
					resource.TestCheckResourceAttr("auth0_connection.sms", "strategy", "sms"),
					resource.TestCheckResourceAttr("auth0_connection.sms", "options.0.twilio_sid", "ABC123"),
					resource.TestCheckResourceAttr("auth0_connection.sms", "options.0.twilio_token", "DEF456"),
					resource.TestCheckResourceAttr("auth0_connection.sms", "options.0.totp.#", "1"),
					resource.TestCheckResourceAttr("auth0_connection.sms", "options.0.totp.0.time_step", "300"),
					resource.TestCheckResourceAttr("auth0_connection.sms", "options.0.totp.0.length", "6"),
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
			{
				Config: random.Template(testAccConnectionEmailConfig, rand),
				Check: resource.ComposeTestCheckFunc(
					random.TestCheckResourceAttr("auth0_connection.email", "name", "Acceptance-Test-Email-{{.random}}", rand),
					resource.TestCheckResourceAttr("auth0_connection.email", "strategy", "email"),
					resource.TestCheckResourceAttr("auth0_connection.email", "options.0.from", "Magic Password <password@example.com>"),
					resource.TestCheckResourceAttr("auth0_connection.email", "options.0.subject", "Sign in!"),
					resource.TestCheckResourceAttr("auth0_connection.email", "options.0.totp.#", "1"),
					resource.TestCheckResourceAttr("auth0_connection.email", "options.0.totp.0.time_step", "300"),
					resource.TestCheckResourceAttr("auth0_connection.email", "options.0.totp.0.length", "6"),
				),
			},
			{
				Config: random.Template(testAccConnectionEmailConfigUpdate, rand),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("auth0_connection.email", "options.0.totp.#", "1"),
					resource.TestCheckResourceAttr("auth0_connection.email", "options.0.totp.0.time_step", "360"),
					resource.TestCheckResourceAttr("auth0_connection.email", "options.0.totp.0.length", "4"),
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

const testAccConnectionEmailConfigUpdate = `

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
			time_step = 360
			length = 4
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
			{
				Config: random.Template(testAccConnectionSalesforceConfig, rand),
				Check: resource.ComposeTestCheckFunc(
					random.TestCheckResourceAttr("auth0_connection.salesforce_community", "name", "Acceptance-Test-Salesforce-Connection-{{.random}}", rand),
					resource.TestCheckResourceAttr("auth0_connection.salesforce_community", "strategy", "salesforce-community"),
					resource.TestCheckResourceAttr("auth0_connection.salesforce_community", "options.0.community_base_url", "https://salesforce.example.com"),
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
			{
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

func TestAccConnectionApple(t *testing.T) {

	rand := random.String(6)

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"auth0": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: random.Template(testAccConnectionAppleConfig, rand),
				Check: resource.ComposeTestCheckFunc(
					random.TestCheckResourceAttr("auth0_connection.apple", "name", "Acceptance-Test-Apple-{{.random}}", rand),
					resource.TestCheckResourceAttr("auth0_connection.apple", "strategy", "apple"),
					resource.TestCheckResourceAttr("auth0_connection.apple", "options.0.client_id", "client_id"),
					resource.TestCheckResourceAttr("auth0_connection.apple", "options.0.client_secret", "-----BEGIN PRIVATE KEY-----\nMIHBAgEAMA0GCSqGSIb3DQEBAQUABIGsMIGpAgEAAiEA3+luhVHxSJ8cv3VNzQDP\nEL6BPs7FjBq4oro0MWM+QJMCAwEAAQIgWbq6/pRK4/ZXV+ZTSj7zuxsWZuK5i3ET\nfR2TCEkZR3kCEQD2ElqDr/pY5aHA++9HioY9AhEA6PIxC1c/K3gJqu+K+EsfDwIQ\nG5MS8Y7Wzv9skOOqfKnZQQIQdG24vaZZ2GwiyOD5YKiLWQIQYNtrb3j0BWsT4LI+\nN9+l1g==\n-----END PRIVATE KEY-----"),
					resource.TestCheckResourceAttr("auth0_connection.apple", "options.0.team_id", "team_id"),
					resource.TestCheckResourceAttr("auth0_connection.apple", "options.0.key_id", "key_id"),
				),
			},
			{
				Config: random.Template(testAccConnectionAppleConfigUpdate, rand),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("auth0_connection.apple", "options.0.team_id", "team_id_update"),
					resource.TestCheckResourceAttr("auth0_connection.apple", "options.0.key_id", "key_id_update"),
				),
			},
		},
	})
}

const testAccConnectionAppleConfig = `

resource "auth0_connection" "apple" {
	name = "Acceptance-Test-Apple-{{.random}}"
	is_domain_connection = false
	strategy = "apple"
	options {
		client_id = "client_id"
		client_secret = "-----BEGIN PRIVATE KEY-----\nMIHBAgEAMA0GCSqGSIb3DQEBAQUABIGsMIGpAgEAAiEA3+luhVHxSJ8cv3VNzQDP\nEL6BPs7FjBq4oro0MWM+QJMCAwEAAQIgWbq6/pRK4/ZXV+ZTSj7zuxsWZuK5i3ET\nfR2TCEkZR3kCEQD2ElqDr/pY5aHA++9HioY9AhEA6PIxC1c/K3gJqu+K+EsfDwIQ\nG5MS8Y7Wzv9skOOqfKnZQQIQdG24vaZZ2GwiyOD5YKiLWQIQYNtrb3j0BWsT4LI+\nN9+l1g==\n-----END PRIVATE KEY-----"
		team_id = "team_id"
		key_id = "key_id"
	}
}
`

const testAccConnectionAppleConfigUpdate = `

resource "auth0_connection" "apple" {
	name = "Acceptance-Test-Apple-{{.random}}"
	is_domain_connection = false
	strategy = "apple"
	options {
		client_id = "client_id"
		client_secret = "-----BEGIN PRIVATE KEY-----\nMIHBAgEAMA0GCSqGSIb3DQEBAQUABIGsMIGpAgEAAiEA3+luhVHxSJ8cv3VNzQDP\nEL6BPs7FjBq4oro0MWM+QJMCAwEAAQIgWbq6/pRK4/ZXV+ZTSj7zuxsWZuK5i3ET\nfR2TCEkZR3kCEQD2ElqDr/pY5aHA++9HioY9AhEA6PIxC1c/K3gJqu+K+EsfDwIQ\nG5MS8Y7Wzv9skOOqfKnZQQIQdG24vaZZ2GwiyOD5YKiLWQIQYNtrb3j0BWsT4LI+\nN9+l1g==\n-----END PRIVATE KEY-----"
		team_id = "team_id_update"
		key_id = "key_id_update"
	}
}
`

func TestAccConnectionGitHub(t *testing.T) {

	rand := random.String(6)

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"auth0": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: random.Template(testAccConnectionGitHubConfig, rand),
				Check: resource.ComposeTestCheckFunc(
					random.TestCheckResourceAttr("auth0_connection.github", "name", "Acceptance-Test-GitHub-{{.random}}", rand),
					resource.TestCheckResourceAttr("auth0_connection.github", "strategy", "github"),
					resource.TestCheckResourceAttr("auth0_connection.github", "options.0.client_id", "client-id"),
					resource.TestCheckResourceAttr("auth0_connection.github", "options.0.client_secret", "client-secret"),
					resource.TestCheckResourceAttr("auth0_connection.github", "options.0.scopes.#", "20"),
					resource.TestCheckResourceAttr("auth0_connection.github", "options.0.scopes.881205744", "email"),
					resource.TestCheckResourceAttr("auth0_connection.github", "options.0.scopes.4080487570", "profile"),
					resource.TestCheckResourceAttr("auth0_connection.github", "options.0.scopes.862208977", "follow"),
					resource.TestCheckResourceAttr("auth0_connection.github", "options.0.scopes.347111084", "read_repo_hook"),
					resource.TestCheckResourceAttr("auth0_connection.github", "options.0.scopes.718177942", "admin_public_key"),
					resource.TestCheckResourceAttr("auth0_connection.github", "options.0.scopes.2480957806", "write_public_key"),
					resource.TestCheckResourceAttr("auth0_connection.github", "options.0.scopes.356496889", "write_repo_hook"),
					resource.TestCheckResourceAttr("auth0_connection.github", "options.0.scopes.3006585776", "write_org"),
					resource.TestCheckResourceAttr("auth0_connection.github", "options.0.scopes.855904415", "read_user"),
					resource.TestCheckResourceAttr("auth0_connection.github", "options.0.scopes.1560560783", "admin_repo_hook"),
					resource.TestCheckResourceAttr("auth0_connection.github", "options.0.scopes.2933527251", "admin_org"),
					resource.TestCheckResourceAttr("auth0_connection.github", "options.0.scopes.1314370975", "repo"),
					resource.TestCheckResourceAttr("auth0_connection.github", "options.0.scopes.2175618052", "repo_status"),
					resource.TestCheckResourceAttr("auth0_connection.github", "options.0.scopes.188173322", "read_org"),
					resource.TestCheckResourceAttr("auth0_connection.github", "options.0.scopes.133261078", "gist"),
					resource.TestCheckResourceAttr("auth0_connection.github", "options.0.scopes.1820025999", "repo_deployment"),
					resource.TestCheckResourceAttr("auth0_connection.github", "options.0.scopes.3220703903", "public_repo"),
					resource.TestCheckResourceAttr("auth0_connection.github", "options.0.scopes.2092139895", "notifications"),
					resource.TestCheckResourceAttr("auth0_connection.github", "options.0.scopes.672436223", "delete_repo"),
					resource.TestCheckResourceAttr("auth0_connection.github", "options.0.scopes.2296398814", "read_public_key"),
				),
			},
		},
	})
}

const testAccConnectionGitHubConfig = `

resource "auth0_connection" "github" {
	name = "Acceptance-Test-GitHub-{{.random}}"
	strategy = "github"
	options {
		client_id = "client-id"
		client_secret = "client-secret"
		scopes = [ "email", "profile", "read_user", "follow", "public_repo", "repo", "repo_deployment", "repo_status", 
				   "delete_repo", "notifications", "gist", "read_repo_hook", "write_repo_hook", "admin_repo_hook",
				   "read_org", "admin_org", "read_public_key", "write_public_key", "admin_public_key", "write_org"
		]
	}
}
`

func TestAccConnectionConfiguration(t *testing.T) {

	rand := random.String(6)

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"auth0": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: random.Template(testAccConnectionConfigurationCreate, rand),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("auth0_connection.my_connection", "options.0.configuration.%", "2"),
					resource.TestCheckResourceAttr("auth0_connection.my_connection", "options.0.configuration.foo", "xxx"),
					resource.TestCheckResourceAttr("auth0_connection.my_connection", "options.0.configuration.bar", "zzz"),
				),
			},
			{
				Config: random.Template(testAccConnectionConfigurationUpdate, rand),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("auth0_connection.my_connection", "options.0.configuration.%", "3"),
					resource.TestCheckResourceAttr("auth0_connection.my_connection", "options.0.configuration.foo", "xxx"),
					resource.TestCheckResourceAttr("auth0_connection.my_connection", "options.0.configuration.bar", "yyy"),
					resource.TestCheckResourceAttr("auth0_connection.my_connection", "options.0.configuration.baz", "zzz"),
				),
			},
		},
	})
}

const testAccConnectionConfigurationCreate = `

resource "auth0_connection" "my_connection" {
	name = "Acceptance-Test-Connection-{{.random}}"
	is_domain_connection = true
	strategy = "auth0"
	options {
		configuration = {
			foo = "xxx"
			bar = "zzz"
		}
	}
}
`

const testAccConnectionConfigurationUpdate = `

resource "auth0_connection" "my_connection" {
	name = "Acceptance-Test-Connection-{{.random}}"
	is_domain_connection = true
	strategy = "auth0"
	options {
		configuration = {
			foo = "xxx"
			bar = "yyy"
			baz = "zzz"
		}
	}
}
`
