package auth0

import (
	"testing"

	"github.com/alexkappa/terraform-provider-auth0/auth0/internal/random"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccAction(t *testing.T) {

	rand := random.String(6)

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"auth0": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: random.Template(testAccActionConfigCreate, rand),
				Check: resource.ComposeTestCheckFunc(
					random.TestCheckResourceAttr("auth0_action.my_action", "name", "Test Action {{.random}}", rand),
					resource.TestCheckResourceAttr("auth0_action.my_action", "code", "exports.onExecutePostLogin = async (event, api) => {};"),
					resource.TestCheckResourceAttr("auth0_action.my_action", "secrets.#", "1"),
				),
			},
			{
				Config: random.Template(testAccActionConfigUpdate, rand),
				Check: resource.ComposeTestCheckFunc(
					random.TestCheckResourceAttr("auth0_action.my_action", "name", "Test Action {{.random}}", rand),
					resource.TestCheckResourceAttr("auth0_action.my_action", "code", "exports.onContinuePostLogin = async (event, api) => {};"),
					resource.TestCheckResourceAttrSet("auth0_action.my_action", "version_id"),
					resource.TestCheckResourceAttr("auth0_action.my_action", "secrets.#", "1"),
				),
			},
			{
				Config: random.Template(testAccActionConfigUpdateAgain, rand),
				Check: resource.ComposeTestCheckFunc(
					random.TestCheckResourceAttr("auth0_action.my_action", "name", "Test Action {{.random}}", rand),
					resource.TestCheckResourceAttrSet("auth0_action.my_action", "version_id"),
					resource.TestCheckResourceAttr("auth0_action.my_action", "secrets.#", "0"),
				),
			},
		},
	})
}

const testAccActionConfigCreate = `

resource auth0_action my_action {
	name = "Test Action {{.random}}"
	supported_triggers {
		id = "post-login"
		version = "v2"
	}
	secrets {
		name = "foo"
		value = "123"
	}
	code = "exports.onExecutePostLogin = async (event, api) => {};"
}
`

const testAccActionConfigUpdate = `

resource auth0_action my_action {
	name = "Test Action {{.random}}"
	supported_triggers {
		id = "post-login"
		version = "v2"
	}
	secrets {
		name = "foo"
		value = "123456"
	}
	code = "exports.onContinuePostLogin = async (event, api) => {};"
	deploy = true
}
`

const testAccActionConfigUpdateAgain = `

resource auth0_action my_action {
	name = "Test Action {{.random}}"
	supported_triggers {
		id = "post-login"
		version = "v2"
	}
	code = <<-EOT
	exports.onContinuePostLogin = async (event, api) => { 
		console.log(event) 
	};"
	EOT
	deploy = true
}
`
