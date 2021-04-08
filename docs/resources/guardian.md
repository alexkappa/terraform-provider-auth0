---
layout: "auth0"
page_title: "Auth0: auth0_guardian"
description: |-
  With this reasource, you can configure some of the MFA options
---

# auth0_guardian

Multi-factor Authentication works by requiring additional factors during the login process to prevent unauthorized access. With this resource you can configure some of
the options available for MFA.

## Example Usage

```hcl
resource "auth0_guardian" "default" {
  policy = "all-applications"
  phone {
    enabled       = false
    provider      = "auth0"
    message_types = ["sms"]
    options {
      enrollment_message   = "{{code}}} is your verification code for {{tenant.friendly_name}}. Please enter this code to verify your enrollment"
      verification_message = "{{code}} is your verification code for {{tenant.friendly_name}}"
    }
  }
}
```

## Argument Reference

Arguments accepted by this resource include:

* `policy` - (Required) String. Policy to use. Available options are `never`, `all-applications` and `confidence-score. The option `confidence-score` means the trigger of MFA will be adaptive. See [Auth0 docs](https://auth0.com/docs/mfa/adaptive-mfa)
* `phone` - (Optional) List(Resource). Configuration settings for the phone MFA. For details, see [Phone](#phone).

### Phone

`phone` supports the following arguments:

* `enabled` - (Required) String. API User for your email service.
* `provider` - (Required) String, Case-sensitive. Provider to use, one of `auth0`, `twilio` or `phone-message-hook`.
* `message_types` - (Required) List(String). Message types to use, array of `phone` and or `voice`. Adding both to array should enable the user to choose.
* `secret_access_key` - (Optional) String, Case-sensitive. AWS Secret Key. Will always be encrypted in our database. Used only for AWS.
* `options`- (Required) List(Resource). Options for the various providers. See [Options](#options).

### Options
`options` supports different arguments depending on the provider specificed in [Phone](#phone).

### Auth0
* `enrollment_message` (Optional) String. This message will be sent whenever a user enrolls a new device for the first time using MFA. Supports liquid syntax, see [Auth0 docs](https://auth0.com/docs/mfa/customize-sms-or-voice-messages).
* `verification_message` (Optional) String. This message will be sent whenever a user logs in after the enrollment. Supports liquid syntax, see [Auth0 docs](https://auth0.com/docs/mfa/customize-sms-or-voice-messages).

### Twilio
* `enrollment_message` (Optional) String. This message will be sent whenever a user enrolls a new device for the first time using MFA. Supports liquid syntax, see [Auth0 docs](https://auth0.com/docs/mfa/customize-sms-or-voice-messages).
* `verification_message` (Optional) String. This message will be sent whenever a user logs in after the enrollment. Supports liquid syntax, see [Auth0 docs](https://auth0.com/docs/mfa/customize-sms-or-voice-messages).
* `sid`(Optional) String.
* `auth_token`(Optional) String.
* `from` (Optional) String.
* `messaging_service_sid`(Optional) String.

### Phone message hook
No options. Custom code has to be written in a phone message hook. See [phone message hook docs](https://auth0.com/docs/hooks/extensibility-points/send-phone-message).
