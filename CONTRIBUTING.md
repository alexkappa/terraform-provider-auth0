# Contributing

Before you begin, read through the Terraform documentation on [Extending
Terraform](https://www.terraform.io/docs/extend/index.html) and [Writing Custom
Providers](https://learn.hashicorp.com/collections/terraform/providers).

Finally, the [HashiCorp Provider Design
Principles](https://www.terraform.io/docs/extend/hashicorp-provider-design-principles.html)
explore the underlying principles for the design choices of this provider.

## Prerequisites

- [Go 1.13+](https://go.dev/)
- [Docker](https://docs.docker.com/get-docker/) - used for running acceptance tests.
- [Docker-Compose](https://docs.docker.com/compose/install/) - used for running acceptance tests.

## Getting started

To work on the provider, you'll need [Go](http://www.golang.org) installed on
your machine (version 1.13+ is *required*). You'll also need to correctly set up
a [GOPATH](http://golang.org/doc/code.html#GOPATH), as well as adding
`$GOPATH/bin` to your `$PATH`.

To compile the provider, run `make build`. This will build the provider and
install the provider binary in the `$GOPATH/bin` directory.

```sh
make build
...
$GOPATH/bin/terraform-provider-auth0
...
```

In order to test the provider, you can simply run `make test`.

```sh
make test
```

In order to run the full suite of Acceptance tests, the following environment
variables must be set:

```sh
AUTH0_DOMAIN=<your-auth0-tenant-domain>
AUTH0_CLIENT_ID=<your-auth0-client-id>
AUTH0_CLIENT_SECRET=<your-auth0-client-secret>
```

Then, run `make testacc`. 

**Note:** The acceptance tests make calls to a real Auth0 tenant, and create
real resources. Certain tests also require a paid Auth0 subscription to be able to
run successfully, e.g. `TestAccCustomDomain` and the ones starting with `TestAccLogStream*`.

**Note:** At the time of writing, the following configuration steps are also
required for the test tenant:

* The `Username-Password-Authentication` connection must have _Requires
  Username_ option enabled for the user tests to successfully run.

## Documentation

To make it easier to document new resources a handy script is available. The
script can output documentation of a resource in Markdown format, using the 
schema of the resource itself.

```sh
go run scripts/gendocs.go -resource auth0_action
```

## Releasing

The Auth0 provider follows the [Versioning and
Changelog](https://www.terraform.io/docs/extend/best-practices/versioning.html)
guidelines from HashiCorp.

Prepare for the release by updating the [CHANGELOG](CHANGELOG.md). 

To publish a new version, create a new git tag and push it.

```bash
git tag -a vX.Y.Z
git push origin vX.Y.Z
```

This will trigger the
[Release](https://github.com/alexkappa/terraform-provider-auth0/actions/workflows/release.yml)
GitHub Action which creates a new release.
