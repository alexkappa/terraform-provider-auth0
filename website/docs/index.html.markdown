---
layout: "auth0"
page_title: "Provider: Auth0"
sidebar_current: "docs-auth0-index"
description: |-
  The Auth0 provider is a provider that leverages the Auth0 Management API to
  configure your Auth0 tenant.
---

# Auth0 Provider

The Auth0 provider is a provider that leverages the Auth0 Management API to 
configure your Auth0 tenant.

## Usage

1- Create a `machine-to-machine` application within Auth0.
2- Give it full access to the Auth0 Management API.

See [this guide](https://auth0.com/docs/api/management/v2/create-m2m-app) for 
details.

Use the client ID and client secret of the application you created for the
credentials as demonstrated below.

### Static Credentials

```hcl
provider "auth0" {
  domain        = "<domain>"
  client_id     = "<client-id>"
  client_secret = "<client-secret>"
}
```

### Environment Variables Credentials

```hcl
provider "auth0" {}
```

Usage:

```sh
$ export AUTH0_DOMAIN="<domain>"
$ export AUTH0_CLIENT_ID="<client-id>"
$ export AUTH0_CLIENT_SECRET="<client-secret>"
$ terraform plan
```

## Argument Reference

In addition to [generic `provider` arguments](https://www.terraform.io/docs/configuration/providers.html)
(e.g. `alias` and `version`), the following arguments are supported in the Auth0
 `provider` block:

* `domain` - (Required) This is the Auth0 tenant domain. Example `mytenant.auth0.com`.
  It must be provided, but it can also be sourced from the `AUTH0_DOMAIN` 
  environment variable.

* `client_id` - (Required) This is the Auth0 application client ID. It must be
  provided, but it can also be sourced from the `AUTH0_DOMAIN` environment variable.

* `client_secret` - (Required) This is the Auth0 application secret. It must be
  provided, but it can also be sourced from the `AUTH0_DOMAIN` environment variable.
