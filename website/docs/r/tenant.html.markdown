---
layout: "auth0"
page_title: "Auth0: auth0_tenant"
sidebar_current: "docs-auth0-resource-tenant"
description: |-
  Manages the Auth0 Tenant.
---

# auth0_tenant

The Auth0 Tenant as defined by the domain set in the provider settings.

~> **Note:** Auth0's Management API doesn't support creating multiple Tenants
   and Terraform doesn't support restricting the number of resources per type.
   Therefore, you should only ever have a single Tenant in your Terraform code
   otherwise unexpected behavior will result due to resources overwriting each other's
   changes.

## Example Usage

```hcl
resource "auth0_tenant" "tenant" {
  change_password {
    enabled = true
    html = "${file("./password_reset.html")}"
  }

  guardian_mfa_page {
    enabled = true
    html = "${file("./guardian_multifactor.html")}"
  }

  default_audience  = "<client_id>"
  default_directory = "Connection-Name"

  error_page {
      html          = "${file("./error.html")}"
      show_log_link = true
      url           = "http://mysite/errors"
  }

  friendly_name       = "Tenant Name"
  picture_url         = "http://mysite/logo.png"
  support_email       = "support@mysite"  
  support_url         = "http://mysite/support"
  allowed_logout_urls = [
      "http://mysite/logout"
  ]
  session_lifetime    = 46000
  sandbox_version     = "8"  
}
```

## Argument Reference

The following arguments are supported:

* `change_password` - (Optional) A `change_password` block as defined below.
* `guardian_mfa_page` - (Optional) A `guardian_mfa_page` block as defined below.
  Defined below.
* `default_audience` - (Optional) Default audience (client ID) for API Authorization.
* `default_directory` - (Optional) Name of the connection that will be used for password grants at the token endpoint. Only the following connection types are supported: LDAP, AD, Database Connections, Passwordless, Windows Azure Active Directory, ADFS.
* `error_page` - (Optional) A `error_page` block as defined below.
* `friendly_name` - (Optional) The friendly name of the tenant.
* `picture_url` - (Optional) The URL of the tenant logo (recommended size: 150x150).
* `support_email` - (Optional) User support email.
* `support_url` - (Optional) User support url.
* `allowed_logout_urls` - (Optional) A set of URLs that are valid to redirect to after logout from Auth0.
* `session_lifetime` - (Optional) Force a user to login after they have been inactive for the specified number (unit: hours).
* `sandbox_version` - (Optional) The selected sandbox version to be used for the extensibility environment.

### `change_password` block

* `enabled` - (Required) Whether to enable the custom page.page.
* `html` - (Required) HTML content of the custom page.

### `guardian_mfa_page` block

* `enabled` - (Required) Whether to enable the custom page.
* `html` - (Required) HTML content of the custom page.

### `error_page` block

* `html` - (Required) HTML content of the custom page.
* `show_log_link` - (Required) Set to `true` to show link to log as part of the default error page,
* `url` - (Required) Redirect to specified url instead of show the default error page.
