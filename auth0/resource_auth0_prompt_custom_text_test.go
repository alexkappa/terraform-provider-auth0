package auth0

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccPromptCustomText(t *testing.T) {

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"auth0": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: testAccPromptCustomTextCreate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("auth0_prompt_custom_text.prompt_custom_text", "prompt", "login"),
					resource.TestCheckResourceAttr("auth0_prompt_custom_text.prompt_custom_text", "login", "en"),
					resource.TestCheckResourceAttr("auth0_prompt_custom_text.prompt_custom_text", "body", "{\"login\": { \"alertListTitle\": \"Alerts\", \"buttonText\": \"Continue\"}}"),
				),
			},
			{
				Config: testAccPromptCustomTextUpdate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("auth0_prompt_custom_text.prompt_custom_text", "prompt", "login"),
					resource.TestCheckResourceAttr("auth0_prompt_custom_text.prompt_custom_text", "login", "en"),
					resource.TestCheckResourceAttr("auth0_prompt_custom_text.prompt_custom_text", "body", "{\"login\": { \"alertListTitle\": \"Alerts\", \"buttonText\": \"Continue to Login\"}}"),
				),
			},
		},
	})
}

const testAccPromptCustomTextCreate = `

resource "auth0_prompt_custom_text" "prompt_custom_text" {
  prompt = "login"
  language = "en"
  body = "{\"login\": { \"alertListTitle\": \"Alerts\", \"buttonText\": \"Continue\"}}""
}
`

const testAccPromptCustomTextUpdate = `

resource "auth0_prompt_custom_text" "prompt_custom_text" {
	prompt = "login"
	language = "en"
	body = "{\"login\": { \"alertListTitle\": \"Alerts\", \"buttonText\": \"Continue to Login\"}}""
  }
  `
