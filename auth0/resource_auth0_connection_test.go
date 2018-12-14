package auth0

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccConnection(t *testing.T) {

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"auth0": Provider(),
		},
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccConnectionConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("auth0_connection.my_connection", "name", "Acceptance-Test-Connection"),
					resource.TestCheckResourceAttr("auth0_connection.my_connection", "is_domain_connection", "true"),
					resource.TestCheckResourceAttr("auth0_connection.my_connection", "strategy", "auth0"),
					resource.TestCheckResourceAttr("auth0_connection.my_connection", "options.0.password_policy", "fair"),
					resource.TestCheckResourceAttr("auth0_connection.my_connection", "options.0.enabled_database_customization", "false"),
					resource.TestCheckResourceAttr("auth0_connection.my_connection", "options.0.brute_force_protection", "true"),
					resource.TestCheckResourceAttr("auth0_connection.my_connection", "options.0.import_mode", "true"),
					resource.TestCheckResourceAttr("auth0_connection.my_connection", "options.0.disable_signup", "false"),
					resource.TestCheckResourceAttr("auth0_connection.my_connection", "options.0.requires_username", "true"),
					resource.TestCheckResourceAttr("auth0_connection.my_connection", "options.0.custom_scripts.get_user", "myFunction"),
					resource.TestCheckResourceAttrSet("auth0_connection.my_connection", "options.0.configuration.foo"),
				),
			},
		},
	})
}

const testAccConnectionConfig = `
provider "auth0" {}

resource "auth0_connection" "my_connection" {
	name = "Acceptance-Test-Connection"
	is_domain_connection = true
	strategy = "auth0"
	options = {
		password_policy = "fair"
		password_history = {
			enable = "true"
			size = "5"
		}
		enabled_database_customization = false
		brute_force_protection = true
		import_mode = true
		disable_signup = false
		requires_username = true
		custom_scripts = {
			get_user = "myFunction"
		}
		configuration = {
			foo = "bar"
		}
	}
}
`

func TestAccConnectionAd(t *testing.T) {

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"auth0": Provider(),
		},
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccConnectionAdConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("auth0_connection.my_ad_connection", "name", "Acceptance-Test-Ad-Connection"),
					resource.TestCheckResourceAttr("auth0_connection.my_ad_connection", "strategy", "ad"),
				),
			},
		},
	})
}

const testAccConnectionAdConfig = `
provider "auth0" {}

resource "auth0_connection" "my_ad_connection" {
	name = "Acceptance-Test-Ad-Connection"
	strategy = "ad"
}
`
