Auth0 Terraform Provider
========================

[![Build Status](https://travis-ci.org/alexkappa/terraform-provider-auth0.svg?branch=master)](https://travis-ci.org/alexkappa/terraform-provider-auth0)
[![Maintainability](https://api.codeclimate.com/v1/badges/9c49c10286123b716c79/maintainability)](https://codeclimate.com/github/alexkappa/terraform-provider-auth0/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/9c49c10286123b716c79/test_coverage)](https://codeclimate.com/github/alexkappa/terraform-provider-auth0/test_coverage)

Requirements
------------

-	[Terraform](https://www.terraform.io/downloads.html) 0.11.x
-	[Go](https://golang.org/doc/install) 1.10 (to build the provider plugin)

Building The Provider
---------------------

Clone repository to: `$GOPATH/src/github.com/yieldr/terraform-provider-auth0`

```sh
$ mkdir -p $GOPATH/src/github.com/yieldr; cd $GOPATH/src/github.com/yieldr
$ git clone git@github.com:yieldr/terraform-provider-auth0
```

Enter the provider directory and build the provider

```sh
$ cd $GOPATH/src/github.com/yieldr/terraform-provider-auth0
$ make build
```

Using the provider
------------------

The provider isn't listed in the official Terraform repository, so using `terraform init` to download the provide won't work. To install the auth0 provider, you can [download the binary](https://github.com/yieldr/terraform-provider-auth0/releases) and place in the directory `~/.terraform.d/plugins` (or `%APPDATA%/terraform.d/plugins/` if you're on Windows).

To use the provider define the `auth0` provider in your `*.tf` file.

```
provider "auth0" {
  "domain" = "<domain>"
  "client_id" = "<client-id>"
  "client_secret" = "<client-secret>"
}
```

These variables can also be accessed via the `AUTH0_DOMAIN`, `AUTH0_CLIENT_ID` and `AUTH0_CLIENT_SECRET` environment variables respectively.

Examples of resources can be found in the [examples directory](example/). The currently supported Auth0 resources are described below.

- [x] [Clients (Applications)](https://auth0.com/docs/api/management/v2#!/Clients/get_clients)
- [x] [Client Grants](https://auth0.com/docs/api/management/v2#!/Client_Grants/get_client_grants)
- [x] [Connections](https://auth0.com/docs/api/management/v2#!/Connections/get_connections)
- [x] [Custom Domains](https://auth0.com/docs/api/management/v2#!/Custom_Domains/get_custom_domains)
- [ ] [Device Credentials](https://auth0.com/docs/api/management/v2#!/Device_Credentials/get_device_credentials)
- [ ] [Grants](https://auth0.com/docs/api/management/v2#!/Grants/get_grants)
- [x] [Resource Servers (APIs)](https://auth0.com/docs/api/management/v2#!/Resource_Servers/get_resource_servers)
- [x] [Rules](https://auth0.com/docs/api/management/v2#!/Rules/get_rules)
- [x] [Rules Configs](https://auth0.com/docs/api/management/v2#!/Rules_Configs/get_rules_configs)
- [ ] [User Blocks](https://auth0.com/docs/api/management/v2#!/User_Blocks/get_user_blocks)
- [ ] [Users](https://auth0.com/docs/api/management/v2#!/Users/get_users)
- [ ] [Users By Email](https://auth0.com/docs/api/management/v2#!/Users_By_Email/get_users_by_email)
- [ ] [Blacklists](https://auth0.com/docs/api/management/v2#!/Blacklists/get_tokens)
- [x] [Email Templates](https://auth0.com/docs/api/management/v2#!/Email_Templates/get_email_templates_by_templateName)
- [x] [Emails](https://auth0.com/docs/api/management/v2#!/Emails/get_provider)
- [ ] [Guardian](https://auth0.com/docs/api/management/v2#!/Guardian/get_factors)
- [ ] [Jobs](https://auth0.com/docs/api/management/v2#!/Jobs/get_jobs_by_id)
- [X] [Tenants](https://auth0.com/docs/api/management/v2#!/Tenants/get_settings)
  - Missing attributes: `flags`, `universal_login`, `idle_session_lifetime`.
- [ ] [Tickets](https://auth0.com/docs/api/management/v2#!/Tickets/post_email_verification)

Developing the Provider
---------------------------

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (version 1.10+ is *required*). You'll also need to correctly setup a [GOPATH](http://golang.org/doc/code.html#GOPATH), as well as adding `$GOPATH/bin` to your `$PATH`.

On how to develop custom terraform providers, read the [official guide](https://www.terraform.io/docs/extend/writing-custom-providers.html).

To compile the provider, run `make build`. This will build the provider and install the provider binary in the `$GOPATH/bin` directory.

```sh
$ make build
...
$ $GOPATH/bin/terraform-provider-auth0
...
```

In order to test the provider, you can simply run `make test`.

```sh
$ make test
```

In order to run the full suite of Acceptance tests, run `make testacc`.

*Note:* Acceptance tests create real resources, and often cost money to run.

```sh
$ make testacc
```
