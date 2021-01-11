package auth0

import (
	"log"
	"reflect"
	"strings"
	"testing"

	"github.com/alexkappa/terraform-provider-auth0/auth0/internal/random"
	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"gopkg.in/auth0.v5/management"
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
					management.IncludeFields("id", "name"),
					management.Page(page))
				if err != nil {
					return err
				}
				for _, connection := range l.Connections {
					log.Printf("[DEBUG] ➝ %s", connection.GetName())
					if strings.Contains(connection.GetName(), "Test") {
						if e := api.Connection.Delete(connection.GetID()); e != nil {
							multierror.Append(err, e)
						}
						log.Printf("[DEBUG] ✗ %s", connection.GetName())
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
					resource.TestCheckResourceAttr("auth0_connection.my_connection", "options.0.validation.0.username.0.min", "10"),
					resource.TestCheckResourceAttr("auth0_connection.my_connection", "options.0.validation.0.username.0.max", "40"),
					resource.TestCheckResourceAttr("auth0_connection.my_connection", "options.0.custom_scripts.get_user", "myFunction"),
					resource.TestCheckResourceAttr("auth0_connection.my_connection", "options.0.mfa.0.active", "true"),
					resource.TestCheckResourceAttr("auth0_connection.my_connection", "options.0.mfa.0.return_enroll_settings", "true"),
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
		validation {
			username {
				min = 10
				max = 40
			}
		}
		enabled_database_customization = false
		brute_force_protection = true
		import_mode = false
		requires_username = true
		disable_signup = false
		custom_scripts = {
			get_user = "myFunction"
		}
		configuration = {
			foo = "bar"
		}
	  mfa {
      active                 = true
      return_enroll_settings = true
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
		mfa {
			active                 = true
			return_enroll_settings = true
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

func TestAccConnectionOIDC(t *testing.T) {

	rand := random.String(6)

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"auth0": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: random.Template(testAccConnectionOIDCConfig, rand),
				Check: resource.ComposeTestCheckFunc(
					random.TestCheckResourceAttr("auth0_connection.oidc", "name", "Acceptance-Test-OIDC-{{.random}}", rand),
					resource.TestCheckResourceAttr("auth0_connection.oidc", "strategy", "oidc"),
					resource.TestCheckResourceAttr("auth0_connection.oidc", "options.0.client_id", "123456"),
					resource.TestCheckResourceAttr("auth0_connection.oidc", "options.0.client_secret", "123456"),
					resource.TestCheckResourceAttr("auth0_connection.oidc", "options.0.domain_aliases.#", "2"),
					resource.TestCheckResourceAttr("auth0_connection.oidc", "options.0.domain_aliases.3506632655", "example.com"),
					resource.TestCheckResourceAttr("auth0_connection.oidc", "options.0.domain_aliases.3154807651", "api.example.com"),
					resource.TestCheckResourceAttr("auth0_connection.oidc", "options.0.type", "back_channel"),
					resource.TestCheckResourceAttr("auth0_connection.oidc", "options.0.issuer", "https://api.login.yahoo.com"),
					resource.TestCheckResourceAttr("auth0_connection.oidc", "options.0.jwks_uri", "https://api.login.yahoo.com/openid/v1/certs"),
					resource.TestCheckResourceAttr("auth0_connection.oidc", "options.0.discovery_url", "https://api.login.yahoo.com/.well-known/openid-configuration"),
					resource.TestCheckResourceAttr("auth0_connection.oidc", "options.0.token_endpoint", "https://api.login.yahoo.com/oauth2/get_token"),
					resource.TestCheckResourceAttr("auth0_connection.oidc", "options.0.userinfo_endpoint", "https://api.login.yahoo.com/openid/v1/userinfo"),
					resource.TestCheckResourceAttr("auth0_connection.oidc", "options.0.authorization_endpoint", "https://api.login.yahoo.com/oauth2/request_auth"),
					resource.TestCheckResourceAttr("auth0_connection.oidc", "options.0.scopes.#", "3"),
					resource.TestCheckResourceAttr("auth0_connection.oidc", "options.0.scopes.2517049750", "openid"),
					resource.TestCheckResourceAttr("auth0_connection.oidc", "options.0.scopes.4080487570", "profile"),
					resource.TestCheckResourceAttr("auth0_connection.oidc", "options.0.scopes.881205744", "email"),
				),
			},
			{
				Config: random.Template(testAccConnectionOIDCConfigUpdate, rand),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("auth0_connection.oidc", "options.0.client_id", "1234567"),
					resource.TestCheckResourceAttr("auth0_connection.oidc", "options.0.client_secret", "1234567"),
					resource.TestCheckResourceAttr("auth0_connection.oidc", "options.0.domain_aliases.#", "1"),
					resource.TestCheckResourceAttr("auth0_connection.oidc", "options.0.domain_aliases.3506632655", "example.com"),
					resource.TestCheckResourceAttr("auth0_connection.oidc", "options.0.type", "front_channel"),
					resource.TestCheckResourceAttr("auth0_connection.oidc", "options.0.issuer", "https://www.paypalobjects.com"),
					resource.TestCheckResourceAttr("auth0_connection.oidc", "options.0.jwks_uri", "https://api.paypal.com/v1/oauth2/certs"),
					resource.TestCheckResourceAttr("auth0_connection.oidc", "options.0.discovery_url", "https://www.paypalobjects.com/.well-known/openid-configuration"),
					resource.TestCheckResourceAttr("auth0_connection.oidc", "options.0.token_endpoint", "https://api.paypal.com/v1/oauth2/token"),
					resource.TestCheckResourceAttr("auth0_connection.oidc", "options.0.userinfo_endpoint", "https://api.paypal.com/v1/oauth2/token/userinfo"),
					resource.TestCheckResourceAttr("auth0_connection.oidc", "options.0.authorization_endpoint", "https://www.paypal.com/signin/authorize"),
					resource.TestCheckResourceAttr("auth0_connection.oidc", "options.0.scopes.#", "2"),
					resource.TestCheckResourceAttr("auth0_connection.oidc", "options.0.scopes.2517049750", "openid"),
					resource.TestCheckResourceAttr("auth0_connection.oidc", "options.0.scopes.881205744", "email"),
				),
			},
		},
	})
}

