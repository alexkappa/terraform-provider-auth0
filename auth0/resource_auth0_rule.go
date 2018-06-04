package auth0

import (
	"regexp"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
	"github.com/yieldr/go-auth0/management"
)

func newRule() *schema.Resource {
	return &schema.Resource{

		Create: createRule,
		Read:   readRule,
		Update: updateRule,
		Delete: deleteRule,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.StringMatch(
					regexp.MustCompile("^[^\\s-][\\w-]+[^\\s-]$"),
					"Can only contain alphanumeric characters, spaces and '-'. "+
						"Can neither start nor end with '-' or spaces."),
			},
			"script": {
				Type:     schema.TypeString,
				Required: true,
			},
			"order": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"enabled": {
				Type:     schema.TypeBool,
				Optional: true,
			},
		},
	}
}

func createRule(d *schema.ResourceData, m interface{}) error {
	c := buildRule(d)
	api := m.(*management.Management)
	if err := api.Rule.Create(c); err != nil {
		return err
	}
	d.SetId(c.ID)
	return readRule(d, m)
}

func readRule(d *schema.ResourceData, m interface{}) error {
	api := m.(*management.Management)
	c, err := api.Rule.Read(d.Id())
	if err != nil {
		return err
	}
	d.Set("name", c.Name)
	d.Set("script", c.Script)
	d.Set("order", c.Order)
	d.Set("enabled", c.Enabled)
	return nil
}

func updateRule(d *schema.ResourceData, m interface{}) error {
	c := buildRule(d)
	api := m.(*management.Management)
	err := api.Rule.Update(d.Id(), c)
	if err != nil {
		return err
	}
	return readRule(d, m)
}

func deleteRule(d *schema.ResourceData, m interface{}) error {
	api := m.(*management.Management)
	return api.Rule.Delete(d.Id())
}

func buildRule(d *schema.ResourceData) *management.Rule {
	return &management.Rule{
		Name:    d.Get("name").(string),
		Script:  d.Get("script").(string),
		Order:   d.Get("order").(int),
		Enabled: d.Get("enabled").(bool),
	}
}
