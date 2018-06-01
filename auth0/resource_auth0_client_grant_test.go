package auth0

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccClientGrant(t *testing.T) {

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"auth0": Provider(),
		},
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccClientGrantConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("auth0_client_grant.my_client_grant", "audience", "https://api.example.com/client-grant-test"),
					resource.TestCheckResourceAttr("auth0_client_grant.my_client_grant", "scope.0", "create:foo"),
				),
			},
		},
	})
}

const testAccClientGrantConfig = `
provider "auth0" {}

resource "auth0_client" "my_client" {
  name = "Application - Client Grant - Acceptance Test"
}

resource "auth0_resource_server" "my_resource_server" {
  name = "Resource Server - Client Grant - Acceptance Test"
  identifier = "https://api.example.com/client-grant-test"
  scopes = {
  	value = "create:foo"
  	description = "Create foos"
  }
  scopes = {
  	value = "create:bar"
  	description = "Create bars"
  }
}

resource "auth0_client_grant" "my_client_grant" {
  client_id = "${auth0_client.my_client.id}"
  audience = "${auth0_resource_server.my_resource_server.identifier}"
  scope = [ "create:foo" ]
}
`