const testAccConnectionOIDCConfig = `

resource "auth0_connection" "oidc" {
	name     = "Acceptance-Test-OIDC-{{.random}}"
	strategy = "oidc"
	options {
		client_id     = "123456"
		client_secret = "123456"
		domain_aliases = [
			"example.com",
			"api.example.com"
		]
		type                   = "back_channel"
		issuer                 = "https://api.login.yahoo.com"
		jwks_uri               = "https://api.login.yahoo.com/openid/v1/certs"
		discovery_url          = "https://api.login.yahoo.com/.well-known/openid-configuration"
		token_endpoint         = "https://api.login.yahoo.com/oauth2/get_token"
		userinfo_endpoint      = "https://api.login.yahoo.com/openid/v1/userinfo"
		authorization_endpoint = "https://api.login.yahoo.com/oauth2/request_auth"
		scopes = [ "openid", "email", "profile" ]
	}
}
`

const testAccConnectionOIDCConfigUpdate = `

resource "auth0_connection" "oidc" {
	name     = "Acceptance-Test-OIDC-{{.random}}"
	strategy = "oidc"
	options {
		client_id     = "1234567"
		client_secret = "1234567"
		domain_aliases = [
			"example.com"
		]
		type                   = "front_channel"
		issuer                 = "https://www.paypalobjects.com"
		jwks_uri               = "https://api.paypal.com/v1/oauth2/certs"
		discovery_url          = "https://www.paypalobjects.com/.well-known/openid-configuration"
		token_endpoint         = "https://api.paypal.com/v1/oauth2/token"
		userinfo_endpoint      = "https://api.paypal.com/v1/oauth2/token/userinfo"
		authorization_endpoint = "https://www.paypal.com/signin/authorize"
		scopes = [ "openid", "email" ]
	}
}
`

