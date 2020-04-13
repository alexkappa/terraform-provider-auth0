package auth0

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccBranding(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"auth0": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: testAccBrandingConfigCreate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("auth0_branding.my_branding", "logo_url", "https://mycompany.org/logo.png"),
					resource.TestCheckResourceAttr("auth0_branding.my_branding", "favicon_url", "https://mycompany.org/favicon.png"),
					resource.TestCheckResourceAttr("auth0_branding.my_branding", "font.0.url", "https://mycompany.org/font.ttf"),
					resource.TestCheckResourceAttr("auth0_branding.my_branding", "colors.0.primary", "#ffffff"),
					resource.TestCheckResourceAttr("auth0_branding.my_branding", "colors.0.page_background", "#000000"),
				),
			},
			{
				Config: testAccBrandingConfigPageBackgroundGradient,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("auth0_branding.my_branding", "colors.0.page_background_gradient.0.type", "linear-gradient"),
					resource.TestCheckResourceAttr("auth0_branding.my_branding", "colors.0.page_background_gradient.0.start", "#aaaaaa"),
					resource.TestCheckResourceAttr("auth0_branding.my_branding", "colors.0.page_background_gradient.0.end", "#333333"),
					resource.TestCheckResourceAttr("auth0_branding.my_branding", "colors.0.page_background_gradient.0.angle_deg", "35"),
				),
			},
		},
	})

}

const testAccBrandingConfigCreate = `
resource "auth0_branding" "my_branding" {
	logo_url    = "https://mycompany.org/logo.png"
	favicon_url = "https://mycompany.org/favicon.png"

	font {
		url = "https://mycompany.org/font.ttf"
	}

	colors {
		primary         = "#ffffff"
		page_background = "#000000"
	}
}
`

const testAccBrandingConfigPageBackgroundGradient = `
resource "auth0_branding" "my_branding" {
	logo_url    = "https://mycompany.org/logo.png"
	favicon_url = "https://mycompany.org/favicon.png"

	font {
		url = "https://mycompany.org/font.ttf"
	}

	colors {
		primary = "#ffffff"

		page_background_gradient {
			type      = "linear-gradient"
			start     = "#aaaaaa"
			end       = "#333333"
			angle_deg = 35
		}
	}
}
`
