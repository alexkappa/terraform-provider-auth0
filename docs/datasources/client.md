---
layout: "auth0"
page_title: "Auth0: auth0_client"
description: |-
  Use this data source to retrieve information about applications that use Auth0 for authentication.
---

# auth0_client

With this resource, you can set up applications that use Auth0 for authentication and configure allowed callback URLs and secrets for these applications. Depending on your plan, you may also configure add-ons to allow your application to call another application's API (such as Firebase and AWS) on behalf of an authenticated user.

## Example Usage

```hcl
data "auth0_client" "my_client" {
  name = "Application - Acceptance Test"
}
```

## Argument Reference

Arguments accepted by this resource include:

* `name` - (Required) String. Name of the client. Make sure to use a unique name if possible. Only unique names are accepted.

## Attribute Reference

Attributes exported by this resource include:

* `name` - String. Name of the client.
* `description` - String, (Max length = 140 characters). Description of the purpose of the client.
* `client_id` - String. ID of the client.
* `client_secret` - String. Secret for the client; keep this private.