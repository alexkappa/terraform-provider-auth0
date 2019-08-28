package auth0

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccUserRole(t *testing.T) {

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"auth0": Provider(),
		},
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccUserRole,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("auth0_user_role.my_user_role", "user_id", "auth0|1234567890"),
					resource.TestCheckResourceAttr("auth0_user_role.my_user_role", "roles.0", "roleId"),
				),
			},
		},
	})
}

const testAccUserRole = `
provider "auth0" {}

resource "auth0_user" "my_user" {
  connection_name = "Username-Password-Authentication"
  user_id = "auth0|1234567890"
  email = "test@test.com"
  password = "passpass$12$12"
  nickname = "testnick"
}

resource "auth0_role" "my_role" {
	name = "User Role Acceptance Test"
}

resource "auth0_user_role" "my_user_role" {
	user_id = "${auth0_user.my_user.user_id}"
	roles = [ "${auth0_role.my_role.id}" ]
}
`
