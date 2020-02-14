package auth0

import (
	"strings"
	"testing"

	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
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
				l, err := api.Connection.List(management.Page(page))
				if err != nil {
					return err
				}
				for _, connection := range l.Connections {
					if strings.Contains(connection.GetName(), "Acceptance-Test") {
						if e := api.Client.Delete(connection.GetID()); e != nil {
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
					random.TestCheckResourceAttr("auth0_connection.my_ad_connection", "name", "Acceptance-Test-AD-Connection-{{.random}}", rand),
					resource.TestCheckResourceAttr("auth0_connection.my_ad_connection", "strategy", "ad"),
				),
			},
		},
	})
}

const testAccConnectionADConfig = `

resource "auth0_connection" "my_ad_connection" {
	name = "Acceptance-Test-AD-Connection-{{.random}}"
	strategy = "ad"
}
`

func TestAccConnectionWAAD(t *testing.T) {

	rand := random.String(6)

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"auth0": Provider(),
		},
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: random.Template(testAccConnectionWAADConfig, rand),
				Check: resource.ComposeTestCheckFunc(
					random.TestCheckResourceAttr("auth0_connection.my_connection", "name", "Acceptance-Test-WAAD-Connection-{{.random}}", rand),
					resource.TestCheckResourceAttr("auth0_connection.my_connection", "strategy", "waad"),
				),
			},
		},
	})
}

const testAccConnectionWAADConfig = `

resource "auth0_connection" "my_connection" {
	name     = "Acceptance-Test-WAAD-Connection-{{.random}}"
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

func TestAccTwilioConnection(t *testing.T) {

	rand := random.String(6)

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"auth0": Provider(),
		},
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: random.Template(testAccTwilioConnectionConfig, rand),
				Check: resource.ComposeTestCheckFunc(
					random.TestCheckResourceAttr("auth0_connection.sms_connection", "name", "Acceptance-Test-SMS-Connection-{{.random}}", rand),
					resource.TestCheckResourceAttr("auth0_connection.sms_connection", "strategy", "sms"),
					resource.TestCheckResourceAttr("auth0_connection.sms_connection", "options.0.twilio_token", "DEF456"),
				),
			},
		},
	})
}

const testAccTwilioConnectionConfig = `

resource "auth0_connection" "sms_connection" {
	name = "Acceptance-Test-SMS-Connection-{{.random}}"
	is_domain_connection = false
	strategy = "sms"
	
	options {
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

func TestAccSalesforceCommunityConnection(t *testing.T) {

	rand := random.String(6)

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"auth0": Provider(),
		},
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: random.Template(testAccSalesforceConnectionConfig, rand),
				Check: resource.ComposeTestCheckFunc(
					random.TestCheckResourceAttr("auth0_connection.salesforce_community", "name", "Acceptance-Test-Salesforce-Connection-{{.random}}", rand),
					resource.TestCheckResourceAttr("auth0_connection.salesforce_community", "strategy", "salesforce-community"),
					resource.TestCheckResourceAttr("auth0_connection.salesforce_community", "options.0.community_base_url", "https://salesforce-community.example"),
				),
			},
		},
	})
}

const testAccSalesforceConnectionConfig = `

resource "auth0_connection" "salesforce_community" {
	name = "Acceptance-Test-Salesforce-Connection-{{.random}}"
	is_domain_connection = false
	strategy = "salesforce-community"
	
	options {
		client_id = false
		client_secret = "sms-connection"
		community_base_url = "https://salesforce-community.example"
	}
}
`
