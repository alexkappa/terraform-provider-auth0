package auth0

import (
	"regexp"

	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccUserMissingRequiredParams(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"auth0": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config:      testAccUserMissingRequiredParam,
				ExpectError: regexp.MustCompile(`config is invalid: auth0_user.user: "connection_name": required field is not set`),
			},
		},
	})
}

const testAccUserMissingRequiredParam = `
provider "auth0" {}

resource "auth0_user" "user" {}
`

func TestAccUserCreateUser(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"auth0": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: testAccUserCreateUser,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("auth0_user.user", "user_id", "12345"),
					resource.TestCheckResourceAttr("auth0_user.user", "email", "test@test.com"),
					resource.TestCheckResourceAttr("auth0_user.user", "connection_name", "Username-Password-Authentication"),
				),
			},
		},
	})
}

const testAccUserCreateUser = `
provider "auth0" {}

resource "auth0_user" "user" {
  connection_name = "Username-Password-Authentication"
  user_id = "12345"
  email = "test@test.com"
  password = "testtest$12$12"
}
`
