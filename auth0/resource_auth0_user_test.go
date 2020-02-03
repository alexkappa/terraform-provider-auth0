package auth0

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccUserMissingRequiredParams(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"auth0": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config:      testAccUserMissingRequiredParam,
				ExpectError: regexp.MustCompile(`The argument "connection_name" is required`),
			},
		},
	})
}

const testAccUserMissingRequiredParam = `
provider "auth0" {}

resource "auth0_user" "user" {}
`

func TestAccUser(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"auth0": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: testAccUser_create,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("auth0_user.user", "user_id", "auth0|12345"),
					resource.TestCheckResourceAttr("auth0_user.user", "email", "test@test.com"),
					resource.TestCheckResourceAttr("auth0_user.user", "name", "Firstname Lastname"),
					resource.TestCheckResourceAttr("auth0_user.user", "family_name", "Lastname"),
					resource.TestCheckResourceAttr("auth0_user.user", "given_name", "Firstname"),
					resource.TestCheckResourceAttr("auth0_user.user", "nickname", "some.nickname"),
					resource.TestCheckResourceAttr("auth0_user.user", "connection_name", "Username-Password-Authentication"),
					resource.TestCheckResourceAttr("auth0_user.user", "roles.#", "0"),
					resource.TestCheckResourceAttr("auth0_user.user", "picture", "https://www.example.com/a-valid-picture-url.jpg"),
				),
			},
			{
				Config: testAccUser_update,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("auth0_user.user", "roles.#", "2"),
					resource.TestCheckResourceAttr("auth0_role.owner", "name", "owner"),
					resource.TestCheckResourceAttr("auth0_role.admin", "name", "admin"),
				),
			},
			{
				Config: testAccUser_update_again,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("auth0_user.user", "roles.#", "1"),
				),
			},
		},
	})
}

const testAccUser_create = `
provider auth0 {}

resource auth0_user user {
  connection_name = "Username-Password-Authentication"
  username = "unique_username"
  user_id = "12345"
  email = "test@test.com"
  password = "passpass$12$12"
  name = "Firstname Lastname"
  given_name = "Firstname"
  family_name = "Lastname"
  nickname = "some.nickname"
  picture = "https://www.example.com/a-valid-picture-url.jpg"
  user_metadata = <<EOF
{
  	"foo": "bar",
  	"bar": { "baz": "qux" }
}
EOF
  app_metadata = <<EOF
{
  	"foo": "bar",
  	"bar": { "baz": "qux" }
}
EOF
}
`

const testAccUser_update = `
provider auth0 {}

resource auth0_user user {
  connection_name = "Username-Password-Authentication"
  username = "unique_username"
  user_id = "12345"
  email = "test@test.com"
  password = "passpass$12$12"
  name = "Firstname Lastname"
  given_name = "Firstname"
  family_name = "Lastname"
  nickname = "some.nickname"
  picture = "https://www.example.com/a-valid-picture-url.jpg"
  roles = [ auth0_role.owner.id, auth0_role.admin.id ]
  user_metadata = <<EOF
{
  	"foo": "bar",
  	"bar": { "baz": "qux" }
}
EOF
  app_metadata = <<EOF
{
  	"foo": "bar",
  	"bar": { "baz": "qux" }
}
EOF
}

resource auth0_role owner {
	name = "owner"
	description = "Owner"
}

resource auth0_role admin {
	name = "admin"
	description = "Administrator"
}
`

const testAccUser_update_again = `
provider auth0 {}

resource auth0_user user {
  connection_name = "Username-Password-Authentication"
  username = "unique_username"
  user_id = "12345"
  email = "test@test.com"
  password = "passpass$12$12"
  name = "Firstname Lastname"
  given_name = "Firstname"
  family_name = "Lastname"
  nickname = "some.nickname"
  picture = "https://www.example.com/a-valid-picture-url.jpg"
  roles = [ auth0_role.admin.id ]
  user_metadata = <<EOF
{
  	"foo": "bar",
  	"bar": { "baz": "qux" }
}
EOF
  app_metadata = <<EOF
{
  	"foo": "bar",
  	"bar": { "baz": "qux" }
}
EOF
}

resource auth0_role admin {
	name = "admin"
	description = "Administrator"
}
`
