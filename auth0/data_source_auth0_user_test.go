package auth0

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/alexkappa/terraform-provider-auth0/auth0/internal/random"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccUserDataSourceMissingParams(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"auth0": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config:      "data auth0_user user {}",
				ExpectError: regexp.MustCompile(`The argument "user_id" or "email" should be configured`),
			},
		},
	})
}

func testAccDataSourceAuth0User(resourceName, dataSourceName string, testAttributes []string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		user := s.RootModule().Resources[resourceName]
		userResource := user.Primary.Attributes

		search := s.RootModule().Resources[dataSourceName]
		searchResource := search.Primary.Attributes

		if searchResource["user_id"] == "" {
			return fmt.Errorf("Expected to get a user_id from Auth0")
		}

		for _, attribute := range testAttributes {
			if searchResource[attribute] != userResource[attribute] {
				return fmt.Errorf("Expected the user's %s to be: %s, but got: %s", attribute, userResource[attribute], searchResource[attribute])
			}
		}

		return nil
	}
}

func TestAccUserDataSourceById(t *testing.T) {
	rand := random.String(6)

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"auth0": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: random.Template(testAccUserCreateAnDataByID, rand),
				Check: testAccDataSourceAuth0User("auth0_user.user", "data.auth0_user.user",
					[]string{"username", "user_id", "email", "name", "given_name", "family_name", "nickname", "picture", "user_metadata", "identities"}),
			},
		},
	})
}

const testAccUserCreateAnDataByID = `

resource auth0_user user {
	connection_name = "Username-Password-Authentication"
	username = "{{.random}}"
	user_id = "{{.random}}"
	email = "{{.random}}@acceptance.test.com"
	password = "passpass$12$12"
	name = "Firstname Lastname"
	given_name = "Firstname"
	family_name = "Lastname"
	nickname = "{{.random}}"
	picture = "https://www.example.com/picture.jpg"
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

data "auth0_user" "user" {
  user_id                =  auth0_user.user.user_id
}

`

func TestAccUserDataSourceByEmail(t *testing.T) {
	rand := random.String(6)

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"auth0": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: random.Template(testAccUserCreateAnDataByEmail, rand),
				Check: testAccDataSourceAuth0User("auth0_user.user", "data.auth0_user.user",
					[]string{"username", "user_id", "email", "name", "given_name", "family_name", "nickname", "picture", "user_metadata", "identities"}),
			},
		},
	})
}

const testAccUserCreateAnDataByEmail = `

resource auth0_user user {
	connection_name = "Username-Password-Authentication"
	username = "{{.random}}"
	user_id = "{{.random}}"
	email = "{{.random}}@acceptance.test.com"
	password = "passpass$12$12"
	name = "Firstname Lastname"
	given_name = "Firstname"
	family_name = "Lastname"
	nickname = "{{.random}}"
	picture = "https://www.example.com/picture.jpg"
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

data "auth0_user" "user" {
  email                =  auth0_user.user.email
}

`
