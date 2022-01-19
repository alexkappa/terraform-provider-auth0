---
layout: "auth0"
page_title: "Auth0: auth0_prompt_custom_text"
description: |-
    With this resource, you can manage custom texts on your Auth0 prompts.
---

# auth0_prompt_custom_text

With this resource, you can manage custom text on your Auth0 prompts. You can read more about custom texts
[here](https://auth0.com/docs/customize/universal-login-pages/customize-login-text-prompts).

## Example Usage

```hcl
resource "auth0_prompt_custom_text" "example" {
  prompt   = "login"
  language = "en"
  body = jsonencode(
    {
      "login" : {
        "alertListTitle" : "Alerts",
        "buttonText" : "Continue",
        "description" : "Login to",
        "editEmailText" : "Edit",
        "emailPlaceholder" : "Email address",
        "federatedConnectionButtonText" : "Continue with ${connectionName}",
        "footerLinkText" : "Sign up",
        "footerText" : "Don't have an account?",
        "forgotPasswordText" : "Forgot password?",
        "invitationDescription" : "Log in to accept ${inviterName}'s invitation to join ${companyName} on ${clientName}.",
        "invitationTitle" : "You've Been Invited!",
        "logoAltText" : "${companyName}",
        "pageTitle" : "Log in | ${clientName}",
        "passwordPlaceholder" : "Password",
        "separatorText" : "Or",
        "signupActionLinkText" : "${footerLinkText}",
        "signupActionText" : "${footerText}",
        "title" : "Welcome",
        "usernamePlaceholder" : "Username or email address"
      }
    }
  )
}
```

## Argument Reference

The following arguments are supported:

* `prompt` - (Required) The term `prompt` is used to refer to a specific step in the login flow. Options include `login`, `login-id`, `login-password`, `login-email-verification`, `signup`, `signup-id`, `signup-password`, `reset-password`, `consent`, `mfa-push`, `mfa-otp`, `mfa-voice`, `mfa-phone`, `mfa-webauthn`, `mfa-sms`, `mfa-email`, `mfa-recovery-code`, `mfa`, `status`, `device-flow`, `email-verification`, `email-otp-challenge`, `organizations`, `invitation`, `common`
* `language` - (Required) Language of the custom text. Options include `ar`, `bg`, `bs`, `cs`, `da`, `de`, `el`, `en`, `es`, `et`, `fi`, `fr`, `fr-CA`, `fr-FR`, `he`, `hi`, `hr`, `hu`, `id`, `is`, `it`, `ja`, `ko`, `lt`, `lv`, `nb`, `nl`, `pl`, `pt`, `pt-BR`, `pt-PT`, `ro`, `ru`, `sk`, `sl`, `sr`, `sv`, `th`, `tr`, `uk`, `vi`, `zh-CN`, `zh-TW`
* `body` - (Required) JSON containing the custom texts. You can check the options for each prompt [here](https://auth0.com/docs/customize/universal-login-pages/customize-login-text-prompts#prompt-values)

## Import

auth0_prompt_custom_text can be imported using the import command and specifying the prompt and language separated
by *:* , e.g.

```terminal
terraform import auth0_prompt_custom_text.example login:en
```
