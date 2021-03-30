---
layout: "auth0"
page_title: "Auth0: auth0_branding"
description: |-
  With this resource, you can manage branding, including logo, color.
---

# auth0_branding

With Auth0, you can setting logo, color to maintain a consistent service brand. This resource allows you to manage a branding within your Auth0 tenant.

## Example Usage

```
resource "auth0_branding" "my_brand" {
	logo_url = "https://mycompany.org/logo.png"
	colors {
		primary = "#0059d6"
		page_background = "#000000"
	}
}
```

## Argument Reference

The following arguments are supported:

* `colors` - (Optional) List(Resource). Configuration settings for Universal Login colors. See [Colors](#colors).
* `logo_url` - (Optional) String. Configuration settings for Universal Login logo URL.

### `colors`

`colors` supports the following arguments:

* `page_background` - (Optional) String, Hexadecimal. Background color of login pages.
* `primary` - (Optional) String, Hexadecimal. Primary button background color.