func TestAccConnectionOAuth2(t *testing.T) {

	rand := random.String(6)

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"auth0": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: random.Template(testAccConnectionOAuth2Config, rand),
				Check: resource.ComposeTestCheckFunc(
					random.TestCheckResourceAttr("auth0_connection.oauth2", "name", "Acceptance-Test-OAuth2-{{.random}}", rand),
					resource.TestCheckResourceAttr("auth0_connection.oauth2", "strategy", "oauth2"),
					resource.TestCheckResourceAttr("auth0_connection.oauth2", "options.0.client_id", "123456"),
					resource.TestCheckResourceAttr("auth0_connection.oauth2", "options.0.client_secret", "123456"),
					resource.TestCheckResourceAttr("auth0_connection.oauth2", "options.0.token_endpoint", "https://api.login.yahoo.com/oauth2/get_token"),
					resource.TestCheckResourceAttr("auth0_connection.oauth2", "options.0.authorization_endpoint", "https://api.login.yahoo.com/oauth2/request_auth"),
					resource.TestCheckResourceAttr("auth0_connection.oauth2", "options.0.scopes.#", "3"),
					resource.TestCheckResourceAttr("auth0_connection.oauth2", "options.0.scopes.2517049750", "openid"),
					resource.TestCheckResourceAttr("auth0_connection.oauth2", "options.0.scopes.4080487570", "profile"),
					resource.TestCheckResourceAttr("auth0_connection.oauth2", "options.0.scopes.881205744", "email"),
					resource.TestCheckResourceAttr("auth0_connection.oauth2", "options.0.scripts.fetchUserProfile", "function( { return callback(null) }"),
				),
			},
			{
				Config: random.Template(testAccConnectionOAuth2ConfigUpdate, rand),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("auth0_connection.oauth2", "options.0.client_id", "1234567"),
					resource.TestCheckResourceAttr("auth0_connection.oauth2", "options.0.client_secret", "1234567"),
					resource.TestCheckResourceAttr("auth0_connection.oauth2", "options.0.token_endpoint", "https://api.paypal.com/v1/oauth2/token"),
					resource.TestCheckResourceAttr("auth0_connection.oauth2", "options.0.authorization_endpoint", "https://www.paypal.com/signin/authorize"),
					resource.TestCheckResourceAttr("auth0_connection.oauth2", "options.0.scopes.#", "2"),
					resource.TestCheckResourceAttr("auth0_connection.oauth2", "options.0.scopes.2517049750", "openid"),
					resource.TestCheckResourceAttr("auth0_connection.oauth2", "options.0.scopes.881205744", "email"),
					resource.TestCheckResourceAttr("auth0_connection.oauth2", "options.0.scripts.fetchUserProfile", "function( { return callback(null) }"),
				),
			},
		},
	})
}

const testAccConnectionOAuth2Config = `

resource "auth0_connection" "oauth2" {
	name     = "Acceptance-Test-OAuth2-{{.random}}"
	strategy = "oauth2"
	is_domain_connection = false
	options {
		client_id     = "123456"
		client_secret = "123456"
		token_endpoint         = "https://api.login.yahoo.com/oauth2/get_token"
		authorization_endpoint = "https://api.login.yahoo.com/oauth2/request_auth"
		scopes = [ "openid", "email", "profile" ]
		scripts = {
			fetchUserProfile= "function( { return callback(null) }"
		}
	}
}
`

const testAccConnectionOAuth2ConfigUpdate = `

resource "auth0_connection" "oauth2" {
	name     = "Acceptance-Test-OAuth2-{{.random}}"
	strategy = "oauth2"
	is_domain_connection = false
	options {
		client_id     = "1234567"
		client_secret = "1234567"
		token_endpoint         = "https://api.paypal.com/v1/oauth2/token"
		authorization_endpoint = "https://www.paypal.com/signin/authorize"
		scopes = [ "openid", "email" ]
		scripts = {
			fetchUserProfile= "function( { return callback(null) }"
		}
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
		client_id = "client-id"
		client_secret = "client-secret"
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

func TestAccConnectionFacebook(t *testing.T) {

	rand := random.String(6)

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"auth0": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: random.Template(testAccConnectionFacebookConfig, rand),
				Check: resource.ComposeTestCheckFunc(
					random.TestCheckResourceAttr("auth0_connection.facebook", "name", "Acceptance-Test-Facebook-{{.random}}", rand),
					resource.TestCheckResourceAttr("auth0_connection.facebook", "strategy", "facebook"),
					resource.TestCheckResourceAttr("auth0_connection.facebook", "options.0.client_id", "client_id"),
					resource.TestCheckResourceAttr("auth0_connection.facebook", "options.0.client_secret", "client_secret"),
					resource.TestCheckResourceAttr("auth0_connection.facebook", "options.0.scopes.#", "4"),
					resource.TestCheckResourceAttr("auth0_connection.facebook", "options.0.scopes.3590537325", "public_profile"),
					resource.TestCheckResourceAttr("auth0_connection.facebook", "options.0.scopes.881205744", "email"),
				),
			},
			{
				Config: random.Template(testAccConnectionFacebookConfigUpdate, rand),
				Check: resource.ComposeTestCheckFunc(
					random.TestCheckResourceAttr("auth0_connection.facebook", "name", "Acceptance-Test-Facebook-{{.random}}", rand),
					resource.TestCheckResourceAttr("auth0_connection.facebook", "strategy", "facebook"),
					resource.TestCheckResourceAttr("auth0_connection.facebook", "options.0.client_id", "client_id_update"),
					resource.TestCheckResourceAttr("auth0_connection.facebook", "options.0.client_secret", "client_secret_update"),
					resource.TestCheckResourceAttr("auth0_connection.facebook", "options.0.scopes.#", "2"),
					resource.TestCheckResourceAttr("auth0_connection.facebook", "options.0.scopes.3590537325", "public_profile"),
					resource.TestCheckResourceAttr("auth0_connection.facebook", "options.0.scopes.881205744", "email"),
				),
			},
		},
	})
}

const testAccConnectionFacebookConfig = `

