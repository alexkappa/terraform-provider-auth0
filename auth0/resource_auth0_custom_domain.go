package auth0

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
	"github.com/yieldr/go-auth0/management"
)

func newCustomDomain() *schema.Resource {
	return &schema.Resource{

		Create: createCustomDomain,
		Read:   readCustomDomain,
		Update: updateCustomDomain,
		Delete: deleteCustomDomain,
		Importer: &schema.ResourceImporter{
			State: func(data *schema.ResourceData, i interface{}) ([]*schema.ResourceData, error) {
				results := make([]*schema.ResourceData, 1, 1)
				results[0] = data
				readCustomDomain(data, i)
				return results, nil
			},
		},

		Schema: map[string]*schema.Schema{
			"domain": {
				Type:     schema.TypeString,
				Required: true,
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.StringInSlice([]string{
					"auth0_managed_certs",
					"self_managed_certs",
				}, true),
			},
			"primary": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"verification_method": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"txt"}, true),
			},
			"verification": {
				Type:     schema.TypeList,
				Computed: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"methods": {
							Type:     schema.TypeList,
							Elem:     schema.TypeMap,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func createCustomDomain(d *schema.ResourceData, m interface{}) error {
	c := buildCustomDomain(d)
	api := m.(*management.Management)
	if err := api.CustomDomain.Create(c); err != nil {
		return err
	}
	d.SetId(c.ID)
	verification := []map[string]interface{}{
		{"methods": c.Verification.Methods},
	}
	if err := d.Set("verification", verification); err != nil {
		return err
	}
	return nil
}

func readCustomDomain(d *schema.ResourceData, m interface{}) error {
	api := m.(*management.Management)
	c, err := api.CustomDomain.Read(d.Id())
	if err != nil {
		return err
	}
	d.SetId(c.ID)
	return nil
}

func updateCustomDomain(d *schema.ResourceData, m interface{}) error {
	c := buildCustomDomain(d)
	c.Verification = nil
	api := m.(*management.Management)
	err := api.CustomDomain.Update(d.Id(), c)
	if err != nil {
		return err
	}
	return nil
}

func deleteCustomDomain(d *schema.ResourceData, m interface{}) error {
	api := m.(*management.Management)
	return api.CustomDomain.Delete(d.Id())
}

func buildCustomDomain(d *schema.ResourceData) *management.CustomDomain {

	c := &management.CustomDomain{
		Domain:             d.Get("domain").(string),
		Type:               d.Get("type").(string),
		VerificationMethod: d.Get("verification_method").(string),
	}

	return c
}
