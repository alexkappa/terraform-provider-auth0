package auth0

import (
	"fmt"
	"testing"

	"github.com/alexkappa/terraform-provider-auth0/auth0/internal/random"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

const testAccDataClientConfigInit = `
resource "auth0_client" "test" {
  name = "Acceptance Test - {{.random}}"
  description = "Test Application Long Description"
  app_type = "non_interactive"
  custom_login_page_on = true
  is_first_party = true
  is_token_endpoint_ip_header_trusted = true
  token_endpoint_auth_method = "client_secret_post"
  oidc_conformant = true
  callbacks = [ "https://example.com/callback" ]
  allowed_origins = [ "https://example.com" ]
  allowed_clients = [ "https://allowed.example.com" ]
  grant_types = [ "authorization_code", "http://auth0.com/oauth/grant-type/password-realm", "implicit", "password", "refresh_token" ]
  organization_usage = "deny"
  organization_require_behavior = "no_prompt"
  allowed_logout_urls = [ "https://example.com" ]
  web_origins = [ "https://example.com" ]
  jwt_configuration {
    lifetime_in_seconds = 300
    secret_encoded = true
    alg = "RS256"
    scopes = {
      foo = "bar"
    }
  }
  client_metadata = {
    foo = "zoo"
  }
  addons {
    firebase = {
      client_email = "john.doe@example.com"
      lifetime_in_seconds = 1
      private_key = "wer"
      private_key_id = "qwreerwerwe"
    }
    samlp {
      audience = "https://example.com/saml"
      mappings = {
        email = "http://schemas.xmlsoap.org/ws/2005/05/identity/claims/emailaddress"
        name = "http://schemas.xmlsoap.org/ws/2005/05/identity/claims/name"
      }
      create_upn_claim = false
      passthrough_claims_with_no_mapping = false
      map_unknown_claims_as_is = false
      map_identities = false
      name_identifier_format = "urn:oasis:names:tc:SAML:2.0:nameid-format:persistent"
      name_identifier_probes = [
        "http://schemas.xmlsoap.org/ws/2005/05/identity/claims/emailaddress"
      ]
      logout = {
        callback = "http://example.com/callback"
        slo_enabled = true
      }
	  signing_cert = "-----BEGIN PUBLIC KEY-----\nMIGf...bpP/t3\n+JGNGIRMj1hF1rnb6QIDAQAB\n-----END PUBLIC KEY-----\n"
    }
  }
  refresh_token {
    leeway = 42
    token_lifetime = 424242
    rotation_type = "rotating"
    expiration_type = "expiring"
    infinite_token_lifetime = true
    infinite_idle_token_lifetime = false
    idle_token_lifetime = 3600
  }
  mobile {
    ios {
      team_id = "9JA89QQLNQ"
      app_bundle_identifier = "com.my.bundle.id"
    }
  }
  initiate_login_uri = "https://example.com/login"
}
`

const testAccDataClientConfigByName = `
%v
data auth0_client test {
  name = "Acceptance Test - {{.random}}"
}
`

const testAccDataClientConfigById = `
%v
data auth0_client test {
  client_id = auth0_client.test.client_id
}
`

func TestAccDataClientByName(t *testing.T) {
	rand := random.String(6)

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"auth0": Provider(),
		},
		PreventPostDestroyRefresh: true,
		Steps: []resource.TestStep{
			{
				Config: random.Template(testAccDataClientConfigInit, rand), // must initialize resource before reading with data source
			},
			{
				Config: random.Template(fmt.Sprintf(testAccDataClientConfigByName, testAccDataClientConfigInit), rand),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.auth0_client.test", "client_id"),
					resource.TestCheckResourceAttr("data.auth0_client.test", "name", fmt.Sprintf("Acceptance Test - %v", rand)),
					resource.TestCheckResourceAttr("data.auth0_client.test", "app_type", "non_interactive"),
					resource.TestCheckNoResourceAttr("data.auth0_client.test", "client_secret_rotation_trigger"),
				),
			},
		},
	})
}

func TestAccDataClientById(t *testing.T) {
	rand := random.String(6)

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"auth0": Provider(),
		},
		PreventPostDestroyRefresh: true,
		Steps: []resource.TestStep{
			{
				Config: random.Template(testAccDataClientConfigInit, rand),
			},
			{
				Config: random.Template(fmt.Sprintf(testAccDataClientConfigById, testAccDataClientConfigInit), rand),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.auth0_client.test", "id"),
					resource.TestCheckResourceAttrSet("data.auth0_client.test", "name"),
					resource.TestCheckNoResourceAttr("data.auth0_client.test", "client_secret_rotation_trigger"),
				),
			},
		},
	})
}