resource "auth0_connection" "facebook" {
	name = "Acceptance-Test-Facebook-{{.random}}"
	is_domain_connection = false
	strategy = "facebook"
	options {
		client_id = "client_id"
		client_secret = "client_secret"
		scopes = [ "public_profile", "email", "groups_access_member_info", "user_birthday" ]
	}
}
`

const testAccConnectionFacebookConfigUpdate = `

resource "auth0_connection" "facebook" {
	name = "Acceptance-Test-Facebook-{{.random}}"
	is_domain_connection = false
	strategy = "facebook"
	options {
		client_id = "client_id_update"
		client_secret = "client_secret_update"
		scopes = [ "public_profile", "email" ]
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
					resource.TestCheckResourceAttr("auth0_connection.apple", "options.0.scopes.#", "2"),
					resource.TestCheckResourceAttr("auth0_connection.apple", "options.0.scopes.2318696674", "name"),
					resource.TestCheckResourceAttr("auth0_connection.apple", "options.0.scopes.881205744", "email"),
				),
			},
			{
				Config: random.Template(testAccConnectionAppleConfigUpdate, rand),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("auth0_connection.apple", "options.0.team_id", "team_id_update"),
					resource.TestCheckResourceAttr("auth0_connection.apple", "options.0.key_id", "key_id_update"),
					resource.TestCheckResourceAttr("auth0_connection.apple", "options.0.scopes.#", "1"),
					resource.TestCheckResourceAttr("auth0_connection.apple", "options.0.scopes.881205744", "email"),
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
		scopes = ["email", "name"]
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
		scopes = ["email"]
	}
}
`

func TestAccConnectionLinkedin(t *testing.T) {

	rand := random.String(6)

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"auth0": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: random.Template(testAccConnectionLinkedinConfig, rand),
				Check: resource.ComposeTestCheckFunc(
					random.TestCheckResourceAttr("auth0_connection.linkedin", "name", "Acceptance-Test-Linkedin-{{.random}}", rand),
					resource.TestCheckResourceAttr("auth0_connection.linkedin", "strategy", "linkedin"),
					resource.TestCheckResourceAttr("auth0_connection.linkedin", "options.0.client_id", "client_id"),
					resource.TestCheckResourceAttr("auth0_connection.linkedin", "options.0.client_secret", "client_secret"),
					resource.TestCheckResourceAttr("auth0_connection.linkedin", "options.0.strategy_version", "2"),
					resource.TestCheckResourceAttr("auth0_connection.linkedin", "options.0.scopes.#", "3"),
					resource.TestCheckResourceAttr("auth0_connection.linkedin", "options.0.scopes.370042894", "basic_profile"),
					resource.TestCheckResourceAttr("auth0_connection.linkedin", "options.0.scopes.881205744", "email"),
				),
			},
			{
				Config: random.Template(testAccConnectionLinkedinConfigUpdate, rand),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("auth0_connection.linkedin", "options.0.client_id", "client_id_update"),
					resource.TestCheckResourceAttr("auth0_connection.linkedin", "options.0.client_secret", "client_secret_update"),
					resource.TestCheckResourceAttr("auth0_connection.linkedin", "options.0.scopes.370042894", "basic_profile"),
					resource.TestCheckResourceAttr("auth0_connection.linkedin", "options.0.scopes.#", "2"),
				),
			},
		},
	})
}

const testAccConnectionLinkedinConfig = `

