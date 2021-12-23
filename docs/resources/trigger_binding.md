---
layout: "auth0"
page_title: "Auth0: auth0_trigger_binding"
description: |-
	With this resource, you can bind an action to a trigger. Once an action is
	created and deployed, it can be attached (i.e. bound) to a trigger so that it
	will be executed as part of a flow.
---

# auth0_trigger_binding

With this resource, you can bind an action to a trigger. Once an action is
created and deployed, it can be attached (i.e. bound) to a trigger so that it
will be executed as part of a flow. 

The list of actions reflects the order in which they will be executed during the
appropriate flow.

## Example Usage

```hcl
resource auth0_action action_foo {
	name = "Test Trigger Binding Foo {{.random}}"
	supported_triggers {
		id = "post-login"
		version = "v2"
	}
	code = <<-EOT
	exports.onContinuePostLogin = async (event, api) => { 
		console.log("foo") 
	};"
	EOT
	deploy = true
}

resource auth0_action action_bar {
	name = "Test Trigger Binding Bar {{.random}}"
	supported_triggers {
		id = "post-login"
		version = "v2"
	}
	code = <<-EOT
	exports.onContinuePostLogin = async (event, api) => { 
		console.log("bar") 
	};"
	EOT
	deploy = true
}

resource auth0_trigger_binding login_flow {
	trigger = "post-login"
	actions {
		id = auth0_action.action_foo.id
		display_name = auth0_action.action_foo.name
	}
	actions {
		id = auth0_action.action_bar.id
		display_name = auth0_action.action_bar.name
	}
}
```

## Argument Reference

The following arguments are supported:

* `trigger` - (Required) The id of the trigger to bind with
* `actions` - (Required) The actions bound to this trigger. For details, see
  [Actions](#actions).

### Actions

* `id` - (Required) Trigger ID.
* `display_name` - (Required) The name of an action.

## Import

auth0_trigger_binding can be imported using the bindings trigger ID, e.g.

```
$ terraform import auth0_trigger_binding.example "post-login"
```
