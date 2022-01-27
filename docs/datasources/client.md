---
layout: "auth0"
page_title: "Data Source: auth0_client"
description: |-
Use this data source to get information about a specific Auth0 Application client by its 'client_id' or 'name'
---

# Data Source: auth0_client

Use this data source to get information about a specific Auth0 Application client by its 'client_id' or 'name'

## Example Usage

```hcl
data "auth0_client" "some-client-by-name" {
  name = "Name of my Application"
}
data "auth0_client" "some-client-by-id" {
  client_id = "abcdefghkijklmnopqrstuvwxyz0123456789"
}
```

## Argument Reference

Arguments accepted by this data source include (exactly one is required):

- `client_id` - (Optional) String. client_id of the application to retrieve
- `name` - (Optional) String. name of the application to retrieve. Ignored if `client_id` is also specified.

## Attribute Reference

* `client_id` - String. ID of the client.
* `client_secret`<sup>[1](#client-keys)</sup> - String. Secret for the client; keep this private.
* `name` - String. Name of the client.
* `description` - String, Description of the purpose of the client.
* `is_first_party` - Boolean. Indicates whether or not this client is a first-party client.
* `is_token_endpoint_ip_header_trusted` - Boolean
* `oidc_conformant` - Boolean. Indicates whether or not this client will conform to strict OIDC specifications.
* `token_endpoint_auth_method` - String. Defines the requested authentication method for the token endpoint. Options include `none` (public client without a client secret), `client_secret_post` (client uses HTTP POST parameters), `client_secret_basic` (client uses HTTP Basic).
* `app_type` - String. Type of application the client represents. Options include `native`, `spa`, `regular_web`, `non_interactive`, `rms`, `box`, `cloudbees`, `concur`, `dropbox`, `mscrm`, `echosign`, `egnyte`, `newrelic`, `office365`, `salesforce`, `sentry`, `sharepoint`, `slack`, `springcm`, `zendesk`, `zoom`.
* `logo_uri` - String. URL of the logo for the client. Recommended size is 150px x 150px. If none is set, the default badge for the application type will be shown.
* `is_first_party` - Boolean. Indicates whether or not this client is a first-party client.
* `is_token_endpoint_ip_header_trusted` - Boolean. Indicates whether or not the token endpoint IP header is trusted.
* `oidc_conformant` - Boolean. Indicates whether or not this client will conform to strict OIDC specifications.
* `callbacks` - List(String). URLs that Auth0 may call back to after a user authenticates for the client. Make sure to specify the protocol (https://) otherwise the callback may fail in some cases. With the exception of custom URI schemes for native clients, all callbacks should use protocol https://.
* `allowed_logout_urls` - List(String). URLs that Auth0 may redirect to after logout.
* `grant_types` - List(String). Types of grants that this client is authorized to use.
* `allowed_origins` - List(String). URLs that represent valid origins for cross-origin resource sharing. By default, all your callback URLs will be allowed.
* `web_origins` - List(String). URLs that represent valid web origins for use with web message response mode.
* `jwt_configuration` - List(Resource). Configuration settings for the JWTs issued for this client. For details, see [JWT Configuration](#jwt-attribute).
* `refresh_token` - List(Resource). Configuration settings for the refresh tokens issued for this client.  For details, see [Refresh Token Configuration](#refresh-token-attribute).
* `encryption_key` - Map(String).
* `sso` - Boolean. Indicates whether or not the client should use Auth0 rather than the IdP to perform Single Sign-On (SSO). True = Use Auth0.
* `sso_disabled` - Boolean. Indicates whether or not SSO is disabled.
* `cross_origin_auth` - Boolean. Indicates whether or not the client can be used to make cross-origin authentication requests.
* `cross_origin_loc` - String. URL for the location on your site where the cross-origin verification takes place for the cross-origin auth flow. Used when performing auth in your own domain instead of through the Auth0-hosted login page.
* `custom_login_page_on` - Boolean. Indicates whether or not a custom login page is to be used.
* `custom_login_page` - String. Content of the custom login page.

### JWT Attribute

`jwt_configuration` outputs the following attributes:

* `lifetime_in_seconds` - Integer. Number of seconds during which the JWT will be valid.
* `secret_encoded` - Boolean. Indicates whether or not the client secret is base64 encoded.
* `scopes` - Map(String). Permissions (scopes) included in JWTs.
* `alg` - String. Algorithm used to sign JWTs.

### Refresh Token Attribute

`refresh_token` outputs the following attributes:

* `rotation_type` - String. Options include `rotating`, `non-rotating`. When `rotating`, exchanging a refresh token will cause a new refresh token to be issued and the existing token will be invalidated. This allows for automatic detection of token reuse if the token is leaked.
* `leeway` - Integer. The amount of time in seconds in which a refresh token may be reused without trigging reuse detection.
* `expiration_type` - (Optional unless `rotation_type` is `rotating`) String. Options include `expiring`, `non-expiring`. Whether a refresh token will expire based on an absolute lifetime, after which the token can no longer be used. If rotation is `rotating`, this must be set to `expiring`.
* `token_lifetime` - Integer. The absolute lifetime of a refresh token in seconds.
* `infinite_idle_token_lifetime` - Boolean, (Default=false) Whether or not inactive refresh tokens should be remain valid indefinitely.
* `infinite_token_lifetime` - Boolean, (Default=false) Whether or not refresh tokens should remain valid indefinitely. If false, `token_lifetime` should also be set
* `idle_token_lifetime` - Integer. The time in seconds after which inactive refresh tokens will expire.

### Client keys

To access the `client_secret` attribute you need to add the `read:client_keys` scope to the Terraform client. Otherwise, the attribute will contain an empty string.