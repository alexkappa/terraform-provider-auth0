package auth0

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccBrand(t *testing.T) {

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"auth0": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: testAccBrandConfigCreate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("auth0_brand.my_brand", "logo_url", "https://mycompany.org/logo.png"),
					resource.TestCheckResourceAttr("auth0_brand.my_brand", "favicon_url", "https://mycompany.org/favicon.ico"),
					resource.TestCheckResourceAttr("auth0_brand.my_brand", "colors.0.primary", "#0059d6"),
					resource.TestCheckResourceAttr("auth0_brand.my_brand", "colors.0.page_background", "#000000"),
					resource.TestCheckResourceAttr("auth0_brand.my_brand", "font.0.url", "https://mycompany.org/font/myfont.ttf"),
				),
			},
			{
				Config: testAccBrandConfigUpdate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("auth0_brand.my_brand", "logo_url", "https://mycompany.org/logo.png"),
					resource.TestCheckResourceAttr("auth0_brand.my_brand", "favicon_url", "https://mycompany.org/favicon.ico"),
					resource.TestCheckResourceAttr("auth0_brand.my_brand", "colors.0.primary", "#0059d6"),
					resource.TestCheckResourceAttr("auth0_brand.my_brand", "colors.0.page_background", "#ffffff"),
					resource.TestCheckResourceAttr("auth0_brand.my_brand", "font.0.url", "https://mycompany.org/font/myfont.ttf"),
				),
			},
		},
	})
}

const testAccBrandConfigCreate = `
resource "auth0_brand" "my_brand" {
	logo_url = "https://mycompany.org/logo.png"
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

const testAccBrandConfigUpdate = `
resource "auth0_brand" "my_brand" {
	logo_url = "https://mycompany.org/logo.png"
	favicon_url = "https://mycompany.org/favicon.ico"
	colors {
		primary = "#0059d6"
		page_background = "#ffffff"
	}
	font {
		url = "https://mycompany.org/font/myfont.ttf"
	}
}
`
