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
	resource.AddTestSweepers("auth0_connection_auth0", &resource.Sweeper{
		Name: "auth0_connection_auth0",
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

func TestAccConnectionAuth0(t *testing.T) {

	rand := random.String(6)

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"auth0": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: random.Template(testAccConnectionAuth0Config, rand),
				Check: resource.ComposeTestCheckFunc(
					random.TestCheckResourceAttr("auth0_connection_auth0.my_connection", "name", "Acceptance-Test-Connection-{{.random}}", rand),
					resource.TestCheckResourceAttr("auth0_connection_auth0.my_connection", "is_domain_connection", "true"),
					resource.TestCheckResourceAttr("auth0_connection_auth0.my_connection", "strategy", "auth0"),
					resource.TestCheckResourceAttr("auth0_connection_auth0.my_connection", "options.0.password_policy", "fair"),
					resource.TestCheckResourceAttr("auth0_connection_auth0.my_connection", "options.0.password_no_personal_info.0.enable", "true"),
					resource.TestCheckResourceAttr("auth0_connection_auth0.my_connection", "options.0.password_dictionary.0.enable", "true"),
					resource.TestCheckResourceAttr("auth0_connection_auth0.my_connection", "options.0.password_complexity_options.0.min_length", "6"),
					resource.TestCheckResourceAttr("auth0_connection_auth0.my_connection", "options.0.enabled_database_customization", "false"),
					resource.TestCheckResourceAttr("auth0_connection_auth0.my_connection", "options.0.brute_force_protection", "true"),
					resource.TestCheckResourceAttr("auth0_connection_auth0.my_connection", "options.0.import_mode", "false"),
					resource.TestCheckResourceAttr("auth0_connection_auth0.my_connection", "options.0.disable_signup", "false"),
					resource.TestCheckResourceAttr("auth0_connection_auth0.my_connection", "options.0.requires_username", "true"),
					resource.TestCheckResourceAttr("auth0_connection_auth0.my_connection", "options.0.custom_scripts.get_user", "myFunction"),
					resource.TestCheckResourceAttrSet("auth0_connection_auth0.my_connection", "options.0.configuration.foo"),
				),
			},
			{
				Config: random.Template(testAccConnectionAuth0ConfigUpdate, rand),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("auth0_connection_auth0.my_connection", "options.0.brute_force_protection", "false"),
				),
			},
		},
	})
}

const testAccConnectionAuth0Config = `

resource "auth0_connection_auth0" "my_connection" {
	name = "Acceptance-Test-Connection-{{.random}}"
	is_domain_connection = true
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

const testAccConnectionAuth0ConfigUpdate = `

resource "auth0_connection_auth0" "my_connection" {
	name = "Acceptance-Test-Connection-{{.random}}"
	is_domain_connection = true
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

func TestAccConnectionAuth0WithEnbledClients(t *testing.T) {

	rand := random.String(6)

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"auth0": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: random.Template(testAccGenericConnectionWithEnabledClientsConfig("auth0"), rand),
				Check: resource.ComposeTestCheckFunc(
					random.TestCheckResourceAttr("auth0_connection_auth0.my_connection", "name", "Acceptance-Test-Connection-{{.random}}", rand),
					resource.TestCheckResourceAttr("auth0_connection_auth0.my_connection", "enabled_clients.#", "4"),
				),
			},
		},
	})
}

func TestAccConnectionAuth0Configuration(t *testing.T) {

	rand := random.String(6)

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"auth0": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: random.Template(testAccConnectionAuth0ConfigurationCreate, rand),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("auth0_connection_auth0.my_connection", "options.0.configuration.%", "2"),
					resource.TestCheckResourceAttr("auth0_connection_auth0.my_connection", "options.0.configuration.foo", "xxx"),
					resource.TestCheckResourceAttr("auth0_connection_auth0.my_connection", "options.0.configuration.bar", "zzz"),
				),
			},
			{
				Config: random.Template(testAccConnectionAuth0ConfigurationUpdate, rand),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("auth0_connection_auth0.my_connection", "options.0.configuration.%", "3"),
					resource.TestCheckResourceAttr("auth0_connection_auth0.my_connection", "options.0.configuration.foo", "xxx"),
					resource.TestCheckResourceAttr("auth0_connection_auth0.my_connection", "options.0.configuration.bar", "yyy"),
					resource.TestCheckResourceAttr("auth0_connection_auth0.my_connection", "options.0.configuration.baz", "zzz"),
				),
			},
		},
	})
}

const testAccConnectionAuth0ConfigurationCreate = `
	resource "auth0_connection_auth0" "my_connection" {
		name = "Acceptance-Test-Connection-{{.random}}"
		is_domain_connection = true
		options {
			configuration = {
				foo = "xxx"
				bar = "zzz"
			}
		}
	}
`

const testAccConnectionAuth0ConfigurationUpdate = `
	resource "auth0_connection_auth0" "my_connection" {
		name = "Acceptance-Test-Connection-{{.random}}"
		is_domain_connection = true
		options {
			configuration = {
				foo = "xxx"
				bar = "yyy"
				baz = "zzz"
			}
		}
	}
`
