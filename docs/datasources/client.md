---
layout: "auth0"
page_title: "Data Source: auth0_client"
description: |-
Data source to retrieve a specific Auth0 Application client by 'client_id' or 'name'
---

# Data Source: auth0_client

Data source to retrieve a specific Auth0 Application client by 'client_id' or 'name'

## Example Usage

```hcl
data "auth0_client" "some-client-by-name" {
  name = "Name of my Application"
}
data "auth0_client" "some-client-by-id" {
  client_id = "abcdefghkijklmnopqrstuvwxyz0123456789"
}
```

## Argument Reference

At least one of the following arguments required:

- `client_id` - (Optional) String. client_id of the application.
- `name` - (Optional) String. Name of the application. Ignored if `client_id` is also specified.

## Attribute Reference

The client data source possesses the same attributes as the `auth0_client` resource, with the exception of `client_secret_rotation_trigger`. Refer to the [auth0_client resource documentation](../resources/client.md) for a list of returned attributes.