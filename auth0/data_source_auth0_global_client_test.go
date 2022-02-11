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

func TestAccDataGlobalClient(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"auth0": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: testAccGlobalClientConfigWithCustomLogin,
			},
			{
				Config: fmt.Sprintf(testAccDataGlobalClientConfig, testAccGlobalClientConfigWithCustomLogin),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.auth0_global_client.global", "custom_login_page", "<html>TEST123</html>"),
					resource.TestCheckResourceAttr("data.auth0_global_client.global", "custom_login_page_on", "true"),
					resource.TestCheckResourceAttrSet("data.auth0_global_client.global", "client_id"),
					resource.TestCheckResourceAttr("data.auth0_global_client.global", "app_type", ""),
					resource.TestCheckResourceAttr("data.auth0_global_client.global", "name", "All Applications"),
					resource.TestCheckNoResourceAttr("data.auth0_global_client.global", "client_secret"),
				),
			},
		},
	})
}
