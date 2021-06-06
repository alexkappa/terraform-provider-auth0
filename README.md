Auth0 Terraform Provider
========================

[![Build](https://github.com/alexkappa/terraform-provider-auth0/workflows/Build/badge.svg)](https://github.com/alexkappa/terraform-provider-auth0/actions)
[![Maintainability](https://api.codeclimate.com/v1/badges/9c49c10286123b716c79/maintainability)](https://codeclimate.com/github/alexkappa/terraform-provider-auth0/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/9c49c10286123b716c79/test_coverage)](https://codeclimate.com/github/alexkappa/terraform-provider-auth0/test_coverage)
[![Gitter](https://badges.gitter.im/terraform-provider-auth0/community.svg)](https://gitter.im/terraform-provider-auth0/community)

Sponsors
--------

| <img width="50" src="https://cdn.auth0.com/blog/open-source-sponsorship.png"> | <div style="text-align: left;">If you want to quickly implement a secure authentication flow with Terraform, create a free plan at [auth0.com/developers](https://auth0.com/developers?utm_source=GHsponsor&utm_medium=GHsponsor&utm_campaign=terraform_auth0_provider&utm_content=auth).</div> |
| :-: | :- |
| <img width="50" src="https://placehold.co/50x50?text=?"> | If you or your company relies on this provider and would like to ensure its continuing support please consider [sponsoring](https://github.com/sponsors/alexkappa). |

Usage
-----

**Terraform 0.13+**

Terraform 0.13 and higher uses the [Terraform Registry](https://registry.terraform.io/) to download and install providers. To install this provider, copy and paste this code into your Terraform configuration. Then, run `terraform init`.

```tf
terraform {
  required_providers {
    auth0 = {
      source  = "alexkappa/auth0"
      version = "0.17.1"
    }
  }
}

provider "auth0" {}
```

```sh
$ terraform init
```

**Terraform 0.12.x**

For older versions of Terraform, binaries are available at the [releases](https://github.com/alexkappa/terraform-provider-auth0/releases) page. Download one that corresponds to your operating system / architecture, and move to the `~/.terraform.d/plugins/` directory. Finally, run terraform init.

```
provider "auth0" {}
```


```sh
$ terraform init
```

See the [Auth0 Provider documentation](https://registry.terraform.io/providers/alexkappa/auth0/latest/docs) for all the available resources.

Contributing
------------

See [CONTRIBUTING.md](CONTRIBUTING.md).