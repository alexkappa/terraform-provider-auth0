---
layout: "auth0"
page_title: "Auth0: auth0_prompt"
description: |-
  With this resource, you can manage your Auth0 prompts, including choosing the login experience version.
---

# auth0_prompt

With this resource, you can manage your Auth0 prompts, including choosing the login experience version.

## Example Usage

```
resource "auth0_prompt" "example" {
  universal_login_experience = "classic"
}
```

## Argument Reference

The following arguments are supported:

- `universal_login_experience` - (Optional) Which login experience to use. Options include `classic` and `new`.
