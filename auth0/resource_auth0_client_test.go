package auth0

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccClient(t *testing.T) {

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"auth0": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: testAccClientConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("auth0_client.my_client", "name", "Application - Acceptance Test"),
				),
			},
		},
	})
}

const testAccClientConfig = `
provider "auth0" {}

resource "auth0_client" "my_client" {
  name = "Application - Acceptance Test"
  description = "Test Applications Long Description"
  app_type = "non_interactive"
  custom_login_page_on = true
  is_first_party = true
  oidc_conformant = false
  callbacks = [ "https://example.com/callback" ]
  allowed_origins = [ "https://example.com" ]
  grant_types = [ "authorization_code", "http://auth0.com/oauth/grant-type/password-realm", "implicit", "password", "refresh_token" ] 
  allowed_logout_urls = [ "https://example.com" ]
  web_origins = [ "https://example.com" ]
  jwt_configuration = {
    lifetime_in_seconds = 300
    secret_encoded = true
    alg = "RS256"
    scopes = {
    	foo = "bar"
    }
  }
  mobile = {
    ios = {
      team_id = "9JA89QQLNQ"
      app_bundle_identifier = "com.my.bundle.id"
    }
  }
}
`
