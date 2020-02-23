package auth0

import (
	"errors"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccGlobalClient(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"auth0": Provider(),
		},
		Steps: []resource.TestStep{
			{
				/* Update the custom login page and its enabled flag */
				Config: testAccGlobalClientConfigWithCustomLogin,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("auth0_global_client.global", "client_id"),
					resource.TestCheckResourceAttrSet("auth0_global_client.global", "client_secret"),
					resource.TestCheckResourceAttr("auth0_global_client.global", "custom_login_page", "<html>TEST123</html>"),
					resource.TestCheckResourceAttr("auth0_global_client.global", "custom_login_page_on", "true"),
				),
			},
			{
				/* Delete the resource which should cause no change in auth0 configuration as the delete operation is a no op */
				Config: testAccGlobalClientConfigEmpty,
				Check: resource.ComposeTestCheckFunc(
					func(state *terraform.State) error {
						for _, m := range state.Modules {
							if len(m.Resources) > 0 {
								if _, ok := m.Resources["auth0_global_client.global"]; ok {
									return errors.New("auth0_global_client.global exists when it should have been removed")
								}
							}
						}
						return nil
					},
				),
			},
			{
				/* Create the resource again with no config settings and verify the values are set to what they were in step 2*/
				Config:             testAccGlobalClientConfigDefault,
				ExpectNonEmptyPlan: true,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("auth0_global_client.global", "custom_login_page", "<html>TEST123</html>"),
					resource.TestCheckResourceAttr("auth0_global_client.global", "custom_login_page_on", "true"),
				),
			},

			{
				/* Disable custom login */
				Config: testAccGlobalClientConfigNoCustomLogin,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("auth0_global_client.global", "custom_login_page_on", "false"),
				),
			},
		},
	})
}

const testAccGlobalClientConfigEmpty = `
`

const testAccGlobalClientConfigDefault = `
resource "auth0_global_client" "global" {
}
`

const testAccGlobalClientConfigWithCustomLogin = `
resource "auth0_global_client" "global" {
    custom_login_page = "<html>TEST123</html>"
    custom_login_page_on = true
}
`

const testAccGlobalClientConfigNoCustomLogin = `
resource "auth0_global_client" "global" {
    custom_login_page_on = false
}
`
