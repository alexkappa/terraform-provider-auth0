---
layout: "auth0"
page_title: "Auth0: auth0_connection"
description: |-
  With this resource, you can configure and manage sources of users, which may include identity providers, databases, or passwordless authentication methods.
---

# auth0_connection

With Auth0, you can define sources of users, otherwise known as connections, which may include identity providers (such as Google or LinkedIn), databases, or passwordless authentication methods. This resource allows you to configure and manage connections to be used with your clients and users.

## Example Usage

```hcl
resource "auth0_connection" "my_connection" {
  name = "Example-Connection"
  strategy = "auth0"
  options {
    password_policy = "excellent"
    password_history {
      enable = true
      size = 3
    }
    brute_force_protection = "true"
    enabled_database_customization = "true"
    custom_scripts = {
      get_user = <<EOF
function getByEmail (email, callback) {
  return callback(new Error("Whoops!"))
}
EOF
    }

    configuration = {
      foo = "bar"
      bar = "baz"
    }
  }
}
```

~> **Note**: The Auth0 dashboard displays only one connection per social provider. Although the Auth0 Management API allowes the creation of multiple connections per strategy, the additional connections may not be visible in the Auth0 dashboard.

## Argument Reference

Arguments accepted by this resource include:

* `name` - (Required) Name of the connection.
* `is_domain_connection` - (Optional) Indicates whether or not the connection is domain level.
* `strategy` - (Required) Type of the connection, which indicates the identity provider. Options include `ad`, `adfs`, `amazon`, `aol`, `apple`, `auth0`, `auth0-adldap`, `auth0-oidc`, `baidu`, `bitbucket`, `bitly`, `box`, `custom`, `daccount`, `dropbox`, `dwolla`, `email`, `evernote`, `evernote-sandbox`, `exact`, `facebook`, `fitbit`, `flickr`, `github`, `google-apps`, `google-oauth2`, `guardian`, `instagram`, `ip`, `line`, `linkedin`, `miicard`, `oauth1`, `oauth2`, `office365`, `oidc`, `paypal`, `paypal-sandbox`, `pingfederate`, `planningcenter`, `renren`, `salesforce`, `salesforce-community`, `salesforce-sandbox` `samlp`, `sharepoint`, `shopify`, `sms`, `soundcloud`, `thecity`, `thecity-sandbox`, `thirtysevensignals`, `twitter`, `untappd`, `vkontakte`, `waad`, `weibo`, `windowslive`, `wordpress`, `yahoo`, `yammer`, `yandex`.
* `options` - (Optional) Configuration settings for connection options. For details, see [Options](#options).
* `enabled_clients` - (Optional) IDs of the clients for which the connection is enabled. If not specified, no clients are enabled.
* `realms` - (Optional) Defines the realms for which the connection will be used (i.e., email domains). If not specified, the connection name is added as the realm.

### Options

`options` supports different arguments depending on the connection `strategy` defined in [Argument Reference](#argument-reference).

### Auth0

With the `auth0` connection strategy, `options` supports the following arguments:

* `validation` - (Optional) A map defining the validation options.
* `password_policy` - (Optional) Indicates level of password strength to enforce during authentication. A strong password policy will make it difficult, if not improbable, for someone to guess a password through either manual or automated means. Options include `none`, `low`, `fair`, `good`, `excellent`.
* `password_history` - (Optional) Configuration settings for the password history that is maintained for each user to prevent the reuse of passwords. For details, see [Password History](#password-history).
* `password_no_personal_info` - (Optional) Configuration settings for the password personal info check, which does not allow passwords that contain any part of the user's personal data, including user's name, username, nickname, user_metadata.name, user_metadata.first, user_metadata.last, user's email, or first part of the user's email. For details, see [Password No Personal Info](#password-no-personal-info).
* `password_dictionary` - (Optional) Configuration settings for the password dictionary check, which does not allow passwords that are part of the password dictionary. For details, see [Password Dictionary](#password-dictionary).
* `password_complexity_options` - (Optional) Configuration settings for password complexity. For details, see [Password Complexity Options](#password-complexity-options).
* `api_enable_users` - (Optional)
* `enabled_database_customization` - (Optional)
* `brute_force_protection` - (Optional) Indicates whether or not to enable brute force protection, which will limit the number of signups and failed logins from a suspicious IP address.
* `import_mode` - (Optional) Indicates whether or not you have a legacy user store and want to gradually migrate those users to the Auth0 user store. [Learn more](https://auth0.com/docs/users/guides/configure-automatic-migration).
* `disable_signup` - (Optional) Boolean. Indicates whether or not to allow user sign-ups to your application.
* `requires_username` - (Optional) Indicates whether or not the user is required to provide a username in addition to an email address.
* `custom_scripts` - (Optional) Custom database action scripts. For more information, read [Custom Database Action Script Templates](https://auth0.com/docs/connections/database/custom-db/templates).
* `configuration` - (Optional) A case-sensitive map of key value pairs used as configuration variables for the `custom_script`.

#### Password History

`password_history` supports the following arguments:

* `enable` - (Optional) Indicates whether password history is enabled for the connection. When enabled, any existing users in this connection will be unaffected; the system will maintain their password history going forward.
* `size` - (Optional) Indicates the number of passwords to keep in history with a maximum of 24.

#### Password No Personal Info

`password_no_personal_info` supports the following arguments:

* `enable` - (Optional) Indicates whether the password personal info check is enabled for this connection.

#### Password Dictionary

`passsword_dictionary` supports the following arguments:

* `enable` - (Optional) Indicates whether the password dictionary check is enabled for this connection.
* `dictionary` - (Optional) Customized contents of the password dictionary. By default, the password dictionary contains a list of the [10,000 most common passwords](https://github.com/danielmiessler/SecLists/blob/master/Passwords/Common-Credentials/10k-most-common.txt); your customized content is used in addition to the default password dictionary. Matching is not case-sensitive.

#### Password Complexity Options

`password_complexity_options` supports the following arguments:

* `min_length`- (Optional) Minimum number of characters allowed in passwords.

### Google OAuth2

~> **Note**: Your Auth0 account may be pre-configured with a `google-oauth2` connection. To manage that connection with terraform see the [import example](#import).

With the `google-oauth2` connection strategy, `options` supports the following arguments:

* `client_id` - (Optional) Facebook client ID.
* `client_secret` - (Optional) Facebook client secret.
* `allowed_audiences` - (Optional) List of allowed audiences.
* `scopes` - (Optional) Scopes.



**Example**:

```hcl
resource "auth0_connection" "google_oauth2" {
  name = "Google-OAuth2-Connection"
  strategy = "google-oauth2"
  options {
    client_id = "<client-id>"
    client_secret = "<client-secret>"
    allowed_audiences = [ "example.com", "api.example.com" ]
    scopes = [ "email", "profile", "gmail", "youtube" ]
  }
}
```

### Facebook

With the `facebook` connection strategy, `options` supports the following arguments:

* `client_id` - (Optional) Facebook client ID.
* `client_secret` - (Optional) Facebook client secret.
* `scopes` - (Optional) Scopes.

**Example**:

```hcl
resource "auth0_connection" "facebook" {
  name = "Facebook-Connection"
  strategy = "facebook"
  options {
    client_id = "<client-id>"
    client_secret = "<client-secret>"
    scopes = [ "public_profile",  "email",  "groups_access_member_info",  "user_birthday" ]
  }
}
```

### Apple

With the `apple` connection strategy, `options` supports the following arguments:

* `client_id` - (Optional) Apple client ID.
* `client_secret` - (Optional) App secret.
* `team_id` - (Optional) Team ID.
* `key_id` - (Optional) Key ID.
* `scopes` - (Optional) Scopes.

**Example**:

```hcl
resource "auth0_connection" "apple" {
  name = "Apple-Connection"
  strategy = "apple"
  options {
    client_id = "<client-id>"
    client_secret = "<private-key>"
    team_id = "<team-id>"
    key_id = "<key-id>"
    scopes = ["email", "name"]
  }
}
```

### Linkedin

With the `linkedin` connection strategy, `options` supports the following arguments:

* `client_id` - (Optional) Linkedin API key.
* `client_secret` - (Optional) Linkedin secret key.
* `strategy_version` - (Optional) Version 1 is deprecated, use version 2.
* `scopes` - (Optional) Scopes.

**Example**:

```hcl
resource "auth0_connection" "linkedin" {
  name = "Linkedin-Connection"
  strategy = "linkedin"
  options {
    client_id = "<client-id>"
    client_secret = "<client-secret>"
    strategy_version = 2
    scopes = [ "basic_profile", "profile", "email" ]
  }
}
```

### GitHub

With the `github` connection strategy, `options` supports the following arguments:

* `client_id` - (Optional) GitHub client ID.
* `client_secret` - (Optional) GitHub client secret.
* `set_user_root_attributes` - (Optional)

**Example**:

```hcl
resource "auth0_connection" "github" {
  name = "GitHub-Connection"
  strategy = "github"
  options {
    client_id = "<client-id>"
    client_secret = "<client-secret>"
    scopes = [ "email", "profile", "public_repo", "repo" ]
  }
}
```

### Salesforce

With the `salesforce`, `salesforce-community` and `salesforce-sandbox` connection strategies, `options` supports the following arguments:

* `client_id` - (Optional) The Salesforce client ID.
* `client_secret` - (Optional) The Salesforce client secret.
* `community_base_url` - (Optional) String.
* `scopes` - (Optional) Scopes.

**Example**:

```hcl
resource "auth0_connection" "salesforce" {
	name = "Salesforce-Connection"
	strategy = "salesforce"
	options {
		client_id = "<client-id>"
		client_secret = "<client-secret>"
		community_base_url = "https://salesforce.example.com"
	}
}
```

### OIDC

With the `oidc` connection strategy, `options` supports the following arguments:

* `client_id` - (Optional) OIDC provider client ID.
* `client_secret` - (Optional) OIDC provider client secret.
* `type` - (Optional) Value can be `back_channel` or `front_channel`.
* `scopes` - (Optional) Scopes required by the connection. The value must be a list, for example `["openid", "profile", "email"]`.
* `issuer` - (Optional) Issuer URL. E.g. `https://auth.example.com`
* `discovery_url` - (Optional) OpenID discovery URL. E.g. `https://auth.example.com/.well-known/openid-configuration`.
* `jwks_uri` - (Optional)
* `token_endpoint` - (Optional)
* `userinfo_endpoint` - (Optional)
* `authorization_endpoint` - (Optional)

### Azure AD

With the `waad` connection strategy, `options` supports the following arguments:

* `app_id` - (Optional) Azure AD app ID.
* `app_domain` - (Optional) Azure AD domain name.
* `client_id` - (Optional) Client ID for your Azure AD application.
* `client_secret` - (Optional) Client secret for your Azure AD application.
* `domain_aliases` - (Optional) List of the domains that can be authenticated using the Identity Provider. Only needed for Identifier First authentication flows.
* `max_groups_to_retrieve` - (Optional) Maximum number of groups to retrieve.
* `tenant_domain` - (Optional)
* `use_wsfed` - (Optional)
* `waad_protocol` - (Optional)
* `waad_common_endpoint` - (Optional) Indicates whether or not to use the common endpoint rather than the default endpoint. Typically enabled if you're using this for a multi-tenant application in Azure AD.

### Twilio / SMS

With the `sms` connection strategy, `options` supports the following arguments:

* `name` - (Optional)
* `twilio_sid` - (Optional) SID for your Twilio account.
* `twilio_token` - (Optional) AuthToken for your Twilio account.
* `from` - (Optional) SMS number for the sender. Used when SMS Source is From.
* `syntax` - (Optional) Syntax of the SMS. Options include `markdown` and `liquid`.
* `template` - (Optional) Template for the SMS. You can use `@@password@@` as a placeholder for the password value.
* `totp` - (Optional) Configuration options for one-time passwords. For details, see [TOTP](#totp).
* `messaging_service_sid` - (Optional) SID for Copilot. Used when SMS Source is Copilot.

#### TOTP

`totp` supports the following arguments:

* `time_step` - (Optional) Integer. Seconds between allowed generation of new passwords.
* `length` - (Optional) Integer. Length of the one-time password.

**Example**:

```hcl
resource "auth0_connection" "sms" {
  name = "SMS-Connection"
  strategy = "sms"
  options {
    name = "SMS OTP"
    twilio_sid = "<twilio-sid>"
    twilio_token = "<twilio-token>"
    from = "<phone-number>"
    syntax = "md_with_macros"
    template = "Your one-time password is @@password@@"
    messaging_service_sid = "<messaging-service-sid>"
    disable_signup = false
    brute_force_protection = true
    totp {
      time_step = 300
      length = 6
    }
  }
}
```

### ADFS

With the `adfs` connection strategy, `options` supports the following arguments:

* `adfs_server` - (Optional) ADFS Metadata source.

### SAML

With the `samlp` connection strategy, `options` supports the following arguments:

* `signing_cert` - The X.509 signing certificate (encoded in PEM or CER) you retrieved from the IdP, Base64-encoded
* `binding_method` - (Optional) The SAML Response Binding - how the SAML token is received by Auth0 from IdP. Two possible values are `urn:oasis:names:tc:SAML:2.0:bindings:HTTP-Redirect` (default) and `urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST`
* `tenant_domain` - (Optional)
* `domain_aliases` - (Optional) List of the domains that can be authenticated using the Identity Provider. Only needed for Identifier First authentication flows.
* `sign_in_endpoint` - SAML single login URL for the connection.
* `sign_out_endpoint` - (Optional) SAML single logout URL for the connection.
* `fields_map` - (Optional) SAML Attributes mapping. If you're configuring a SAML enterprise connection for a non-standard PingFederate Server, you must update the attribute mappings.
* `sign_saml_request` - (Optional) (Boolean) When enabled, the SAML authentication request will be signed.
* `signature_algorithm` - (Optional) Sign Request Algorithm
* `digest_algorithm` - (Optional) Sign Request Algorithm Digest

**Example**:
```hcl
resource "auth0_connection" "samlp" {
	name = "SAML-Connection"
	strategy = "samlp"
	options {
		signing_cert = <<EOF
-----BEGIN CERTIFICATE-----
MIIDujCCAqKgAwIBAgIIE31FZVaPXTUwDQYJKoZIhvcNAQEFBQAwSTELMAkGA1UE
BhMCVVMxEzARBgNVBAoTCkdvb2dsZSBJbmMxJTAjBgNVBAMTHEdvb2dsZSBJbnRl
cm5ldCBBdXRob3JpdHkgRzIwHhcNMTQwMTI5MTMyNzQzWhcNMTQwNTI5MDAwMDAw
WjBpMQswCQYDVQQGEwJVUzETMBEGA1UECAwKQ2FsaWZvcm5pYTEWMBQGA1UEBwwN
TW91bnRhaW4gVmlldzETMBEGA1UECgwKR29vZ2xlIEluYzEYMBYGA1UEAwwPbWFp
bC5nb29nbGUuY29tMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEfRrObuSW5T7q
5CnSEqefEmtH4CCv6+5EckuriNr1CjfVvqzwfAhopXkLrq45EQm8vkmf7W96XJhC
7ZM0dYi1/qOCAU8wggFLMB0GA1UdJQQWMBQGCCsGAQUFBwMBBggrBgEFBQcDAjAa
BgNVHREEEzARgg9tYWlsLmdvb2dsZS5jb20wCwYDVR0PBAQDAgeAMGgGCCsGAQUF
BwEBBFwwWjArBggrBgEFBQcwAoYfaHR0cDovL3BraS5nb29nbGUuY29tL0dJQUcy
LmNydDArBggrBgEFBQcwAYYfaHR0cDovL2NsaWVudHMxLmdvb2dsZS5jb20vb2Nz
cDAdBgNVHQ4EFgQUiJxtimAuTfwb+aUtBn5UYKreKvMwDAYDVR0TAQH/BAIwADAf
BgNVHSMEGDAWgBRK3QYWG7z2aLV29YG2u2IaulqBLzAXBgNVHSAEEDAOMAwGCisG
AQQB1nkCBQEwMAYDVR0fBCkwJzAloCOgIYYfaHR0cDovL3BraS5nb29nbGUuY29t
L0dJQUcyLmNybDANBgkqhkiG9w0BAQUFAAOCAQEAH6RYHxHdcGpMpFE3oxDoFnP+
gtuBCHan2yE2GRbJ2Cw8Lw0MmuKqHlf9RSeYfd3BXeKkj1qO6TVKwCh+0HdZk283
TZZyzmEOyclm3UGFYe82P/iDFt+CeQ3NpmBg+GoaVCuWAARJN/KfglbLyyYygcQq
0SgeDh8dRKUiaW3HQSoYvTvdTuqzwK4CXsr3b5/dAOY8uMuG/IAR3FgwTbZ1dtoW
RvOTa8hYiU6A475WuZKyEHcwnGYe57u2I2KbMgcKjPniocj4QzgYsVAVKW3IwaOh
yE+vPxsiUkvQHdO2fojCkY8jg70jxM+gu59tPDNbw3Uh/2Ij310FgTHsnGQMyA==
-----END CERTIFICATE-----
EOF
		sign_in_endpoint = "https://saml.provider/sign_in"
		sign_out_endpoint = "https://saml.provider/sign_out"
		tenant_domain = "example.con"
		domain_aliases = ["example.con", "example.coz"]
		binding_method = "urn:oasis:names:tc:SAML:2.0:bindings:HTTP-Post"
		signature_algorithm = "rsa-sha256"
		digest_algorithm = "sha256"
		fields_map = {
			foo = "bar"
			baz = "baa"
		}
	}
}
```

## Attribute Reference

Attributes exported by this resource include:

* `is_domain_connection` - Boolean. Indicates whether or not the connection is domain level.
* `options` - List(Resource). Configuration settings for connection options. For details, see [Options Attributes](#options-attributes).
* `realms` - List(String). Defines the realms for which the connection will be used (i.e., email domains). If the array is empty or the property is not specified, the connection name is added as the realm.

### Options Attributes

`options` exports the following attributes:

* `password_history` - List(Resource). Configuration settings for the password history that is maintained for each user to prevent the reuse of passwords. For details, see [Password History Attributes](#password-attributes-history).

#### Password History Attributes

`password_history` exports the following attributes:

* `enable` - Boolean. Indicates whether password history is enabled for the connection. When enabled, any existing users in this connection will be unaffected; the system will maintain their password history going forward.
* `size` - Integer. Indicates the number of passwords to keep in history.

### Import

Connections can be imported using their id, e.g.

```
$ terraform import auth0_connection.google con_a17f21fdb24d48a0
```
