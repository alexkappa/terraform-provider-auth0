package auth0

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccCustomDomainVerification(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"auth0": providerWithTestingConfiguration(),
		},
		Steps: []resource.TestStep{
			{
				Config: testAccCustomDomainVerification,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("auth0_custom_domain.my_custom_domain", "domain", "terraform-provider.auth0.com"),
					resource.TestCheckResourceAttr("auth0_custom_domain.my_custom_domain", "type", "auth0_managed_certs"),
					resource.TestCheckResourceAttrSet("auth0_custom_domain_verification.my_custom_domain_verification", "custom_domain_id"),
				),
			},
		},
	})
}

const testAccCustomDomainVerification = `
resource "auth0_custom_domain" "my_custom_domain" {
	domain = "terraform-provider.auth0.com"
	type = "auth0_managed_certs"
}

resource "auth0_custom_domain_verification" "my_custom_domain_verification" {
	custom_domain_id = auth0_custom_domain.my_custom_domain.id
	timeouts { create = "15m" }
}
`
