package auth0

import (
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"gopkg.in/auth0.v4"
	"gopkg.in/auth0.v4/management"
)

func newCustomDomainVerification() *schema.Resource {
	return &schema.Resource{

		Create: createCustomDomainVerification,
		Read:   readCustomDomainVerification,
		Delete: deleteCustomDomainVerification,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"custom_domain_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func createCustomDomainVerification(d *schema.ResourceData, m interface{}) error {
	api := m.(*management.Management)
	c, err := api.CustomDomain.Verify(d.Get("custom_domain_id").(string))
	if err != nil {
		return err
	}
	d.SetId(auth0.StringValue(c.ID))
	return nil
}

func readCustomDomainVerification(d *schema.ResourceData, m interface{}) error {
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
	d.Set("custom_domain_id", auth0.StringValue(c.ID))
	return nil
}

func deleteCustomDomainVerification(d *schema.ResourceData, m interface{}) error {
	return nil
}
