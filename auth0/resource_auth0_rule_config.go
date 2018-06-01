package auth0

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/yieldr/go-auth0/management"
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
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func createRuleConfig(d *schema.ResourceData, m interface{}) error {
	r := buildRuleConfig(d)
	key := r.Key
	r.Key = ""
	api := m.(*management.Management)
	if err := api.RuleConfig.Upsert(key, r); err != nil {
		return err
	}
	d.SetId(r.Key)
	return nil
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
	r.Key = ""
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
		Key:   d.Get("key").(string),
		Value: d.Get("value").(string),
	}
}
