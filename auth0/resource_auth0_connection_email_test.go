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
	resource.AddTestSweepers("auth0_connection_email", &resource.Sweeper{
		Name: "auth0_connection_email",
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
					random.TestCheckResourceAttr("auth0_connection_email.my_connection", "name", "Acceptance-Test-Email-{{.random}}", rand),
					resource.TestCheckResourceAttr("auth0_connection_email.my_connection", "strategy", "email"),
					resource.TestCheckResourceAttr("auth0_connection_email.my_connection", "options.0.from", "Magic Password <password@example.com>"),
					resource.TestCheckResourceAttr("auth0_connection_email.my_connection", "options.0.subject", "Sign in!"),
					resource.TestCheckResourceAttr("auth0_connection_email.my_connection", "options.0.totp.#", "1"),
					resource.TestCheckResourceAttr("auth0_connection_email.my_connection", "options.0.totp.0.time_step", "300"),
					resource.TestCheckResourceAttr("auth0_connection_email.my_connection", "options.0.totp.0.length", "6"),
				),
			},
			{
				Config: random.Template(testAccConnectionEmailConfigUpdate, rand),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("auth0_connection_email.my_connection", "options.0.totp.#", "1"),
					resource.TestCheckResourceAttr("auth0_connection_email.my_connection", "options.0.totp.0.time_step", "360"),
					resource.TestCheckResourceAttr("auth0_connection_email.my_connection", "options.0.totp.0.length", "4"),
				),
			},
		},
	})
}

const testAccConnectionEmailConfig = `

resource "auth0_connection_email" "my_connection" {
	name = "Acceptance-Test-Email-{{.random}}"
	is_domain_connection = false

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

resource "auth0_connection_email" "my_connection" {
	name = "Acceptance-Test-Email-{{.random}}"
	is_domain_connection = false

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

func TestAccConnectionEmailWithEnbledClients(t *testing.T) {

	rand := random.String(6)

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"auth0": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: random.Template(testAccGenericConnectionWithEnabledClientsConfig("email"), rand),
				Check: resource.ComposeTestCheckFunc(
					random.TestCheckResourceAttr("auth0_connection_email.my_connection", "name", "Acceptance-Test-Connection-{{.random}}", rand),
					resource.TestCheckResourceAttr("auth0_connection_email.my_connection", "enabled_clients.#", "4"),
				),
			},
		},
	})
}
