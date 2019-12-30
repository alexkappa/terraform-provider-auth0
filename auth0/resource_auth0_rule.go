package auth0

import (
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"

	"gopkg.in/auth0.v2"
	"gopkg.in/auth0.v2/management"
)

var ruleNameRegexp = regexp.MustCompile("^[^\\s-][\\w -]+[^\\s-]$")

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
					ruleNameRegexp,
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
	d.SetId(auth0.StringValue(c.ID))
	return readRule(d, m)
}

func readRule(d *schema.ResourceData, m interface{}) error {
	api := m.(*management.Management)
	c, err := api.Rule.Read(d.Id())
	if err != nil && strings.HasPrefix(err.Error(), "404") {
		d.SetId("")
		return nil
	}
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
	err := api.Rule.Delete(d.Id())
	if err != nil && strings.HasPrefix(err.Error(), "404") {
		d.SetId("")
		return nil
	}
	return err
}

func buildRule(d *schema.ResourceData) *management.Rule {
	return &management.Rule{
		Name:    String(d, "name"),
		Script:  String(d, "script"),
		Order:   Int(d, "order"),
		Enabled: Bool(d, "enabled"),
	}
}
