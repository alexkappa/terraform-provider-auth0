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
	resource.AddTestSweepers("auth0_connection_waad", &resource.Sweeper{
		Name: "auth0_connection_waad",
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
					random.TestCheckResourceAttr("auth0_connection_waad.my_connection", "name", "Acceptance-Test-Azure-AD-{{.random}}", rand),
					resource.TestCheckResourceAttr("auth0_connection_waad.my_connection", "strategy", "waad"),
					resource.TestCheckResourceAttr("auth0_connection_waad.my_connection", "options.0.client_id", "123456"),
					resource.TestCheckResourceAttr("auth0_connection_waad.my_connection", "options.0.client_secret", "123456"),
					resource.TestCheckResourceAttr("auth0_connection_waad.my_connection", "options.0.tenant_domain", "example.onmicrosoft.com"),
					resource.TestCheckResourceAttr("auth0_connection_waad.my_connection", "options.0.domain", "example.onmicrosoft.com"),
					resource.TestCheckResourceAttr("auth0_connection_waad.my_connection", "options.0.domain_aliases.#", "2"),
					resource.TestCheckResourceAttr("auth0_connection_waad.my_connection", "options.0.domain_aliases.3506632655", "example.com"),
					resource.TestCheckResourceAttr("auth0_connection_waad.my_connection", "options.0.domain_aliases.3154807651", "api.example.com"),
					resource.TestCheckResourceAttr("auth0_connection_waad.my_connection", "options.0.scopes.#", "3"),
					resource.TestCheckResourceAttr("auth0_connection_waad.my_connection", "options.0.scopes.370042894", "basic_profile"),
					resource.TestCheckResourceAttr("auth0_connection_waad.my_connection", "options.0.scopes.1268340351", "ext_profile"),
					resource.TestCheckResourceAttr("auth0_connection_waad.my_connection", "options.0.scopes.541325467", "ext_groups"),
				),
			},
		},
	})
}

const testAccConnectionAzureADConfig = `

resource "auth0_connection_waad" "my_connection" {
	name     = "Acceptance-Test-Azure-AD-{{.random}}"
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

func TestAccConnectionAzureADWithEnbledClients(t *testing.T) {

	rand := random.String(6)

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"auth0": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: random.Template(testAccGenericConnectionWithEnabledClientsConfig("waad"), rand),
				Check: resource.ComposeTestCheckFunc(
					random.TestCheckResourceAttr("auth0_connection_waad.my_connection", "name", "Acceptance-Test-Connection-{{.random}}", rand),
					resource.TestCheckResourceAttr("auth0_connection_waad.my_connection", "enabled_clients.#", "4"),
				),
			},
		},
	})
}
