package auth0

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

const testAccDataGlobalClientConfig = `
%v

data auth0_global_client global {
}
`

func clientGlobalDataSourceFields() []string {
	return []string{
		"client_id",
		"name",
		"app_type",
		"logo_uri",
		"custom_login_page_on",
		"initiate_login_uri",
		"oidc_conformant",
		"sso",
		"cross_origin_auth",
		"description",
		"token_endpoint_auth_method",
		"client_secret",
		"is_first_party",
		"grant_types.0",
		"grant_types.1",
		"grant_types.2",
		"grant_types.3",
		"grant_types.4",
		"grant_types.#",
		"callbacks.0",
		"callbacks.#",
		"web_origins.0",
		"web_origins.#",
		"allowed_origins.0",
		"allowed_origins.#",
		"allowed_logout_urls.0",
		"allowed_logout_urls.#",
	}
}

func TestAccDataGlobalClient(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"auth0": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(testAccDataGlobalClientConfig, testAccGlobalClientConfigWithCustomLogin),
				Check:  checkDataSourceStateMatchesResourceState("data.auth0_global_client.global", "auth0_global_client.global", clientGlobalDataSourceFields()),
			},
		},
	})
}
