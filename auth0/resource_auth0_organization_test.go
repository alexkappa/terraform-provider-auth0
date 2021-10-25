package auth0

import (
	"testing"

	"github.com/alexkappa/terraform-provider-auth0/auth0/internal/random"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccOrganization(t *testing.T) {

	rand := random.String(6)

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"auth0": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: random.Template(testAccOrganizationCreate, rand),
				Check: resource.ComposeTestCheckFunc(
					random.TestCheckResourceAttr("auth0_organization.alexkappa", "name", "alexkappa", rand),
					resource.TestCheckResourceAttr("auth0_organization.alexkappa", "display_name", "alexkappa.com"),
				),
			},
			{
				Config: random.Template(testAccOrganizationUpdate, rand),
				Check: resource.ComposeTestCheckFunc(
					random.TestCheckResourceAttr("auth0_organization.alexkappa", "name", "alexkappa", rand),
					resource.TestCheckResourceAttr("auth0_organization.alexkappa", "display_name", "alexkappa.com"),
					resource.TestCheckResourceAttr("auth0_organization.alexkappa", "branding.#", "1"),
					resource.TestCheckResourceAttr("auth0_organization.alexkappa", "branding.0.logo_url", "https://alexkappa.com/logo.svg"),
					resource.TestCheckResourceAttr("auth0_organization.alexkappa", "branding.0.colors.%", "2"),
					resource.TestCheckResourceAttr("auth0_organization.alexkappa", "branding.0.colors.primary", "#e3e2f0"),
					resource.TestCheckResourceAttr("auth0_organization.alexkappa", "branding.0.colors.page_background", "#e3e2ff"),
				),
			},
		},
	})
}

const testAccOrganizationCreate = `

resource auth0_organization alexkappa {
	name = "alexkappa"
	display_name = "alexkappa.com"
}
`

const testAccOrganizationUpdate = `

resource auth0_organization alexkappa {
	name = "alexkappa"
	display_name = "alexkappa.com"
	branding {
		logo_url = "https://alexkappa.com/logo.svg"
		colors = {
			primary = "#e3e2f0"
			page_background = "#e3e2ff"
		}
	}
}
`
