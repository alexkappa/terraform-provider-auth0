---
layout: "auth0"
page_title: "Auth0: auth0_client"
description: |-
With this resource, you can create and configure the Auth0 global client which is used to specify the custom login page for the Universal Login.
---

# auth0_client

With this resource, you can create and configure the Auth0 global client which is used to specify the custom login page for the Universal Login.

## Example Usage

```hcl
resource "auth0_global_client" "my_client" {
  custom_login_page_on = true
  custom_login_page = "--custom universal login page HTML--"
  client_metadata = {
    foo = "zoo"
  }
}
```

## Argument Reference

Arguments accepted by this resource include:

* `custom_login_page_on` - (Optional) Boolean. Indicates whether or not a custom login page is to be used.
* `custom_login_page` - (Optional) String. Content of the custom login page.
* `custom_login_page_preview` - (Optional) String.
* `client_metadata` - (Optional) Map(String)

## Attribute Reference

Attributes exported by this resource include:

* `client_id` - String. ID of the client.
* `client_secret`<sup>[1](#client-keys)</sup> - String. Secret for the global client; keep this private. This is extremely sensitive as it gives unfettered access to your Auth0 client.
* `custom_login_page_on` - Boolean. Indicates whether or not a custom login page is to be used.

### Client keys

To access the `client_secret` attribute you need to add the `read:client_keys` scope to the Terraform client. Otherwise, the attribute will contain an empty string.
