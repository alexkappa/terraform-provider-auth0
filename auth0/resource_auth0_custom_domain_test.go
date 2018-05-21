package auth0

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccCustomDomain(t *testing.T) {

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"auth0": Provider(),
		},
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCustomDomain,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("auth0_custom_domain.my_custom_domain", "domain", "auth.example.com"),
					resource.TestCheckResourceAttr("auth0_custom_domain.my_custom_domain", "type", "auth0_managed_certs"),
					resource.TestCheckResourceAttr("auth0_custom_domain.my_custom_domain", "verification_method", "txt"),
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
