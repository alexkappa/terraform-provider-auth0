---
layout: "auth0"
page_title: "Auth0: auth0_hook"
description: |-
  Hooks are secure, self-contained functions that allow you to customize the behavior of Auth0 when executed for selected extensibility points of the Auth0 platform. Auth0 invokes Hooks during runtime to execute your custom Node.js code.

  Depending on the extensibility point, you can use Hooks with Database Connections and/or Passwordless Connections.
---

# auth0_hook

Hooks are secure, self-contained functions that allow you to customize the behavior of Auth0 when executed for selected extensibility points of the Auth0 platform. Auth0 invokes Hooks during runtime to execute your custom Node.js code.

Depending on the extensibility point, you can use Hooks with Database Connections and/or Passwordless Connections.

## Example Usage

```hcl
resource "auth0_hook" "my_hook" {
  name = "My Pre User Registration Hook"
  script = <<EOF
function (user, context, callback) {
  callback(null, { user });
}
EOF
  trigger_id = "pre-user-registration"
  enabled = true

  secrets = {
    foo = "bar"    
  }

  dependencies = {
    auth0 = "2.30.0"
  }

  secrets = {
    secretKey1 = secretValue1
    secretKey2 = secretValue2
  }
}
```

## Argument Reference

The following arguments are supported:

* `enabled` - (Optional) Whether the hook is enabled, or disabled
* `name` - (Required) Name of this hook
* `script` - (Required) Code to be executed when this hook runs
* `trigger_id` - (Required) Execution stage of this rule. Can be credentials-exchange, pre-user-registration, post-user-registration, post-change-password, or send-phone-message
* `dependencies` - (Optional) Dependencies of this hook used by webtask server
* `secrets` - (Optional) Map(String), sets the hook secrets associated with this hook. If specified, any secrets not listed here will be removed from the hook.