resource "auth0_connection" "linkedin" {
	name = "Acceptance-Test-Linkedin-{{.random}}"
	is_domain_connection = false
	strategy = "linkedin"
	options {
		client_id = "client_id"
		client_secret = "client_secret"
		strategy_version = 2
		scopes = [ "basic_profile", "profile", "email" ]
	}
}
`

const testAccConnectionLinkedinConfigUpdate = `

resource "auth0_connection" "linkedin" {
	name = "Acceptance-Test-Linkedin-{{.random}}"
	is_domain_connection = false
	strategy = "linkedin"
	options {
		client_id = "client_id_update"
		client_secret = "client_secret_update"
		strategy_version = 2
		scopes = [ "basic_profile", "profile" ]
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

func TestConnectionInstanceStateUpgradeV0(t *testing.T) {

	for _, tt := range []struct {
		name            string
		version         interface{}
		versionExpected int
	}{
		{
			name:            "Empty",
			version:         "",
			versionExpected: 0,
		},
		{
			name:            "Zero",
			version:         "0",
			versionExpected: 0,
		},
		{
			name:            "NonZero",
			version:         "123",
			versionExpected: 123,
		},
		{
			name:            "Invalid",
			version:         "foo",
			versionExpected: 0,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {

			state := map[string]interface{}{
				"options": []interface{}{
					map[string]interface{}{"strategy_version": tt.version},
				},
			}

			actual, err := connectionSchemaUpgradeV0(state, nil)
			if err != nil {
				t.Fatalf("error migrating state: %s", err)
			}

			expected := map[string]interface{}{
				"options": []interface{}{
					map[string]interface{}{"strategy_version": tt.versionExpected},
				},
			}

			if !reflect.DeepEqual(expected, actual) {
				t.Fatalf("\n\nexpected:\n\n%#v\n\ngot:\n\n%#v\n\n", expected, actual)
			}
		})
	}
}

func TestConnectionInstanceStateUpgradeV1(t *testing.T) {

	for _, tt := range []struct {
		name               string
		validation         map[string]string
		validationExpected []map[string][]interface{}
	}{
		{
			name: "Only Min",
			validation: map[string]string{
				"min": "5",
			},
			validationExpected: []map[string][]interface{}{
				{
					"username": []interface{}{
						map[string]string{
							"min": "5",
						},
					},
				},
			},
		},
		{
			name: "Min and Max",
			validation: map[string]string{
				"min": "5",
				"max": "10",
			},
			validationExpected: []map[string][]interface{}{
				{
					"username": []interface{}{
						map[string]string{
							"min": "5",
							"max": "10",
						},
					},
				},
			},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {

			state := map[string]interface{}{
				"options": []interface{}{
					map[string]interface{}{"validation": tt.validation},
				},
			}

			actual, err := connectionSchemaUpgradeV1(state, nil)
			if err != nil {
				t.Fatalf("error migrating state: %s", err)
			}

			expected := map[string]interface{}{
				"options": []interface{}{
					map[string]interface{}{"validation": tt.validationExpected},
				},
			}

			if !reflect.DeepEqual(expected, actual) {
				t.Fatalf("\n\nexpected:\n\n%#v\n\ngot:\n\n%#v\n\n", expected, actual)
			}
		})
	}
}

func TestAccConnectionSAML(t *testing.T) {
	rand := random.String(6)

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"auth0": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: random.Template(testConnectionSAMLConfigCreate, rand),
				Check: resource.ComposeTestCheckFunc(
					random.TestCheckResourceAttr("auth0_connection.my_connection", "name", "Acceptance-Test-SAML-{{.random}}", rand),
					resource.TestCheckResourceAttr("auth0_connection.my_connection", "strategy", "samlp"),
				),
			},
			{
				Config: random.Template(testConnectionSAMLConfigUpdate, rand),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("auth0_connection.my_connection", "options.0.idp_initiated.0.client_authorize_query", "type=code&timeout=60"),
					resource.TestCheckResourceAttr("auth0_connection.my_connection", "options.0.sign_out_endpoint", ""),
				),
			},
		},
	})
}

