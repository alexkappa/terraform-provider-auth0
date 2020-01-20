package auth0

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func init() {
	resource.AddTestSweepers("auth0_email", &resource.Sweeper{
		Name: "auth0_email",
		F: func(_ string) error {
			api, err := Auth0()
			if err != nil {
				return err
			}
			return api.Email.Delete()
		},
	})
}

func TestAccEmail(t *testing.T) {

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"auth0": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: `
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
				`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("auth0_email.my_email_provider", "name", "ses"),
					resource.TestCheckResourceAttr("auth0_email.my_email_provider", "enabled", "true"),
					resource.TestCheckResourceAttr("auth0_email.my_email_provider", "default_from_address", "accounts@example.com"),
					resource.TestCheckResourceAttr("auth0_email.my_email_provider", "credentials.0.access_key_id", "AKIAXXXXXXXXXXXXXXXX"),
					resource.TestCheckResourceAttr("auth0_email.my_email_provider", "credentials.0.secret_access_key", "7e8c2148xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
					resource.TestCheckResourceAttr("auth0_email.my_email_provider", "credentials.0.region", "us-east-1"),
				),
			},
			{
				Config: `
				resource "auth0_email" "my_email_provider" {
					name = "ses"
					enabled = true
					default_from_address = "accounts@example.com"
					credentials {
						access_key_id = "AKIAXXXXXXXXXXXXXXXY"
						secret_access_key = "7e8c2148xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
						region = "us-east-1"
					}
				}
				`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("auth0_email.my_email_provider", "name", "ses"),
					resource.TestCheckResourceAttr("auth0_email.my_email_provider", "enabled", "true"),
					resource.TestCheckResourceAttr("auth0_email.my_email_provider", "default_from_address", "accounts@example.com"),
					resource.TestCheckResourceAttr("auth0_email.my_email_provider", "credentials.0.access_key_id", "AKIAXXXXXXXXXXXXXXXY"),
					resource.TestCheckResourceAttr("auth0_email.my_email_provider", "credentials.0.secret_access_key", "7e8c2148xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"),
					resource.TestCheckResourceAttr("auth0_email.my_email_provider", "credentials.0.region", "us-east-1"),
				),
			},
			{
				Config: `
				resource "auth0_email" "my_email_provider" {
					name = "mailgun"
					enabled = true
					default_from_address = "accounts@example.com"
					credentials {
						api_key = "MAILGUNXXXXXXXXXXXXXXX"
						domain = "example.com"
						region = "eu"
					}
				}
				`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("auth0_email.my_email_provider", "name", "mailgun"),
					resource.TestCheckResourceAttr("auth0_email.my_email_provider", "enabled", "true"),
					resource.TestCheckResourceAttr("auth0_email.my_email_provider", "default_from_address", "accounts@example.com"),
					resource.TestCheckResourceAttr("auth0_email.my_email_provider", "credentials.0.domain", "example.com"),
					resource.TestCheckResourceAttr("auth0_email.my_email_provider", "credentials.0.region", "eu"),
				),
			},
		},
	})
}
