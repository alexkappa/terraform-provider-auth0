---
layout: "auth0"
page_title: "Provider: Auth0"
description: |-
  The Auth0 provider is used to interact with Auth0 applications and APIs.
---

# Auth0 Provider

The Auth0 provider is used to interact with Auth0 applications and APIs. It provides resources that allow you to create and manage clients, resource servers, client grants, connections, email providers and templates, rules and rule variables, users, roles, tenants, and custom domains as part of a Terraform deployment.

Use the navigation to the left to read about the available resources.

## Example Usage

```hcl
provider "auth0" {
  domain = "<domain>"
  client_id = "<client-id>"
  client_secret = "<client-secret>"
  debug = "<debug>"
}
```

~> Hard-coding credentials into any Terraform configuration is not recommended, and risks secret leakage should this file ever be committed to a public version control system. See [Environment Variables](#environment-variables) for a better alternative.

## Argument Reference

* `domain` - (Required) Your Auth0 domain name. It can also be sourced from the `AUTH0_DOMAIN` environment variable.
* `client_id` - (Optional) Your Auth0 client ID. It can also be sourced from the `AUTH0_CLIENT_ID` environment variable.
* `client_secret` - (Optional) Your Auth0 client secret. It can also be sourced from the `AUTH0_CLIENT_SECRET` environment variable.
* `api_token` - (Optional) Your Auth0 [management api access token](https://auth0.com/docs/security/tokens/access-tokens/management-api-access-tokens).
  It can also be sourced from the `AUTH0_API_TOKEN` environment variable. Can be
  used instead of `client_id` + `client_secret`. If both are specified,
  `management_token` will be used over `client_id` + `client_secret` fields.
* `debug` - (Optional) Indicates whether or not to turn on debug mode.

## Environment Variables

You can provide your credentials via the `AUTH0_DOMAIN`, `AUTH0_CLIENT_ID` and `AUTH0_CLIENT_SECRET` environment variables, respectively.

```hcl
provider "auth0" {}
```

Usage:

```bash
$ export AUTH0_DOMAIN="<domain>"
$ export AUTH0_CLIENT_ID="<client-id>"
$ export AUTH0_CLIENT_SECRET="<client_secret>"
$ terraform plan
```

## Importing resources

To import Auth0 resources, you will need to know their id. You can use the [Auth0 API Explorer](https://auth0.com/docs/api/management/v2) to easily find your resource id.

