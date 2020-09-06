package auth0

import (
	"testing"

	"github.com/alexkappa/terraform-provider-auth0/auth0/internal/random"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccClientGrant(t *testing.T) {

	rand := random.String(6)

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"auth0": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: random.Template(testAccClientGrantConfigCreate, rand),
				Check: resource.ComposeTestCheckFunc(
					random.TestCheckResourceAttr("auth0_client_grant.my_client_grant", "audience", "https://uat.tf.alexkappa.com/client-grant/{{.random}}", rand),
					resource.TestCheckResourceAttr("auth0_client_grant.my_client_grant", "scope.#", "0"),
				),
			},
			{
				Config: random.Template(testAccClientGrantConfigUpdate, rand),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("auth0_client_grant.my_client_grant", "scope.#", "1"),
					resource.TestCheckResourceAttr("auth0_client_grant.my_client_grant", "scope.0", "create:foo"),
				),
			},
			{
				Config: random.Template(testAccClientGrantConfigUpdateAgain, rand),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("auth0_client_grant.my_client_grant", "scope.#", "0"),
				),
			},
			{
				Config: random.Template(testAccClientGrantConfigUpdateChangeClient, rand),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("auth0_client_grant.my_client_grant", "scope.#", "0"),
				),
			},
		},
	})
}

const testAccClientGrantAuxConfig = `

resource "auth0_client" "my_client" {
	name = "Acceptance Test - Client Grant - {{.random}}"
	custom_login_page_on = true
	is_first_party = true
}

resource "auth0_resource_server" "my_resource_server" {
	name = "Acceptance Test - Client Grant - {{.random}}"
	identifier = "https://uat.tf.alexkappa.com/client-grant/{{.random}}"
	scopes {
		value = "create:foo"
		description = "Create foos"
	}
	scopes {
		value = "create:bar"
		description = "Create bars"
	}
}
`

const testAccClientGrantConfigCreate = testAccClientGrantAuxConfig + `

resource "auth0_client_grant" "my_client_grant" {
	client_id = "${auth0_client.my_client.id}"
	audience = "${auth0_resource_server.my_resource_server.identifier}"
	scope = [ ]
}
`

const testAccClientGrantConfigUpdate = testAccClientGrantAuxConfig + `

resource "auth0_client_grant" "my_client_grant" {
	client_id = "${auth0_client.my_client.id}"
	audience = "${auth0_resource_server.my_resource_server.identifier}"
	scope = [ "create:foo" ] 
}
`

const testAccClientGrantConfigUpdateAgain = testAccClientGrantAuxConfig + `

resource "auth0_client_grant" "my_client_grant" {
	client_id = "${auth0_client.my_client.id}"
	audience = "${auth0_resource_server.my_resource_server.identifier}"
	scope = [ ]
}
`

const testAccClientGrantConfigUpdateChangeClient = testAccClientGrantAuxConfig + `

resource "auth0_client" "my_client_alt" {
	name = "Acceptance Test - Client Grant Alt - {{.random}}"
	custom_login_page_on = true
	is_first_party = true
}

resource "auth0_client_grant" "my_client_grant" {
	client_id = "${auth0_client.my_client_alt.id}"
	audience = "${auth0_resource_server.my_resource_server.identifier}"
	scope = [ ]
}
`
