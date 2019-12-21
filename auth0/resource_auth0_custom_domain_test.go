package auth0

import (
	"strings"
	"testing"

	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func init() {
	resource.AddTestSweepers("auth0_custom_domain", &resource.Sweeper{
		Name: "auth0_custom_domain",
		F: func(_ string) (err error) {
			api, err := Auth0()
			if err != nil {
				return
			}
			domains, err := api.CustomDomain.List()
			if err != nil {
				return
			}
			for _, domain := range domains {
				if strings.Contains(*domain.Domain, "auth.example.com") {
					if e := api.CustomDomain.Delete(*domain.ID); e != nil {
						multierror.Append(err, e)
					}
				}
			}
			return
		},
	})
}

func TestAccCustomDomain(t *testing.T) {

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"auth0": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: testAccCustomDomain,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("auth0_custom_domain.my_custom_domain", "domain", "auth.example.com"),
					resource.TestCheckResourceAttr("auth0_custom_domain.my_custom_domain", "type", "auth0_managed_certs"),
					resource.TestCheckResourceAttr("auth0_custom_domain.my_custom_domain", "status", "pending_verification"),
				),
			},
		},
	})
}

const testAccCustomDomain = `
provider "auth0" {}

resource "auth0_custom_domain" "my_custom_domain" {
  domain = "auth.example.com"
  type = "auth0_managed_certs"
  verification_method = "txt"
}
`
