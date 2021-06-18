---
layout: "auth0"
page_title: "Auth0: auth0_custom_domain_verification"
description: |-
  With this resource, you can verify a custom domain created with the `auth0_custom_domain` resource.
---

# auth0_custom_domain_verification

With Auth0, you can use a custom domain to maintain a consistent user experience. This is a three-step process; you must configure the custom domain in Auth0, then create a DNS record for the domain, then verify the DNS record in Auth0. This resources allows for automating the verification part of the process.

## Example Usage

```hcl
resource "auth0_custom_domain" "my_custom_domain" {
	domain = "login.example.com"
	type = "auth0_managed_certs"
}

resource "auth0_custom_domain_verification" "my_custom_domain_verification" {
	custom_domain_id = auth0_custom_domain.my_custom_domain.id
	timeouts { create = "15m" }
	depends_on = [ digitalocean_record.my_domain_name_record ]
}

resource "digitalocean_record" "my_domain_name_record" {
	domain = "example.com"
	type = upper(auth0_custom_domain.my_custom_domain.verification[0].methods[0].name)
	name = "${auth0_custom_domain.my_custom_domain.domain}."
	value = "${auth0_custom_domain.my_custom_domain.verification[0].methods[0].record}."
}
```

## Argument Reference

Arguments accepted by this resource include:

* `custom_domain_id` - (Required) String. ID of the custom domain resource.

## Meta-Arguments

`auth0_custom_domain_verification` can be used with the `depends_on` [meta-argument](https://www.terraform.io/docs/language/resources/syntax.html#meta-arguments) to explicitly wait for the domain name record (DNS) to be created before attempting to verify the custom domain. 

## Operation Timeouts

`auth0_custom_domain_verification` provides the following [`timeouts`](https://www.terraform.io/docs/configuration/blocks/resources/syntax.html#operation-timeouts):

`create` - (Default `5m`) How long to wait for a certificate to be issued.