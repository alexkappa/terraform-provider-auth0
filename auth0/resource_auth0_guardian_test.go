package auth0

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"testing"
)

func TestAccGuardian(t *testing.T) {

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"auth0": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: testAccConfigureCustomPhone,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("auth0_guardian.foo", "policy", "all-applications"),
					resource.TestCheckResourceAttr("auth0_guardian.foo", "phone.0.enabled", "true"),
					resource.TestCheckResourceAttr("auth0_guardian.foo", "phone.0.message_types.0", "voice"),
					resource.TestCheckResourceAttr("auth0_guardian.foo", "phone.0.provider", "phone-message-hook"),
				),
			},

			{
				Config: testAccConfigureTwilio,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("auth0_guardian.foo", "policy", "all-applications"),
					resource.TestCheckResourceAttr("auth0_guardian.foo", "phone.0.enabled", "true"),
					resource.TestCheckResourceAttr("auth0_guardian.foo", "phone.0.message_types.0", "voice"),
					resource.TestCheckResourceAttr("auth0_guardian.foo", "phone.0.provider", "twilio"),
					resource.TestCheckResourceAttr("auth0_guardian.foo", "phone.0.options.0.enrollment_message", "enroll foo"),
					resource.TestCheckResourceAttr("auth0_guardian.foo", "phone.0.options.0.verification_message", "verify foo"),
					resource.TestCheckResourceAttr("auth0_guardian.foo", "phone.0.options.0.from", "from bar"),
					resource.TestCheckResourceAttr("auth0_guardian.foo", "phone.0.options.0.messaging_service_sid", "foo"),
					resource.TestCheckResourceAttr("auth0_guardian.foo", "phone.0.options.0.auth_token", "bar"),
				),
			},
			{
				Config: testAccConfigureAuth0,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("auth0_guardian.foo", "policy", "all-applications"),
					resource.TestCheckResourceAttr("auth0_guardian.foo", "phone.0.enabled", "true"),
					resource.TestCheckResourceAttr("auth0_guardian.foo", "phone.0.message_types.0", "voice"),
					resource.TestCheckResourceAttr("auth0_guardian.foo", "phone.0.provider", "auth0"),
					resource.TestCheckResourceAttr("auth0_guardian.foo", "phone.0.options.0.enrollment_message", "enroll foo"),
					resource.TestCheckResourceAttr("auth0_guardian.foo", "phone.0.options.0.verification_message", "verify foo"),
				),
			},
			{
				Config: testAccConfigureNoPhone,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("auth0_guardian.foo", "policy", "all-applications"),
					resource.TestCheckNoResourceAttr("auth0_guardian.foo", "phone"),
				),
			},
		},
	})
}

const testAccConfigureCustomPhone = `

resource "auth0_guardian" "foo" {
  policy = "all-applications"
  phone {
	enabled = true
	provider = "phone-message-hook"
	message_types = ["voice"]
}
}
`
const testAccConfigureTwilio = `

resource "auth0_guardian" "foo" {
  policy = "all-applications"
  phone {
	enabled = true
	provider = "twilio"
	message_types = ["voice"]
	options {
		enrollment_message = "enroll foo"
		verification_message = "verify foo"
		from = "from bar"
		messaging_service_sid = "foo"
		auth_token = "bar"
		sid = "foo"
	}
}
}
`

const testAccConfigureAuth0 = `

resource "auth0_guardian" "foo" {
  policy = "all-applications"
  phone {
	enabled = true
	provider = "auth0"
	message_types = ["voice"]
	options {
		enrollment_message = "enroll foo"
		verification_message = "verify foo"
	}
}
}
`
const testAccConfigureNoPhone = `

resource "auth0_guardian" "foo" {
  policy = "all-applications"
}
`
