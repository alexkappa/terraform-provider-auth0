package auth0

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccRuleConfig(t *testing.T) {

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"auth0": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: testAccRuleConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("auth0_rule_config.foo", "id", "foo"),
					resource.TestCheckResourceAttr("auth0_rule_config.foo", "key", "foo"),
					resource.TestCheckResourceAttr("auth0_rule_config.foo", "value", "bar"),
				),
			},
		},
	})
}

const testAccRuleConfig = `
provider "auth0" {}

resource "auth0_rule_config" "foo" {
  key = "foo"
  value = "bar"
}
`
