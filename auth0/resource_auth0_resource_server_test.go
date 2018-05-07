package auth0

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccResourceServer(t *testing.T) {

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"auth0": Provider(),
		},
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccResourceServerConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("auth0_resource_server.my_resource_server", "name", "Test Resource Server"),
				),
			},
		},
	})
}

const testAccResourceServerConfig = `
provider "auth0" {}

resource "auth0_resource_server" "my_resource_server" {
  name = "Test Resource Server"
  identifier = "https://api.example.com/v2"
  signing_alg = "RS256"
  scopes = {
  	value = "create:foo"
  	description = "Create foos"
  }
  scopes = {
  	value = "create:bar"
  	description = "Create bars"
  }
  allow_offline_access = true
  token_lifetime = 8600
  skip_consent_for_verifiable_first_party_clients = true
}
`
