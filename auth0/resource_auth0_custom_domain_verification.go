package auth0

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"gopkg.in/auth0.v5/management"
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

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
		},
	}
}

func createCustomDomainVerification(d *schema.ResourceData, m interface{}) error {
	api := m.(*management.Management)
	return resource.Retry(d.Timeout(schema.TimeoutCreate), func() *resource.RetryError {
		c, err := api.CustomDomain.Verify(d.Get("custom_domain_id").(string))
		if err != nil {
			return resource.NonRetryableError(err)
		}
		if c.GetStatus() != "ready" {
			return resource.RetryableError(fmt.Errorf("Custom domain has status %q", c.GetStatus()))
		}
		log.Printf("[INFO] Custom domain %s verified", c.GetDomain())
		d.SetId(c.GetID())
		return resource.NonRetryableError(readCustomDomainVerification(d, m))
	})
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
	d.Set("custom_domain_id", c.GetID())
	return nil
}

func deleteCustomDomainVerification(d *schema.ResourceData, m interface{}) error {
	return nil
}
