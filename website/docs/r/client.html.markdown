---
layout: "auth0"
page_title: "auth0_client"
sidebar_current: "docs-auth0-resource-client"
description: |-
  Auth0 Client (Application)
---

# Auth0 Client (Application)

The `auth0_client` resource provides an [Auth0 Client](https://auth0.com/docs/api/management/v2#!/Clients/get_clients).

## Example Usage

Client's are an integral part of an Auth0 setup. It can

```hcl
resource "auth0_client" "my_app_client" {
  name            = "Example Application"
  description     = "Example Application Description"
  app_type        = "regular_web"
  is_first_party  = true
  oidc_conformant = false
  callbacks       = ["https://example.com/callback"]
  allowed_origins = ["https://example.com"]
  web_origins     = ["https://example.com"]
  grant_types     = ["authorization_code", "http://auth0.com/oauth/grant-type/password-realm", "implicit", "password", "refresh_token"]

  jwt_configuration = {
    lifetime_in_seconds = 120
    secret_encoded      = true
    alg                 = "RS256"
  }

  custom_login_page_on = "true"
}
```

## Argument Reference

The following arguments are supported:

* `name` - The name of the client.
* `description` - Free text description of the purpose of the Client. Maximum character length is `140` characters.
* `app_type` - (Optional) The type of application this client represents
* `logo_uri` - (Optional) The URL of the client logo (recommended size: 150x150)
* `is_first_party` - (Optional) Whether this client a first party client or not
* `is_token_endpoint_ip_header_trusted` - (Optional)
* `oidc_conformant` - (Optional) Whether this client will conform to strict OIDC specifications
* `callbacks` - (Optional) The URLs that Auth0 can use to as a callback for the client
* `allowed_logout_urls` - (Optional)
* `allowed_origins` - (Optional)
* `web_origins` - (Optional) A set of URLs that represents valid web origins for use with web message response mode
* `sso` - (Optional) Bool
* `sso_disabled` - (Optional) Bool
* `cross_origin_auth` - (Optional) Bool
* `cross_origin_loc` - (Optional) String
* `custom_login_page_on` - (Optional) Bool
* `custom_login_page` - (Optional) String
* `custom_login_page_preview` - (Optional) String
* `form_template` - (Optional) String
* `token_endpoint_auth_method` - (Optional) Defines the requested authentication method for the token endpoint. Possible values are `none` (public client without a client secret), `client_secret_post` (client uses HTTP POST parameters) or `client_secret_basic` (client uses HTTP Basic).
* `client_metadata` - (Optional) Key value pairs of metadata to associate to this client.

The nested `jwt_configuration` block has the following structure:

* `lifetime_in_seconds` - (Optional) The amount of seconds the JWT will be valid (affects `exp` claim)
* `secret_encoded` - (Optional) Set to `true` if the client secret is base64 encoded, `false` otherwise. Defaults to `true`
* `scopes` - (Optional) Algorithm uses to sign JWTs
* `alg` - (Optional) One of `HS256` or `RS256`

The nested `addons` block has the following structure:

* `aws` - (Optional) Map
* `azure_blob` - (Optional) Map
* `azure_sb` - (Optional) Map
* `rms` - (Optional) Map
* `mscrm` - (Optional) Map
* `slack` - (Optional) Map
* `sentry` - (Optional) Map
* `box` - (Optional) Map
* `cloudbees` - (Optional) Map
* `concur` - (Optional) Map
* `dropbox` - (Optional) Map
* `echosign` - (Optional) Map
* `egnyte` - (Optional) Map
* `firebase` - (Optional) Map
* `newrelic` - (Optional) Map
* `office365` - (Optional) Map
* `salesforce` - (Optional) Map
* `salesforce_api` - (Optional) Map
* `salesforce_sandbox_api` - (Optional) Map
* `layer` - (Optional) Map
* `sap_api` - (Optional) Map
* `sharepoint` - (Optional) Map
* `springcm` - (Optional) Map
* `wams` - (Optional) Map
* `wsfed` - (Optional) Map
* `zendesk` - (Optional) Map
* `zoom` - (Optional) Map
* `addons.samlp`
  - `audience` - (Optional) String
  - `recipient` - (Optional) String
  - `create_upn_claim` - (Optional) Bool
  - `passthrough_claims_with_no_mapping` - (Optional) Bool
  - `map_unknown_claims_as_is` - (Optional) Bool
  - `map_identities` - (Optional) Bool
  - `signature_algorithm` - (Optional) String
  - `digest_algorithm` - (Optional) String
  - `destination` - (Optional) String
  - `lifetime_in_seconds` - (Optional) Int
  - `sign_response` - (Optional) Bool
  - `typed_attributes` - (Optional) Bool
  - `include_attribute_name_format` - (Optional) Bool
  - `name_identifier_format` - (Optional) String
  - `authn_context_class_ref` - (Optional) String
  - `binding` - (Optional) String
  - `mappings` - (Optional) Map
  - `logout` - (Optional) Map
  - `callback` - (Optional) String
  - `slo_enabled` - (Optional) Bool
  - `name_identifier_probes` - (Optional) List

The nested `mobile` block has the following structure:

* `android`
* `app_package_name` - (Optional) String
* `sha256_cert_fingerprints` - (Optional) List
* `ios` - (Optional) List
* `team_id` - (Optional) String
* `app_bundle_identifier` - (Optional) String

## Attributes Reference

The following attributes are exported:

* `client_id` - The id of the client
* `client_secret` - The client secret
* `grant_types` -

The nested `encryption_key` block has the following structure:

* `pub`
* `cert`
* `subject`
