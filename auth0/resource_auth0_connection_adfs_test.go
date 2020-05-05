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
	resource.AddTestSweepers("auth0_connection_adfs", &resource.Sweeper{
		Name: "auth0_connection_adfs",
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

func TestAccConnectionADFS(t *testing.T) {

	rand := random.String(6)

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"auth0": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: random.Template(testAccConnectionADFSConfig, rand),
				Check: resource.ComposeTestCheckFunc(
					random.TestCheckResourceAttr("auth0_connection_adfs.my_connection", "name", "Acceptance-Test-ADFS-{{.random}}", rand),
					resource.TestCheckResourceAttr("auth0_connection_adfs.my_connection", "strategy", "adfs"),
					resource.TestCheckResourceAttr("auth0_connection_adfs.my_connection", "options.0.tenant_domain", "example.com"),
					resource.TestCheckResourceAttr("auth0_connection_adfs.my_connection", "options.0.adfs_server", "srv.example.com"),
					resource.TestCheckResourceAttr("auth0_connection_adfs.my_connection", "options.0.api_enable_users", "true"),
					resource.TestCheckResourceAttr("auth0_connection_adfs.my_connection", "options.0.domain_aliases.#", "2"),
					resource.TestCheckResourceAttr("auth0_connection_adfs.my_connection", "options.0.domain_aliases.3506632655", "example.com"),
					resource.TestCheckResourceAttr("auth0_connection_adfs.my_connection", "options.0.domain_aliases.3154807651", "api.example.com"),
				),
			},
		},
	})
}

const testAccConnectionADFSConfig = `

resource "auth0_connection_adfs" "my_connection" {
	name     = "Acceptance-Test-ADFS-{{.random}}"
	options {
		tenant_domain = "example.com"
		domain_aliases = [
			"example.com",
			"api.example.com"
		]
		adfs_server = "srv.example.com"
		api_enable_users = true
	}
}
`

func TestAccConnectionADFSWithEnbledClients(t *testing.T) {

	rand := random.String(6)

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"auth0": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: random.Template(testAccGenericConnectionWithEnabledClientsConfig("adfs"), rand),
				Check: resource.ComposeTestCheckFunc(
					random.TestCheckResourceAttr("auth0_connection_adfs.my_connection", "name", "Acceptance-Test-Connection-{{.random}}", rand),
					resource.TestCheckResourceAttr("auth0_connection_adfs.my_connection", "enabled_clients.#", "4"),
				),
			},
		},
	})
}
