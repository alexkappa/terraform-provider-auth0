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
	resource.AddTestSweepers("auth0_connection_github", &resource.Sweeper{
		Name: "auth0_connection_github",
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
					random.TestCheckResourceAttr("auth0_connection_github.my_connection", "name", "Acceptance-Test-GitHub-{{.random}}", rand),
					resource.TestCheckResourceAttr("auth0_connection_github.my_connection", "strategy", "github"),
					resource.TestCheckResourceAttr("auth0_connection_github.my_connection", "options.0.client_id", "client-id"),
					resource.TestCheckResourceAttr("auth0_connection_github.my_connection", "options.0.client_secret", "client-secret"),
					resource.TestCheckResourceAttr("auth0_connection_github.my_connection", "options.0.scopes.#", "20"),
					resource.TestCheckResourceAttr("auth0_connection_github.my_connection", "options.0.scopes.881205744", "email"),
					resource.TestCheckResourceAttr("auth0_connection_github.my_connection", "options.0.scopes.4080487570", "profile"),
					resource.TestCheckResourceAttr("auth0_connection_github.my_connection", "options.0.scopes.862208977", "follow"),
					resource.TestCheckResourceAttr("auth0_connection_github.my_connection", "options.0.scopes.347111084", "read_repo_hook"),
					resource.TestCheckResourceAttr("auth0_connection_github.my_connection", "options.0.scopes.718177942", "admin_public_key"),
					resource.TestCheckResourceAttr("auth0_connection_github.my_connection", "options.0.scopes.2480957806", "write_public_key"),
					resource.TestCheckResourceAttr("auth0_connection_github.my_connection", "options.0.scopes.356496889", "write_repo_hook"),
					resource.TestCheckResourceAttr("auth0_connection_github.my_connection", "options.0.scopes.3006585776", "write_org"),
					resource.TestCheckResourceAttr("auth0_connection_github.my_connection", "options.0.scopes.855904415", "read_user"),
					resource.TestCheckResourceAttr("auth0_connection_github.my_connection", "options.0.scopes.1560560783", "admin_repo_hook"),
					resource.TestCheckResourceAttr("auth0_connection_github.my_connection", "options.0.scopes.2933527251", "admin_org"),
					resource.TestCheckResourceAttr("auth0_connection_github.my_connection", "options.0.scopes.1314370975", "repo"),
					resource.TestCheckResourceAttr("auth0_connection_github.my_connection", "options.0.scopes.2175618052", "repo_status"),
					resource.TestCheckResourceAttr("auth0_connection_github.my_connection", "options.0.scopes.188173322", "read_org"),
					resource.TestCheckResourceAttr("auth0_connection_github.my_connection", "options.0.scopes.133261078", "gist"),
					resource.TestCheckResourceAttr("auth0_connection_github.my_connection", "options.0.scopes.1820025999", "repo_deployment"),
					resource.TestCheckResourceAttr("auth0_connection_github.my_connection", "options.0.scopes.3220703903", "public_repo"),
					resource.TestCheckResourceAttr("auth0_connection_github.my_connection", "options.0.scopes.2092139895", "notifications"),
					resource.TestCheckResourceAttr("auth0_connection_github.my_connection", "options.0.scopes.672436223", "delete_repo"),
					resource.TestCheckResourceAttr("auth0_connection_github.my_connection", "options.0.scopes.2296398814", "read_public_key"),
				),
			},
		},
	})
}

const testAccConnectionGitHubConfig = `

resource "auth0_connection_github" "my_connection" {
	name = "Acceptance-Test-GitHub-{{.random}}"
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

func TestAccConnectionGitHubWithEnbledClients(t *testing.T) {

	rand := random.String(6)

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"auth0": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: random.Template(testAccGenericConnectionWithEnabledClientsConfig("github"), rand),
				Check: resource.ComposeTestCheckFunc(
					random.TestCheckResourceAttr("auth0_connection_github.my_connection", "name", "Acceptance-Test-Connection-{{.random}}", rand),
					resource.TestCheckResourceAttr("auth0_connection_github.my_connection", "enabled_clients.#", "4"),
				),
			},
		},
	})
}
