package auth0

import (
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"gopkg.in/auth0.v4"
	"gopkg.in/auth0.v4/management"
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
		if mErr, ok := err.(management.Error); ok {
			if mErr.Status() == http.StatusNotFound {
				d.SetId("")
				return nil
			}
		}
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
	err := api.RuleConfig.Delete(d.Id())
	if err != nil {
		if mErr, ok := err.(management.Error); ok {
			if mErr.Status() == http.StatusNotFound {
				d.SetId("")
				return nil
			}
		}
	}
	return err
}

func buildRuleConfig(d *schema.ResourceData) *management.RuleConfig {
	return &management.RuleConfig{
		Key:   StringIfExists(d, "key"),
		Value: StringIfExists(d, "value"),
	}
}
