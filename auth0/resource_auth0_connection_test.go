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

func TestAccConnectionValidation(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"auth0": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: testAccConnectionValidationConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("auth0_connection.my_connection_with_validation", "name", "Acceptance-Test-Connection-With-Validation"),
					resource.TestCheckResourceAttr("auth0_connection.my_connection_with_validation", "strategy", "auth0"),
					resource.TestCheckResourceAttr("auth0_connection.my_connection_with_validation", "options.0.requires_username", "true"),
					resource.TestCheckResourceAttr("auth0_connection.my_connection_with_validation", "options.0.validation.0.username.0.min", "2"),
					resource.TestCheckResourceAttr("auth0_connection.my_connection_with_validation", "options.0.validation.0.username.0.max", "10"),
				),
			},
		},
	})
}

const testAccConnectionValidationConfig = `
provider "auth0" {}

resource "auth0_connection" "my_connection_with_validation" {
	name = "Acceptance-Test-Connection-With-Validation"
	strategy = "auth0"

	options = {
		requires_username = true

		validation {
			username {
				min = 2
				max = 10
			}
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
