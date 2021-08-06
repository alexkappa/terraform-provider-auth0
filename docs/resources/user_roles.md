---
layout: "auth0"
page_title: "Auth0: auth0_user_roles"
description: |-
  With this resource, you can create and manage collections of roles that are assigned to users.
---

# auth0_role

With this resource, you can create and manage collections of roles that are assigned to users. Permissions (scopes) are created on auth0_resource_server, then associated with roles and optionally, users using this resource.

## Example Usage

```hcl
resource "auth0_resource_server" "my_resource_server" {
  name = "My Resource Server (Managed by Terraform)"
  identifier = "my-resource-server-identifier"
  signing_alg = "RS256"
  token_lifetime = 86400
  skip_consent_for_verifiable_first_party_clients = true

  enforce_policies = true

  scopes {
    value = "read:something"
    description = "read something"
  }
}

resource "auth0_role" "my_role" {
  name = "My Role - (Managed by Terraform)"
  description = "Role Description..."

  permissions {
    resource_server_identifier = "${auth0_resource_server.my_resource_server.identifier}"
    name = "read:something"
  }
}

data "auth0_user" "my_user" {
  email = "test@test.com"
}

resource "auth0_user_roles" "my_user" {
  user_id = data.auth0_user.my_user.user_id
  roles = [ "${auth0_role.my_role.id}" ]
}
```

~> The user_roles resource should be used when you are not managing your users with a user resource. When managing users with user resources you should set the roles directly on the user resource.

## Argument Reference

Arguments accepted by this resource include:

* `user_id` - (Required) List(String). IDs of the users to which the role is assigned.
* `roles` - (Required) Set(String). Set of IDs of roles assigned to the user.

## Attribute Reference

Attributes exported by this resource include:

* `id` - String. ID for the role.
