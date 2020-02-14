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
	resource.AddTestSweepers("auth0_role", &resource.Sweeper{
		Name: "auth0_role",
		F: func(_ string) error {
			api, err := Auth0()
			if err != nil {
				return err
			}
			var page int
			for {
				l, err := api.Role.List(management.Page(page))
				if err != nil {
					return err
				}
				for _, role := range l.Roles {
					if strings.Contains(role.GetName(), "Acceptance Test") {
						if e := api.Role.Delete(role.GetID()); e != nil {
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

func TestAccRole(t *testing.T) {

	rand := random.String(6)

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"auth0": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: random.Template(testAccRoleCreate, rand),
				Check: resource.ComposeTestCheckFunc(
					random.TestCheckResourceAttr("auth0_role.the_one", "name", "The One - Acceptance Test - {{.random}}", rand),
					resource.TestCheckResourceAttr("auth0_role.the_one", "description", "The One - Acceptance Test"),
					resource.TestCheckResourceAttr("auth0_role.the_one", "permissions.#", "1"),
				),
			},
			{
				Config: random.Template(testAccRoleUpdate, rand),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("auth0_role.the_one", "description", "The One who will bring peace - Acceptance Test"),
					resource.TestCheckResourceAttr("auth0_role.the_one", "permissions.#", "2"),
				),
			},
		},
	})
}

const testAccRoleAux = `

resource auth0_resource_server matrix {
    name = "Role - Acceptance Test - {{.random}}"
    identifier = "https://matrix.com/"
    scopes {
        value = "stop:bullets"
        description = "Stop bullets"
    }
    scopes {
        value = "bring:peace"
        description = "Bring peace"
    }
}`

const testAccRoleCreate = testAccRoleAux + `

resource auth0_role the_one {
  name = "The One - Acceptance Test - {{.random}}"
  description = "The One - Acceptance Test"
  permissions {
    name = "stop:bullets"
    resource_server_identifier = auth0_resource_server.matrix.identifier
  }
}
`

const testAccRoleUpdate = testAccRoleAux + `

resource auth0_role the_one {
  name = "The One - Acceptance Test - {{.random}}"
  description = "The One who will bring peace - Acceptance Test"
  permissions {
    name = "stop:bullets"
    resource_server_identifier = auth0_resource_server.matrix.identifier
  }
  permissions {
    name = "bring:peace"
    resource_server_identifier = auth0_resource_server.matrix.identifier
  }
}
`
