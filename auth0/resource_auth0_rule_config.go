package auth0

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"gopkg.in/auth0.v2"
	"gopkg.in/auth0.v2/management"
)

func newRuleConfig() *schema.Resource {
	return &schema.Resource{

		Create: createRuleConfig,
		Read:   readRuleConfig,
		Update: updateRuleConfig,
		Delete: deleteRuleConfig,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"key": {
				Type:     schema.TypeString,
				Required: true,
			},
			"value": {
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
			},
		},
	}
}

func createRuleConfig(d *schema.ResourceData, m interface{}) error {
	r := buildRuleConfig(d)
	key := auth0.StringValue(r.Key)
	r.Key = nil
	api := m.(*management.Management)
	if err := api.RuleConfig.Upsert(key, r); err != nil {
		return err
	}
	d.SetId(auth0.StringValue(r.Key))
	return readRuleConfig(d, m)
}

func readRuleConfig(d *schema.ResourceData, m interface{}) error {
	api := m.(*management.Management)
	r, err := api.RuleConfig.Read(d.Id())
	if err != nil {
		return err
	}
	d.Set("key", r.Key)
	return nil
}

func updateRuleConfig(d *schema.ResourceData, m interface{}) error {
	r := buildRuleConfig(d)
	r.Key = nil
	api := m.(*management.Management)
	err := api.RuleConfig.Upsert(d.Id(), r)
	if err != nil {
		return err
	}
	return readRuleConfig(d, m)
}

func deleteRuleConfig(d *schema.ResourceData, m interface{}) error {
	api := m.(*management.Management)
	return api.RuleConfig.Delete(d.Id())
}

func buildRuleConfig(d *schema.ResourceData) *management.RuleConfig {
	return &management.RuleConfig{
		Key:   String(d, "key"),
		Value: String(d, "value"),
	}
}
