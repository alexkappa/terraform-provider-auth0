package auth0

import (
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"

	"gopkg.in/auth0.v4"
	"gopkg.in/auth0.v4/management"
)

func newCustomDomain() *schema.Resource {
	return &schema.Resource{

		Create: createCustomDomain,
		Read:   readCustomDomain,
		Update: updateCustomDomain,
		Delete: deleteCustomDomain,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"domain": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
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
				ForceNew:     true,
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
			"verification_trigger": {
				Type:     schema.TypeMap,
				Optional: true,
			},
		},
	}
}

func createCustomDomain(d *schema.ResourceData, m interface{}) error {
	c := expandCustomDomain(d)
	api := m.(*management.Management)
	if err := api.CustomDomain.Create(c); err != nil {
		return err
	}
	d.SetId(auth0.StringValue(c.ID))
	return readCustomDomain(d, m)
}

func readCustomDomain(d *schema.ResourceData, m interface{}) error {
	api := m.(*management.Management)
	c, err := api.CustomDomain.Read(d.Id())
	if err != nil {
		if mErr, ok := err.(management.Error); ok {
			if mErr.Status() == http.StatusNotFound {
				d.SetId("")
				return nil
			}
		}
		return err
	}

	d.SetId(auth0.StringValue(c.ID))
	d.Set("domain", c.Domain)
	d.Set("type", c.Type)
	d.Set("primary", c.Primary)
	d.Set("status", c.Status)

	if c.Verification != nil {
		d.Set("verification", []map[string]interface{}{
			{"methods": c.Verification.Methods},
		})
	}

	return nil
}

func updateCustomDomain(d *schema.ResourceData, m interface{}) error {
	// If "verification_trigger" was updated, run a verification request. Since
	// most fields either force a new resource or get computed server side, no
	// other updates are needed.
	err := verifyCustomDomain(d, m)
	if err != nil {
		return err
	}
	return readCustomDomain(d, m)
}

func deleteCustomDomain(d *schema.ResourceData, m interface{}) error {
	api := m.(*management.Management)
	err := api.CustomDomain.Delete(d.Id())
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

func expandCustomDomain(d *schema.ResourceData) *management.CustomDomain {
	return &management.CustomDomain{
		Domain:             String(d, "domain"),
		Type:               String(d, "type"),
		VerificationMethod: String(d, "verification_method"),
	}
}

func verifyCustomDomain(d *schema.ResourceData, m interface{}) error {
	api := m.(*management.Management)
	if d.HasChange("verification_trigger") {
		v, err := api.CustomDomain.Verify(d.Id())
		if err != nil {
			return err
		}
		d.Set("status", v.Status)
	}
	return nil
}
