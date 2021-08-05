---
layout: "auth0"
page_title: "Data Source: auth0_user"
description: |-
  Use this data source to get information about a user
---

# Data Source: auth0_user

Use this data source to get information about a user

## Example Usage

```data "auth0_user" "user" {
  email = "test@test.com"
}

data "auth0_user" "example" {
  email           = "test@gmail.com"
  connection_name = "google-oauth2"
}
```

## Argument Reference

Arguments accepted by this resource include:

- `user_id` - (Optional) String. ID of the user.
- `email` - (Optional) String. Email address of the user.
- `connection_name` - (Optional) String. Name of the connection from which the user information was sourced.

## Attributes Reference

The following attributes are exported:

- `user_id` - (Optional) String. ID of the user.
- `connection_name` - (Required) String. Name of the connection from which the user information was sourced.
- `username` - (Optional) String. Username of the user. Only valid if the connection requires a username.
- `nickname` - (Optional) String. Preferred nickname or alias of the user.
- `email` - (Optional) String. Email address of the user.
- `email_verified` - (Optional) Boolean. Indicates whether or not the email address has been verified.
- `verify_email` - (Optional) Boolean. Indicates whether or not the user will receive a verification email after creation. Overrides behavior of `email_verified` parameter.
- `phone_number` - (Optional) String. Phone number for the user; follows the E.164 recommendation. Used for SMS connections.
- `phone_verified` - (Optional) Boolean. Indicates whether or not the phone number has been verified.
- `user_metadata` - (Optional) String, JSON format. Custom fields that store info about the user that does not impact a user's core functionality. Examples include work address, home address, and user preferences.
- `app_metadata` (Optional) String, JSON format. Custom fields that store info about the user that impact the user's core functionality, such as how an application functions or what the user can access. Examples include support plans and IDs for external accounts.
- `created_at` - (Optional) String. Date where user was created.
- `updated_at` - (Optional) String. Last date where user was updated.
- `identities` - (Optional) String. All connections where user is.

### Identities Attributes

- `connection` - (Optional) String. Name of the connection from which the user information was sourced.
- `user_id` - (Optional) String. ID of the user.
- `provider` - (Optional) String. Type of the connection, which indicates the identity provider.
- `is_social` - (Optional) String. If it's a social provider.
