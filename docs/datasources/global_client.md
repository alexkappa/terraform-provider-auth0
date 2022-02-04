---
layout: "auth0"
page_title: "Data Source: auth0_global_client"
description: |-
Retrieves a tenant's global Auth0 Application client
---

# Data Source: auth0_global_client

Retrieves a tenant's global Auth0 Application client

## Example Usage

```hcl
data "auth0_global_client" "global" {
}
```

## Argument Reference

No arguments accepted.

## Attribute Reference

* `client_id` - String. ID of the client.
* `client_secret`<sup>[1](#client-keys)</sup> - String. Secret for the client; keep this private.
* `custom_login_page_on` - Boolean. Indicates whether or not a custom login page is to be used.
* `custom_login_page` - String. Content of the custom login page.
* `client_metadata` - (Optional) Map(String)

### Client Keys

To access the `client_secret` attribute you need to add the `read:client_keys` scope to the Terraform client. Otherwise, the attribute will contain an empty string.