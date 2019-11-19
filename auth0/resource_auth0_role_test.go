package auth0

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccRole(t *testing.T) {

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"auth0": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: testAccRoleCreate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("auth0_role.my_role", "name", "Application - Role Acceptance Test"),
					resource.TestCheckResourceAttr("auth0_role.my_role", "description", "Test Applications Role Long Description"),
				),
			},
			{
				Config: testAccRoleUpdate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("auth0_role.my_role", "description", "Test Applications Role Long Description And Then Some"),
					resource.TestCheckResourceAttr("auth0_role.my_role", "user_ids.0", "auth0|neo"),
					resource.TestCheckResourceAttr("auth0_role.my_role", "user_ids.1", "auth0|trinity"),
				),
			},
		},
	})
}

const testAccRoleCreate = `
provider "auth0" {}

resource "auth0_role" "my_role" {
	name = "Application - Role Acceptance Test"
	description = "Test Applications Role Long Description"
}
`

const testAccRoleUpdate = `
provider "auth0" {}

resource "auth0_user" "neo" {
  connection_name = "Username-Password-Authentication"
  email = "neo@matrix.com"
  username = "neo"
  nickname = "neo"
  password = "IAmThe#1"
  user_id = "neo"
}

resource "auth0_user" "trinity" {
  connection_name = "Username-Password-Authentication"
  email = "trinity@matrix.com"
  username = "trinity"
  nickname = "trinity"
  password = "TheM4trixH4$Y0u"
  user_id = "trinity"
}

resource "auth0_role" "my_role" {
  name = "Application Role Acceptance Test"
  description = "Test Applications Role Long Description And Then Some"
  user_ids = [ 
    auth0_user.neo.id, 
    auth0_user.trinity.id
  ]
}
`
