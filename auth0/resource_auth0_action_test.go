package auth0

import (
	"testing"

	"github.com/alexkappa/terraform-provider-auth0/auth0/internal/random"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccAction(t *testing.T) {

	rand := random.String(6)

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"auth0": Provider(),
		},
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: random.Template(testAccAction, rand),
				Check: resource.ComposeTestCheckFunc(
					random.TestCheckResourceAttr("auth0_action.my_action", "name", "acceptance-test-{{.random}}", rand),
					resource.TestCheckResourceAttr("auth0_action.my_action", "code", "function (user, context, callback) { callback(null, user, context); }"),
					resource.TestCheckResourceAttr("auth0_action.my_action", "runtime", "node12"),
				),
			},
		},
	})
}

const testAccAction = `

resource "auth0_action" "my_action" {
  name = "acceptance-test-{{.random}}"
  dependencies = [
	  {
		"name": "lodash",
		"version": "1.0.0"
	  }
	]
  code = "function (user, context, callback) { callback(null, user, context); }"
  runtime = "node12"
  secrets = [
	  {
		  "name": "mySecret",
		  "value": "mySecretValue"
	  }
  ]
}
`

func TestActionNameRegexp(t *testing.T) {

	vf := validation.StringMatch(actionNameRegexp, "invalid name")

	for name, valid := range map[string]bool{
		"my-action-1":                 true,
		"1-my-action":                 true,
		"action 2 name with spaces":   true,
		" action with a space prefix": false,
		"action with a space suffix ": false,
		" ":                           false,
		"   ":                         false,
	} {
		_, errs := vf(name, "name")
		if errs != nil && valid {
			t.Fatalf("Expected %q to be valid, but got validation errors %v", name, errs)
		}
		if errs == nil && !valid {
			t.Fatalf("Expected %q to be invalid, but got no validation errors.", name)
		}
	}
}
