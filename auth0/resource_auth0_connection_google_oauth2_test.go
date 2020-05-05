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
	resource.AddTestSweepers("auth0_connection_google_oauth2", &resource.Sweeper{
		Name: "auth0_connection_google_oauth2",
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
					random.TestCheckResourceAttr("auth0_connection_google_oauth2.my_connection", "name", "Acceptance-Test-Google-OAuth2-{{.random}}", rand),
					resource.TestCheckResourceAttr("auth0_connection_google_oauth2.my_connection", "strategy", "google-oauth2"),
					resource.TestCheckResourceAttr("auth0_connection_google_oauth2.my_connection", "options.0.client_id", ""),
					resource.TestCheckResourceAttr("auth0_connection_google_oauth2.my_connection", "options.0.client_secret", ""),
					resource.TestCheckResourceAttr("auth0_connection_google_oauth2.my_connection", "options.0.allowed_audiences.#", "2"),
					resource.TestCheckResourceAttr("auth0_connection_google_oauth2.my_connection", "options.0.allowed_audiences.3506632655", "example.com"),
					resource.TestCheckResourceAttr("auth0_connection_google_oauth2.my_connection", "options.0.allowed_audiences.3154807651", "api.example.com"),
					resource.TestCheckResourceAttr("auth0_connection_google_oauth2.my_connection", "options.0.scopes.#", "4"),
					resource.TestCheckResourceAttr("auth0_connection_google_oauth2.my_connection", "options.0.scopes.881205744", "email"),
					resource.TestCheckResourceAttr("auth0_connection_google_oauth2.my_connection", "options.0.scopes.4080487570", "profile"),
				),
			},
		},
	})
}

const testAccConnectionGoogleOAuth2Config = `

resource "auth0_connection_google_oauth2" "my_connection" {
	name = "Acceptance-Test-Google-OAuth2-{{.random}}"
	is_domain_connection = false
	options {
		client_id = ""
		client_secret = ""
		allowed_audiences = [ "example.com", "api.example.com" ]
		scopes = [ "email", "profile", "gmail", "youtube" ]
	}
}
`

func TestAccConnectionGoogleOAuth2WithEnbledClients(t *testing.T) {

	rand := random.String(6)

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"auth0": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: random.Template(testAccGenericConnectionWithEnabledClientsConfig("google_oauth2"), rand),
				Check: resource.ComposeTestCheckFunc(
					random.TestCheckResourceAttr("auth0_connection_google_oauth2.my_connection", "name", "Acceptance-Test-Connection-{{.random}}", rand),
					resource.TestCheckResourceAttr("auth0_connection_google_oauth2.my_connection", "enabled_clients.#", "4"),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}
