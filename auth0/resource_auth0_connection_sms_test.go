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
	resource.AddTestSweepers("auth0_connection_sms", &resource.Sweeper{
		Name: "auth0_connection_sms",
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
					random.TestCheckResourceAttr("auth0_connection_sms.my_connection", "name", "Acceptance-Test-SMS-{{.random}}", rand),
					resource.TestCheckResourceAttr("auth0_connection_sms.my_connection", "strategy", "sms"),
					resource.TestCheckResourceAttr("auth0_connection_sms.my_connection", "options.0.twilio_sid", "ABC123"),
					resource.TestCheckResourceAttr("auth0_connection_sms.my_connection", "options.0.twilio_token", "DEF456"),
					resource.TestCheckResourceAttr("auth0_connection_sms.my_connection", "options.0.totp.#", "1"),
					resource.TestCheckResourceAttr("auth0_connection_sms.my_connection", "options.0.totp.0.time_step", "300"),
					resource.TestCheckResourceAttr("auth0_connection_sms.my_connection", "options.0.totp.0.length", "6"),
				),
			},
		},
	})
}

const testAccConnectionSMSConfig = `

resource "auth0_connection_sms" "my_connection" {
	name = "Acceptance-Test-SMS-{{.random}}"
	is_domain_connection = false

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

func TestAccConnectionSMSWithEnbledClients(t *testing.T) {

	rand := random.String(6)

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"auth0": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: random.Template(testAccGenericConnectionWithEnabledClientsConfig("sms"), rand),
				Check: resource.ComposeTestCheckFunc(
					random.TestCheckResourceAttr("auth0_connection_sms.my_connection", "name", "Acceptance-Test-Connection-{{.random}}", rand),
					resource.TestCheckResourceAttr("auth0_connection_sms.my_connection", "enabled_clients.#", "4"),
				),
			},
		},
	})
}
