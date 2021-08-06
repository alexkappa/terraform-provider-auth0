package auth0

import (
	"log"
	"regexp"
	"testing"

	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"gopkg.in/auth0.v5/management"

	"github.com/alexkappa/terraform-provider-auth0/auth0/internal/random"
)

func init() {
	resource.AddTestSweepers("auth0_user_roles", &resource.Sweeper{
		Name: "auth0_user",
		F: func(_ string) error {
			api, err := Auth0()
			if err != nil {
				return err
			}
			var page int
			for {
				l, err := api.User.Search(
					management.Page(page),
					management.Query(`email.domain:"acceptance.test.com"`))
				if err != nil {
					return err
				}
				for _, user := range l.Users {
					log.Printf("[DEBUG] ✗ %s", user.GetName())
					if e := api.User.Delete(user.GetID()); e != nil {
						multierror.Append(err, e)
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

func TestAccUserRolesMissingRequiredParams(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"auth0": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config:      testAccUserRolesMissingUserId,
				ExpectError: regexp.MustCompile(`The argument "user_id" is required`),
			},
			{
				Config:      testAccUserRolesMissingRoles,
				ExpectError: regexp.MustCompile(`The argument "roles" is required`),
			},
		},
	})
}

func TestAccUserRoles(t *testing.T) {

	rand := random.String(6)

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"auth0": Provider(),
		},
		// ExpectNonEmptyPlan is set since we are creating the user and then modifying it.
		// This wouldn't normally be the case of how you use this, but for a test
		Steps: []resource.TestStep{
			{
				Config: random.Template(testAccUserRolesCreate, rand),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("auth0_user.user", "roles.#", "0"),
					resource.TestCheckResourceAttr("auth0_user_roles.user", "roles.#", "1"),
					resource.TestCheckResourceAttr("auth0_role.owner", "name", "owner"),
				),
				ExpectNonEmptyPlan: true,
			},
			{
				Config: random.Template(testAccUserRolesAddRole, rand),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("auth0_user.user", "roles.#", "0"),
					resource.TestCheckResourceAttr("auth0_user_roles.user", "roles.#", "2"),
					resource.TestCheckResourceAttr("auth0_role.owner", "name", "owner"),
					resource.TestCheckResourceAttr("auth0_role.admin", "name", "admin"),
				),
				ExpectNonEmptyPlan: true,
			},
			{
				Config: random.Template(testAccUserRolesRemoveRole, rand),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("auth0_user.user", "roles.#", "0"),
					resource.TestCheckResourceAttr("auth0_user_roles.user", "roles.#", "1"),
					resource.TestCheckResourceAttr("auth0_role.admin", "name", "admin"),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

const testAccUserRolesMissingUserId = `

resource auth0_user_roles user {
	roles = [ "rol_1" ]
}
`

const testAccUserRolesMissingRoles = `

resource auth0_user_roles user {
	user_id = "{{.random}}"
}
`

const testAccUserRolesCreate = `

resource auth0_user user {
	connection_name = "Username-Password-Authentication"
	username = "{{.random}}"
	email = "{{.random}}@acceptance.test.com"
	password = "passpass$12$12"
	name = "Firstname Lastname"
	given_name = "Firstname"
	family_name = "Lastname"
	nickname = "{{.random}}"
}

resource auth0_user_roles user {
	user_id = auth0_user.user.user_id
	roles = [ auth0_role.owner.id ]
}

resource auth0_role owner {
	name = "owner"
	description = "Owner"
}
`

const testAccUserRolesAddRole = `

resource auth0_user user {
	connection_name = "Username-Password-Authentication"
	username = "{{.random}}"
	email = "{{.random}}@acceptance.test.com"
	password = "passpass$12$12"
	name = "Firstname Lastname"
	given_name = "Firstname"
	family_name = "Lastname"
	nickname = "{{.random}}"
}

resource auth0_user_roles user {
	user_id = auth0_user.user.user_id
	roles = [ auth0_role.owner.id, auth0_role.admin.id ]
}

resource auth0_role owner {
	name = "owner"
	description = "Owner"
}

resource auth0_role admin {
	name = "admin"
	description = "Admin"
}
`

const testAccUserRolesRemoveRole = `

resource auth0_user user {
	connection_name = "Username-Password-Authentication"
	username = "{{.random}}"
	email = "{{.random}}@acceptance.test.com"
	password = "passpass$12$12"
	name = "Firstname Lastname"
	given_name = "Firstname"
	family_name = "Lastname"
	nickname = "{{.random}}"
}

resource auth0_user_roles user {
	user_id = auth0_user.user.user_id
	roles = [ auth0_role.admin.id ]
}

resource auth0_role admin {
	name = "admin"
	description = "Admin"
}
`
