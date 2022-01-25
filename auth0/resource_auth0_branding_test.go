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
					resource.TestCheckResourceAttr("auth0_branding.my_brand", "logo_url", "https://mycompany.org/v1/logo.png"),
					resource.TestCheckResourceAttr("auth0_branding.my_brand", "favicon_url", "https://mycompany.org/favicon.ico"),
					resource.TestCheckResourceAttr("auth0_branding.my_brand", "colors.0.primary", "#0059d6"),
					resource.TestCheckResourceAttr("auth0_branding.my_brand", "colors.0.page_background", "#000000"),
					resource.TestCheckResourceAttr("auth0_branding.my_brand", "font.0.url", "https://mycompany.org/font/myfont.ttf"),
				),
			},
			{
				Config: testAccBrandingConfigUpdate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("auth0_branding.my_brand", "logo_url", "https://mycompany.org/v2/logo.png"),
					resource.TestCheckResourceAttr("auth0_branding.my_brand", "favicon_url", "https://mycompany.org/favicon.ico"),
					resource.TestCheckResourceAttr("auth0_branding.my_brand", "colors.0.primary", "#ffa629"),
					resource.TestCheckResourceAttr("auth0_branding.my_brand", "colors.0.page_background", "#ffffff"),
					resource.TestCheckResourceAttr("auth0_branding.my_brand", "font.0.url", "https://mycompany.org/font/myfont.ttf"),
				),
			},
		},
	})
}

const testAccBrandingConfigCreate = `
resource "auth0_branding" "my_brand" {
	logo_url = "https://mycompany.org/v1/logo.png"
	favicon_url = "https://mycompany.org/favicon.ico"
	colors {
		primary = "#0059d6"
		page_background = "#000000"
	}
	font {
		url = "https://mycompany.org/font/myfont.ttf"
	}
}
`

const testAccBrandingConfigUpdate = `
resource "auth0_branding" "my_brand" {
	logo_url = "https://mycompany.org/v2/logo.png"
	favicon_url = "https://mycompany.org/favicon.ico"
	colors {
		primary = "#ffa629"
		page_background = "#ffffff"
	}
	font {
		url = "https://mycompany.org/font/myfont.ttf"
	}
}
`
