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
					resource.TestCheckResourceAttr("auth0_client.my_client", "addons.0.firebase.client_email", "john.doe@example.com"),
					resource.TestCheckResourceAttr("auth0_client.my_client", "addons.0.firebase.lifetime_in_seconds", "1"),
					resource.TestCheckResourceAttr("auth0_client.my_client", "addons.0.samlp.0.audience", "https://example.com/saml"),
					resource.TestCheckResourceAttr("auth0_client.my_client", "addons.0.samlp.0.map_identities", "false"),
					resource.TestCheckResourceAttr("auth0_client.my_client", "addons.0.samlp.0.name_identifier_format", "urn:oasis:names:tc:SAML:2.0:nameid-format:persistent"),
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
  addons = {
    firebase = {
      client_email = "john.doe@example.com"
      lifetime_in_seconds = 1
      private_key = "wer"
      private_key_id = "qwreerwerwe"
    },
    samlp = {
      audience = "https://example.com/saml",
      mappings = {
        email = "http://schemas.xmlsoap.org/ws/2005/05/identity/claims/emailaddress",
        name = "http://schemas.xmlsoap.org/ws/2005/05/identity/claims/name"
      },
      create_upn_claim = false,
      passthrough_claims_with_no_mapping = false,
      map_unknown_claims_as_is = false,
      map_identities = false,
      name_identifier_format = "urn:oasis:names:tc:SAML:2.0:nameid-format:persistent",
      name_identifier_probes = [
        "http://schemas.xmlsoap.org/ws/2005/05/identity/claims/emailaddress"
      ]
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
