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
  domain = "auth.example.com"
  type = "auth0_managed_certs"
  verification_method = "txt"
}
resource "digitalocean_record" "auth0_domain" {
  domain = "example.com"
  type   = upper(auth0_custom_domain.my_custom_domain.verification[0].methods[0].name)
  name   = "auth"
  value  = "${auth0_custom_domain.my_custom_domain.verification[0].methods[0].record}."
}
# wait for DNS record to propagate
resource "null_resource" "wait_for_auth0_dns" {
  provisioner "local-exec" {
    command = "while ! nslookup ${digitalocean_record.auth0_domain.fqdn}; do sleep 1; done"
  }
  triggers = {
    dns = digitalocean_record.auth0_domain.id
  }
}
resource "auth0_custom_domain_verification" "my_custom_domain" {
  custom_domain_id = auth0_custom_domain.my_custom_domain.id
  depends_on = [null_resource.wait_for_auth0_dns]
}
```

## Argument Reference

Arguments accepted by this resource include:

* `custom_domain_id` - (Required) String. ID of the custom domain resource.
