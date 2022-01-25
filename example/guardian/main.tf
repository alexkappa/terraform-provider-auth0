provider "auth0" {}

resource "auth0_guardian" "guardian" {
  email  = false
  policy = "all-applications"
  phone {
    provider      = "auth0"
    message_types = ["sms", "voice"]
    options {
      verification_message = "{{code}} is your verification code for {{tenant.friendly_name}}. Please enter this code to verify your enrollment."
      enrollment_message   = "{{code}} is your verification code for {{tenant.friendly_name}}."
    }
  }
}
