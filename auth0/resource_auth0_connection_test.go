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
					resource.TestCheckResourceAttr("auth0_connection.my_connection", "strategy", "waad"),
				),
			},
		},
	})
}

const testAccConnectionConfig = `
provider "auth0" {
	domain = ""
	client_id = ""
	client_secret = ""
}

resource "auth0_connection" "my_connection" {
	name     = "Acceptance-Test-Connection"
  	strategy = "waad"
  	options = {
    	client_id     = "123456"
    	client_secret = "123456"
    	tenant_domain = "example.com"
		domain_aliases = [
			"example.io",
		]
		use_wsfed            = false
		waad_protocol        = "openid-connect"
		waad_common_endpoint = false
		app_domain       = "my-auth0-app.eu.auth0.com"
		api_enable_users = true
		basic_profile    = true
		ext_groups       = true
		ext_profile      = true
	}
}
`