const testConnectionSAMLConfigCreate = `
resource "auth0_connection" "my_connection" {
	name = "Acceptance-Test-SAML-{{.random}}"
	strategy = "samlp"
	options {
		signing_cert = <<EOF
-----BEGIN CERTIFICATE-----
MIID6TCCA1ICAQEwDQYJKoZIhvcNAQEFBQAwgYsxCzAJBgNVBAYTAlVTMRMwEQYD
VQQIEwpDYWxpZm9ybmlhMRYwFAYDVQQHEw1TYW4gRnJhbmNpc2NvMRQwEgYDVQQK
EwtHb29nbGUgSW5jLjEMMAoGA1UECxMDRW5nMQwwCgYDVQQDEwNhZ2wxHTAbBgkq
hkiG9w0BCQEWDmFnbEBnb29nbGUuY29tMB4XDTA5MDkwOTIyMDU0M1oXDTEwMDkw
OTIyMDU0M1owajELMAkGA1UEBhMCQVUxEzARBgNVBAgTClNvbWUtU3RhdGUxITAf
BgNVBAoTGEludGVybmV0IFdpZGdpdHMgUHR5IEx0ZDEjMCEGA1UEAxMaZXVyb3Bh
LnNmby5jb3JwLmdvb2dsZS5jb20wggIiMA0GCSqGSIb3DQEBAQUAA4ICDwAwggIK
AoICAQC6pgYt7/EibBDumASF+S0qvqdL/f+nouJw2T1Qc8GmXF/iiUcrsgzh/Fd8
pDhz/T96Qg9IyR4ztuc2MXrmPra+zAuSf5bevFReSqvpIt8Duv0HbDbcqs/XKPfB
uMDe+of7a9GCywvAZ4ZUJcp0thqD9fKTTjUWOBzHY1uNE4RitrhmJCrbBGXbJ249
bvgmb7jgdInH2PU7PT55hujvOoIsQW2osXBFRur4pF1wmVh4W4lTLD6pjfIMUcML
ICHEXEN73PDic8KS3EtNYCwoIld+tpIBjE1QOb1KOyuJBNW6Esw9ALZn7stWdYcE
qAwvv20egN2tEXqj7Q4/1ccyPZc3PQgC3FJ8Be2mtllM+80qf4dAaQ/fWvCtOrQ5
pnfe9juQvCo8Y0VGlFcrSys/MzSg9LJ/24jZVgzQved/Qupsp89wVidwIzjt+WdS
fyWfH0/v1aQLvu5cMYuW//C0W2nlYziL5blETntM8My2ybNARy3ICHxCBv2RNtPI
WQVm+E9/W5rwh2IJR4DHn2LHwUVmT/hHNTdBLl5Uhwr4Wc7JhE7AVqb14pVNz1lr
5jxsp//ncIwftb7mZQ3DF03Yna+jJhpzx8CQoeLT6aQCHyzmH68MrHHT4MALPyUs
Pomjn71GNTtDeWAXibjCgdL6iHACCF6Htbl0zGlG0OAK+bdn0QIDAQABMA0GCSqG
SIb3DQEBBQUAA4GBAOKnQDtqBV24vVqvesL5dnmyFpFPXBn3WdFfwD6DzEb21UVG
5krmJiu+ViipORJPGMkgoL6BjU21XI95VQbun5P8vvg8Z+FnFsvRFY3e1CCzAVQY
ZsUkLw2I7zI/dNlWdB8Xp7v+3w9sX5N3J/WuJ1KOO5m26kRlHQo7EzT3974g
-----END CERTIFICATE-----
EOF
		sign_in_endpoint = "https://saml.provider/sign_in"
		sign_out_endpoint = "https://saml.provider/sign_out"
		user_id_attribute = "https://saml.provider/imi/ns/identity-200810"
		tenant_domain = "example.com"
		domain_aliases = ["example.com", "example.coz"]
		protocol_binding = "urn:oasis:names:tc:SAML:2.0:bindings:HTTP-Post"
		request_template = "<samlp:AuthnRequest xmlns:samlp=\"urn:oasis:names:tc:SAML:2.0:protocol\"\n@@AssertServiceURLAndDestination@@\n    ID=\"@@ID@@\"\n    IssueInstant=\"@@IssueInstant@@\"\n    ProtocolBinding=\"@@ProtocolBinding@@\" Version=\"2.0\">\n    <saml:Issuer xmlns:saml=\"urn:oasis:names:tc:SAML:2.0:assertion\">@@Issuer@@</saml:Issuer>\n</samlp:AuthnRequest>"
		signature_algorithm = "rsa-sha256"
		digest_algorithm = "sha256"
		icon_url = "https://example.com/logo.svg"
		fields_map = {
			foo = "bar"
			baz = "baa"
		}
		idp_initiated {
			client_id = "client_id"
			client_protocol = "samlp"
			client_authorize_query = "type=code&timeout=30"
		}
	}
}
`

