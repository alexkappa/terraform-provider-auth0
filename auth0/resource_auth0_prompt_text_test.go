package auth0

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccPromptText(t *testing.T) {

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"auth0": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: testAccPromptTextCreate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("auth0_prompt_text.prompt_text", "language", "es"),
					resource.TestCheckResourceAttr("auth0_prompt_text.prompt_text", "prompt_name", "reset_password"),
					// TODO: Validate prompt_text content
				),
			},
		},
	})
}

const testAccPromptTextCreate = `

resource "auth0_prompt_text" "prompt_text" {
  language    = "es"
  prompt_name = "reset_password"

  prompt_content {
    reset_password_request {
	  page_title = "hola"
	  title      = "las vegas"
    }
  }
}
`
