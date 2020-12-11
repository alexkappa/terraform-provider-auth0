package auth0

import (
	"log"
	"strings"
	"testing"

	"github.com/alexkappa/terraform-provider-auth0/auth0/internal/random"
	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func init() {
	resource.AddTestSweepers("auth0_rule_config", &resource.Sweeper{
		Name: "auth0_rule_config",
		F: func(_ string) error {
			api, err := Auth0()
			if err != nil {
				return err
			}
			configurations, err := api.RuleConfig.List()
			if err != nil {
				return err
			}
			for _, c := range configurations {
				log.Printf("[DEBUG] ➝ %s", c.GetKey())
				if strings.Contains(c.GetKey(), "test") {
					if e := api.RuleConfig.Delete(c.GetKey()); e != nil {
						multierror.Append(err, e)
					}
					log.Printf("[DEBUG] ✗ %s", c.GetKey())
				}
			}
			return err
		},
	})
}

func TestAccRuleConfig(t *testing.T) {

	rand := random.String(4)

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"auth0": Provider(),
		},
		Steps: []resource.TestStep{
			{
				Config: random.Template(testAccRuleConfigCreate, rand),
				Check: resource.ComposeTestCheckFunc(
					random.TestCheckResourceAttr("auth0_rule_config.foo", "id", "acc_test_{{.random}}", rand),
					random.TestCheckResourceAttr("auth0_rule_config.foo", "key", "acc_test_{{.random}}", rand),
					resource.TestCheckResourceAttr("auth0_rule_config.foo", "value", "bar"),
				),
			},
			{
				Config: random.Template(testAccRuleConfigUpdateValue, rand),
				Check: resource.ComposeTestCheckFunc(
					random.TestCheckResourceAttr("auth0_rule_config.foo", "id", "acc_test_{{.random}}", rand),
					random.TestCheckResourceAttr("auth0_rule_config.foo", "key", "acc_test_{{.random}}", rand),
					resource.TestCheckResourceAttr("auth0_rule_config.foo", "value", "foo"),
				),
			},
			{
				Config: random.Template(testAccRuleConfigUpdateKey, rand),
				Check: resource.ComposeTestCheckFunc(
					random.TestCheckResourceAttr("auth0_rule_config.foo", "id", "acc_test_key_{{.random}}", rand),
					random.TestCheckResourceAttr("auth0_rule_config.foo", "key", "acc_test_key_{{.random}}", rand),
					resource.TestCheckResourceAttr("auth0_rule_config.foo", "value", "foo"),
				),
			},
		},
	})
}

const testAccRuleConfigCreate = `

resource "auth0_rule_config" "foo" {
  key = "acc_test_{{.random}}"
  value = "bar"
}
`

const testAccRuleConfigUpdateValue = `

resource "auth0_rule_config" "foo" {
  key = "acc_test_{{.random}}"
  value = "foo"
}
`

const testAccRuleConfigUpdateKey = `

resource "auth0_rule_config" "foo" {
  key = "acc_test_key_{{.random}}"
  value = "foo"
}
`
