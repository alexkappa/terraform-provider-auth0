---
layout: "auth0"
page_title: "Auth0: auth0_branding"
description: |-
  With this resource, you can configure the branding for the Universal Login Experience.
---

# auth0_branding

With this resource, you can configure the branding for the Universal Login Experience.

## Example Usage

```hcl
resource "auth0_branding" "example" {
  favicon_url = "https://mysite/favicon.png"
  logo_url    = "https://mysite/logo.png"
  
  font {
    url = "https://mysite/font.ttf"
  }

  colors {
    primary         = "#ffffff"
    page_background = "#000000"
  }
}
```

### With a gradient page background

```hcl
resource "auth0_branding" "example" {
  colors {
    primary = "#ffffff"
    
    page_background_gradient {
      type      = "linear-gradient"
      start     = "#333333"
      end       = "#aaaaaa"
      angle_deg = 35
    }
  }
}
```

## Argument Reference

The following arguments are supported:


* `favicon_url` - (Optional) String. URL for the favicon. Must use HTTPS.
* `logo_url` - (Optional) String. URL for the logo. Must use HTTPS.
* `colors` - (Optional) List(Resource). Configuration settings to customized the branding colors. For details see [Colors](#colors). 
* `font` - (Optional) List(Resource). Configuration settings to customize the font. For details see [Font](#font).

### `colors`

Specify the colors that the universal login experience should use.

For the background, you can either use a static color with `page_background`, or color gradient with `page_background_color`. It is not possible to define both.

#### Arguments

* `primary` - (Optional) String. Accent color.
* `page_background` - (Optional) String. Page background color. Conflicts with `page_background_gradient`.
* `page_background_gradient` - (Optional) List(Resource). Configuration settings to define a gradient page background. For details see [Page Background Gradient](#page_background_gradient).  Conflicts with `page_background`.

#### Attributes

### `font`

#### Arguments

* `url` - (Required) String. URL for the custom font. Must use HTTPS.

#### Attributes

### `page_background_gradient`

#### Arguments

* `type` - (Optional) String. The type of gradient. 
* `start` - (Optional) String. The start color of the gradient.
* `end` - (Optional) String. The end colors of the gradient.
* `angle_deg` - (Optional) Number. Degrees that the background gradient is rotated.

#### Attributes

## Attributes Reference

In addition to the arguments listed above, no other attributes are exported.

## Import

auth0_branding can be imported using an arbitrary id, e.g:

```
$ terraform import auth0_branding.example 123
```
