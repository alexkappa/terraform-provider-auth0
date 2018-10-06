package auth0

import (
	"fmt"
	"os"
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
				ExpectError: regexp.MustCompile(`config is invalid: auth0_user.user: "conn": required field is not set`),
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
				Config: fmt.Sprintf(testAccUserCreateUser, os.Getenv("AUTH0_CLIENT_ID")),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("auth0_user.user", "user_id", "auth0|12345"),
					resource.TestCheckResourceAttr("auth0_user.user", "email", "test@test.com"),
					resource.TestCheckResourceAttr("auth0_user.user", "conn", "database-acc-user"),
				),
			},
		},
	})
}

const testAccUserCreateUser = `
provider "auth0" {}

resource "auth0_connection" "database_acc_user" {
  name = "database-acc-user"
  strategy = "auth0"
  enabled_clients = [
    "%s"
  ]
  options = {
    password_policy = "low"
  }
}

resource "auth0_user" "user" {
  conn = "database-acc-user"
  user_id = "12345"
  email = "test@test.com"
  password = "testtest$12$12"

  depends_on = ["auth0_connection.database_acc_user"]
}
		`
