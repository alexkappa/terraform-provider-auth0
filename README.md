Terraform Provider [![wercker status](https://app.wercker.com/status/a660ce322102e732b9c53a690f6c7078/s/master "wercker status")](https://app.wercker.com/project/byKey/a660ce322102e732b9c53a690f6c7078)
==================

- Website: https://www.terraform.io
- Mailing list: [Google Groups](http://groups.google.com/group/terraform-tool)

<img src="https://cdn.rawgit.com/hashicorp/terraform-website/master/content/source/assets/images/logo-hashicorp.svg" width="600px">

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
----------------------

To use the provider define the `auth0` provider in your `*.tf` file.

```
provider "auth0" {
  "domain" = "<doman>"
  "client_id" = "<client-id>"
  "client_secret" = "<client-secret>"
}
```

These variables can also be accessed via the `AUTH0_DOMAIN`, `AUTH0_CLIENT_ID` and `AUTH0_CLIENT_SECRET` environment variables respectively.

Then you can define Auth0 resources using terraform

```
resource "auth0_client" "my_app_client" {
  name = "My Application (Managed by Terraform)"
  description = "My Applications Long Description"
  app_type = "non_interactive"
  is_first_party = false
  oidc_conformant = false
  callbacks = [ "https://example.com/callback" ]
  allowed_origins = [ "https://example.com" ]
  web_origins = [ "https://example.com" ]
  jwt_configuration = {
    lifetime_in_seconds = 300
    secret_encoded = true
    alg = "RS256"
  }
}
```

Currently this provider supports [Clients (aka Applications)](https://auth0.com/docs/api/management/v2#!/Clients/get_clients) and [Resource Servers (aka APIs)](https://auth0.com/docs/api/management/v2#!/Resource_Servers/get_resource_servers) but we intend to support all entities of the Auth0 Management API.

If you need resources that are not available yet, please help the project by contributing.

Developing the Provider
---------------------------

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (version 1.10+ is *required*). You'll also need to correctly setup a [GOPATH](http://golang.org/doc/code.html#GOPATH), as well as adding `$GOPATH/bin` to your `$PATH`.

To compile the provider, run `make install`. This will build the provider and install the provider binary in the `$GOPATH/bin` directory.

```sh
$ make install
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
