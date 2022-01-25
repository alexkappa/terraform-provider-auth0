package auth0

import (
	"testing"

	"github.com/alexkappa/terraform-provider-auth0/auth0/internal/random"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccRule(t *testing.T) {

	rand := random.String(6)

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"auth0": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: random.Template(testAccRule, rand),
				Check: resource.ComposeTestCheckFunc(
					random.TestCheckResourceAttr("auth0_rule.my_rule", "name", "acceptance-test-{{.random}}", rand),
					resource.TestCheckResourceAttr("auth0_rule.my_rule", "script", "function (user, context, callback) { callback(null, user, context); }"),
					resource.TestCheckResourceAttr("auth0_rule.my_rule", "enabled", "true"),
				),
			},
		},
	})
}

const testAccRule = `

resource "auth0_rule" "my_rule" {
  name = "acceptance-test-{{.random}}"
  script = "function (user, context, callback) { callback(null, user, context); }"
  enabled = true
}
`

func TestRuleNameRegexp(t *testing.T) {

	vf := validation.StringMatch(ruleNameRegexp, "invalid name")

	for name, valid := range map[string]bool{
		"my-rule-1":                 true,
		"1-my-rule":                 true,
		"rule 2 name with spaces":   true,
		" rule with a space prefix": false,
		"rule with a space suffix ": false,
		" ":                         false,
		"   ":                       false,
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
