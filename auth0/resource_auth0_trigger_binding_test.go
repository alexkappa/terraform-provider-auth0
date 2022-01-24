package auth0

import (
	"testing"

	"github.com/alexkappa/terraform-provider-auth0/auth0/internal/random"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccTriggerBinding(t *testing.T) {

	rand := random.String(6)

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"auth0": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: random.Template(testAccTriggerBindingConfigCreate, rand),
				Check: resource.ComposeTestCheckFunc(
					random.TestCheckResourceAttr("auth0_action.action_foo", "name", "Test Trigger Binding Foo {{.random}}", rand),
					random.TestCheckResourceAttr("auth0_action.action_bar", "name", "Test Trigger Binding Bar {{.random}}", rand),
					resource.TestCheckResourceAttr("auth0_trigger_binding.login_flow", "actions.#", "2"),
					random.TestCheckResourceAttr("auth0_trigger_binding.login_flow", "actions.0.display_name", "Test Trigger Binding Foo {{.random}}", rand),
					random.TestCheckResourceAttr("auth0_trigger_binding.login_flow", "actions.1.display_name", "Test Trigger Binding Bar {{.random}}", rand),
				),
			},
			{
				Config: random.Template(testAccTriggerBindingConfigUpdate, rand),
				Check: resource.ComposeTestCheckFunc(
					random.TestCheckResourceAttr("auth0_action.action_foo", "name", "Test Trigger Binding Foo {{.random}}", rand),
					random.TestCheckResourceAttr("auth0_action.action_bar", "name", "Test Trigger Binding Bar {{.random}}", rand),
					resource.TestCheckResourceAttr("auth0_trigger_binding.login_flow", "actions.#", "2"),
					random.TestCheckResourceAttr("auth0_trigger_binding.login_flow", "actions.0.display_name", "Test Trigger Binding Bar {{.random}}", rand),
					random.TestCheckResourceAttr("auth0_trigger_binding.login_flow", "actions.1.display_name", "Test Trigger Binding Foo {{.random}}", rand),
				),
			},
		},
	})
}

const testAccTriggerBindingAction = `

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
`

const testAccTriggerBindingConfigCreate = testAccTriggerBindingAction + `

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
`

const testAccTriggerBindingConfigUpdate = testAccTriggerBindingAction + `

resource auth0_trigger_binding login_flow {
	trigger = "post-login"
	actions {
		id = auth0_action.action_bar.id # <----- change the order of the actions
		display_name = auth0_action.action_bar.name
	}
	actions {
		id = auth0_action.action_foo.id
		display_name = auth0_action.action_foo.name
	}
}
`
