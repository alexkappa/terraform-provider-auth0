package auth0

import (
	"strings"
	"testing"

	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/terraform-providers/terraform-provider-auth0/auth0/internal/random"
	"gopkg.in/auth0.v4/management"
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
    identifier = "https://{{.random}}.matrix.com/"
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

func TestAccRolePermissions(t *testing.T) {

	rand := random.String(6)

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"auth0": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: random.Template(testAccRolePermissions, rand),
				Check: resource.ComposeTestCheckFunc(
					random.TestCheckResourceAttr("auth0_role.role", "name", "The One - Acceptance Test - {{.random}}", rand),
					resource.TestCheckResourceAttr("auth0_role.role", "description", "The One - Acceptance Test"),
					resource.TestCheckResourceAttr("auth0_role.role", "permissions.#", "58"),
				),
			},
		},
	})
}

const testAccRolePermissions = `

locals {
	permissions = {
		"permission:1"   = "Permission 1"
		"permission:2"   = "Permission 2"
		"permission:3"   = "Permission 3"
		"permission:4"   = "Permission 4"
		"permission:5"   = "Permission 5"
		"permission:6"   = "Permission 6"
		"permission:7"   = "Permission 7"
		"permission:8"   = "Permission 8"
		"permission:9"   = "Permission 9"
		"permission:10"   = "Permission 10"
		"permission:11"   = "Permission 11"
		"permission:12"   = "Permission 12"
		"permission:13"   = "Permission 13"
		"permission:14"   = "Permission 14"
		"permission:15"   = "Permission 15"
		"permission:16"   = "Permission 16"
		"permission:17"   = "Permission 17"
		"permission:18"   = "Permission 18"
		"permission:19"   = "Permission 19"
		"permission:20"   = "Permission 20"
		"permission:21"   = "Permission 21"
		"permission:22"   = "Permission 22"
		"permission:23"   = "Permission 23"
		"permission:24"   = "Permission 24"
		"permission:25"   = "Permission 25"
		"permission:26"   = "Permission 26"
		"permission:27"   = "Permission 27"
		"permission:28"   = "Permission 28"
		"permission:29"   = "Permission 29"
		"permission:30"   = "Permission 30"
		"permission:31"   = "Permission 31"
		"permission:32"   = "Permission 32"
		"permission:33"   = "Permission 33"
		"permission:34"   = "Permission 34"
		"permission:35"   = "Permission 35"
		"permission:36"   = "Permission 36"
		"permission:37"   = "Permission 37"
		"permission:38"   = "Permission 38"
		"permission:39"   = "Permission 39"
		"permission:40"   = "Permission 40"
		"permission:41"   = "Permission 41"
		"permission:42"   = "Permission 42"
		"permission:43"   = "Permission 43"
		"permission:44"   = "Permission 44"
		"permission:45"   = "Permission 45"
		"permission:46"   = "Permission 46"
		"permission:47"   = "Permission 47"
		"permission:48"   = "Permission 48"
		"permission:49"   = "Permission 49"
		"permission:50"   = "Permission 50"
		"permission:51"   = "Permission 51"
		"permission:52"   = "Permission 52"
		"permission:53"   = "Permission 53"
		"permission:54"   = "Permission 54"
		"permission:55"   = "Permission 55"
		"permission:56"   = "Permission 56"
		"permission:57"   = "Permission 57"
		"permission:58"   = "Permission 58"
	}
}

resource auth0_resource_server server {
	name = "Role - Acceptance Test - {{.random}}"
	identifier = "https://{{.random}}.matrix.com/"

	dynamic scopes {
		for_each = local.permissions
		iterator = permission
		content {
			value       = permission.key
			description = permission.value
		}
	}
}

resource auth0_role role {
	name = "The One - Acceptance Test - {{.random}}"
	description = "The One - Acceptance Test"
	dynamic permissions {
		for_each = local.permissions
		iterator = permission
		content {
			name = permission.key
			resource_server_identifier = auth0_resource_server.server.identifier
		}
	}
  }
`
