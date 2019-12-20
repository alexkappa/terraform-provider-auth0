package auth0

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccEmailTemplate(t *testing.T) {

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"auth0": Provider(),
		},
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccEmailTemplateConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("auth0_email_template.my_email_template", "template", "welcome_email"),
					resource.TestCheckResourceAttr("auth0_email_template.my_email_template", "body", "<html><body><h1>Welcome!</h1></body></html>"),
					resource.TestCheckResourceAttr("auth0_email_template.my_email_template", "from", "welcome@example.com"),
					resource.TestCheckResourceAttr("auth0_email_template.my_email_template", "result_url", "https://example.com/welcome"),
					resource.TestCheckResourceAttr("auth0_email_template.my_email_template", "subject", "Welcome"),
					resource.TestCheckResourceAttr("auth0_email_template.my_email_template", "syntax", "liquid"),
					resource.TestCheckResourceAttr("auth0_email_template.my_email_template", "url_lifetime_in_seconds", "3600"),
					resource.TestCheckResourceAttr("auth0_email_template.my_email_template", "enabled", "true"),
				),
			},
		},
	})
}

const testAccEmailTemplateConfig = `
provider "auth0" {}

resource "auth0_email" "my_email_provider" {
	name = "ses"
	enabled = true
	default_from_address = "accounts@example.com"
	credentials {
		access_key_id = "AKIAXXXXXXXXXXXXXXXX"
		secret_access_key = "7e8c2148xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
		region = "us-east-1"
	}
}

resource "auth0_email_template" "my_email_template" {
	template = "welcome_email"
	body = "<html><body><h1>Welcome!</h1></body></html>"
	from = "welcome@example.com"
	result_url = "https://example.com/welcome"
	subject = "Welcome"
	syntax = "liquid"
	url_lifetime_in_seconds = 3600
	enabled = true

	depends_on = ["auth0_email.my_email_provider"]
}
`
