package auth0

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccHook(t *testing.T) {

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"auth0": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: testAccHookCreate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("auth0_hook.my_hook", "name", "pre-user-reg-hook"),
					resource.TestCheckResourceAttr("auth0_hook.my_hook", "dependencies.#", "0"),
					resource.TestCheckResourceAttr("auth0_hook.my_hook", "script", "function (user, context, callback) { callback(null, { user }); }"),
					resource.TestCheckResourceAttr("auth0_hook.my_hook", "trigger_id", "pre-user-registration"),
					resource.TestCheckResourceAttr("auth0_hook.my_hook", "enabled", "true"),
				),
			},
			{
				Config: testAccHookUpdate,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("auth0_hook.my_hook", "name", "pre-user-reg-hook"),
					resource.TestCheckResourceAttr("auth0_hook.my_hook", "dependencies.#", "1"),
					resource.TestCheckResourceAttr("auth0_hook.my_hook", "dependencies.0.auth0", "2.30.0"), // TODO figure out the right value
					resource.TestCheckResourceAttr("auth0_hook.my_hook", "script", "function (user, context, callback) { console.log(user); callback(null, { user }); }"),
					resource.TestCheckResourceAttr("auth0_hook.my_hook", "trigger_id", "pre-user-registration"),
					resource.TestCheckResourceAttr("auth0_hook.my_hook", "enabled", "false"),
				),
			},
		},
	})
}

const testAccHookCreate = `

resource "auth0_hook" "my_hook" {
  name = "pre-user-reg-hook"
  trigger_id = "pre-user-registration"
  script = "function (user, context, callback) { callback(null, { user }); }"
  enabled = true
	dependencies {}
}
`

const testAccHookUpdate = `

resource "auth0_hook" "my_hook" {
  name = "pre-user-reg-hook"
  trigger_id = "pre-user-registration"
  script = "function (user, context, callback) { console.log(user); callback(null, { user }); }"
  enabled = false
	dependencies {
		auth0 = "2.30.0"
	}
}
`

func TestHookNameRegexp(t *testing.T) {
	for name, valid := range map[string]bool{
		"my-hook-1":                 true,
		"hook 2 name with spaces":   true,
		" hook with a space prefix": false,
		"hook with a space suffix ": false,
		" ":                         false,
		"   ":                       false,
	} {
		fn := validateHookNameFunc()

		_, errs := fn(name, "name")
		if errs != nil && valid {
			t.Fatalf("Expected %q to be valid, but got validation errors %v", name, errs)
		}

		if errs == nil && !valid {
			t.Fatalf("Expected %q to be invalid, but got no validation errors.", name)
		}
	}
}
