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
			resource.TestStep{
				Config: testAccRole,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("auth0_role.my_role", "name", "Application - Role Acceptance Test"),
					resource.TestCheckResourceAttr("auth0_role.my_role", "description", "Test Applications Role Long Description"),
				),
			},
		},
	})
}

const testAccRole = `
provider "auth0" {}

resource "auth0_role" "my_role" {
	name = "Application - Role Acceptance Test"
	description = "Test Applications Role Long Description"
}
`