const testConnectionSAMLConfigUpdate = `
resource "auth0_connection" "my_connection" {
	name = "Acceptance-Test-SAML-{{.random}}"
	strategy = "samlp"
	options {
		signing_cert = <<EOF
-----BEGIN CERTIFICATE-----
MIID6TCCA1ICAQEwDQYJKoZIhvcNAQEFBQAwgYsxCzAJBgNVBAYTAlVTMRMwEQYD
VQQIEwpDYWxpZm9ybmlhMRYwFAYDVQQHEw1TYW4gRnJhbmNpc2NvMRQwEgYDVQQK
EwtHb29nbGUgSW5jLjEMMAoGA1UECxMDRW5nMQwwCgYDVQQDEwNhZ2wxHTAbBgkq
hkiG9w0BCQEWDmFnbEBnb29nbGUuY29tMB4XDTA5MDkwOTIyMDU0M1oXDTEwMDkw
OTIyMDU0M1owajELMAkGA1UEBhMCQVUxEzARBgNVBAgTClNvbWUtU3RhdGUxITAf
BgNVBAoTGEludGVybmV0IFdpZGdpdHMgUHR5IEx0ZDEjMCEGA1UEAxMaZXVyb3Bh
LnNmby5jb3JwLmdvb2dsZS5jb20wggIiMA0GCSqGSIb3DQEBAQUAA4ICDwAwggIK
AoICAQC6pgYt7/EibBDumASF+S0qvqdL/f+nouJw2T1Qc8GmXF/iiUcrsgzh/Fd8
pDhz/T96Qg9IyR4ztuc2MXrmPra+zAuSf5bevFReSqvpIt8Duv0HbDbcqs/XKPfB
uMDe+of7a9GCywvAZ4ZUJcp0thqD9fKTTjUWOBzHY1uNE4RitrhmJCrbBGXbJ249
bvgmb7jgdInH2PU7PT55hujvOoIsQW2osXBFRur4pF1wmVh4W4lTLD6pjfIMUcML
ICHEXEN73PDic8KS3EtNYCwoIld+tpIBjE1QOb1KOyuJBNW6Esw9ALZn7stWdYcE
qAwvv20egN2tEXqj7Q4/1ccyPZc3PQgC3FJ8Be2mtllM+80qf4dAaQ/fWvCtOrQ5
pnfe9juQvCo8Y0VGlFcrSys/MzSg9LJ/24jZVgzQved/Qupsp89wVidwIzjt+WdS
fyWfH0/v1aQLvu5cMYuW//C0W2nlYziL5blETntM8My2ybNARy3ICHxCBv2RNtPI
WQVm+E9/W5rwh2IJR4DHn2LHwUVmT/hHNTdBLl5Uhwr4Wc7JhE7AVqb14pVNz1lr
5jxsp//ncIwftb7mZQ3DF03Yna+jJhpzx8CQoeLT6aQCHyzmH68MrHHT4MALPyUs
Pomjn71GNTtDeWAXibjCgdL6iHACCF6Htbl0zGlG0OAK+bdn0QIDAQABMA0GCSqG
SIb3DQEBBQUAA4GBAOKnQDtqBV24vVqvesL5dnmyFpFPXBn3WdFfwD6DzEb21UVG
5krmJiu+ViipORJPGMkgoL6BjU21XI95VQbun5P8vvg8Z+FnFsvRFY3e1CCzAVQY
ZsUkLw2I7zI/dNlWdB8Xp7v+3w9sX5N3J/WuJ1KOO5m26kRlHQo7EzT3974g
-----END CERTIFICATE-----
EOF
		sign_in_endpoint = "https://saml.provider/sign_in"
		sign_out_endpoint = ""
		tenant_domain = "example.com"
		domain_aliases = ["example.com", "example.coz"]
		protocol_binding = "urn:oasis:names:tc:SAML:2.0:bindings:HTTP-Post"
		signature_algorithm = "rsa-sha256"
		digest_algorithm = "sha256"
		fields_map = {
			foo = "bar"
			baz = "baa"
		}
		idp_initiated {
			client_id = "client_id"
			client_protocol = "samlp"
			client_authorize_query = "type=code&timeout=60"
		}
	}
}
`
