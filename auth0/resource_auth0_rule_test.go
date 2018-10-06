package auth0

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/helper/validation"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccRule(t *testing.T) {

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"auth0": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: testAccRule,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("auth0_rule.my_rule", "name", "auth0-authorization-extension"),
					resource.TestCheckResourceAttr("auth0_rule.my_rule", "script", "function (user, context, callback) { callback(null, user, context); }"),
					resource.TestCheckResourceAttr("auth0_rule.my_rule", "enabled", "true"),
				),
			},
		},
	})
}

const testAccRule = `
provider "auth0" {}

resource "auth0_rule" "my_rule" {
  name = "auth0-authorization-extension"
  script = "function (user, context, callback) { callback(null, user, context); }"
  enabled = true
}
`

func TestRuleNameRegexp(t *testing.T) {
	testCases := []struct {
		name  string
		valid bool
	}{
		{
			name:  "my-rule-1",
			valid: true,
		},
		{
			name:  "rule 2 name with spaces",
			valid: true,
		},
		{
			name:  " rule with a space prefix",
			valid: false,
		},
		{
			name:  "rule with a space suffix ",
			valid: false,
		},
		{
			name:  " ", // rule with only one space,
			valid: false,
		},
		{
			name:  "   ", // rule with only three spaces,
			valid: false,
		},
	}

	vf := validation.StringMatch(ruleNameRegexp, "invalid name")
	for _, tc := range testCases {
		_, errs := vf(tc.name, "name")
		if errs != nil && tc.valid {
			t.Fatalf("Expected %q to be valid, but got validation errors %v", tc.name, errs)
		}
		if errs == nil && !tc.valid {
			t.Fatalf("Expected %q to be invalid, but got no validation errors.", tc.name)
		}
	}
}
