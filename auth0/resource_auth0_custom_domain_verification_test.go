package auth0

import (
	"testing"

	"github.com/alexkappa/terraform-provider-auth0/auth0/internal/digitalocean"
	"github.com/alexkappa/terraform-provider-auth0/auth0/internal/random"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccCustomDomainVerification(t *testing.T) {

	rand := random.String(6)

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"auth0":        Provider(),
			"digitalocean": digitalocean.Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: random.Template(testAccCustomDomainVerification, rand),
				Check: resource.ComposeTestCheckFunc(
					random.TestCheckResourceAttr("auth0_custom_domain.my_custom_domain", "domain", "{{.random}}.auth.uat.alexkappa.com", rand),
					resource.TestCheckResourceAttr("auth0_custom_domain.my_custom_domain", "type", "auth0_managed_certs"),
					resource.TestCheckResourceAttrSet("auth0_custom_domain_verification.my_custom_domain_verification", "custom_domain_id"),
				),
			},
			{
				Config: random.Template(testAccCustomDomainVerification, rand),
			},
		},
	})
}

const testAccCustomDomainVerification = `

resource "digitalocean_record" "auth0_domain" {
	domain = "alexkappa.com"
	type = upper(auth0_custom_domain.my_custom_domain.verification[0].methods[0].name)
	name = "{{.random}}.auth.uat.alexkappa.com."
	value = "${auth0_custom_domain.my_custom_domain.verification[0].methods[0].record}."
	ttl = 60
}

resource "auth0_custom_domain" "my_custom_domain" {
	domain = "{{.random}}.auth.uat.alexkappa.com"
	type = "auth0_managed_certs"
	verification_method = "txt"
}

resource "auth0_custom_domain_verification" "my_custom_domain_verification" {
	custom_domain_id = auth0_custom_domain.my_custom_domain.id
	timeouts { create = "15m" }
	depends_on = [ digitalocean_record.auth0_domain ]
}
`
