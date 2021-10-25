package auth0

import (
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"gopkg.in/auth0.v5/management"
)

func newOrganization() *schema.Resource {
	return &schema.Resource{

		Create: createOrganization,
		Read:   readOrganization,
		Update: updateOrganization,
		Delete: deleteOrganization,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of this organization",
			},
			"display_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Friendly name of this organization",
			},
			"branding": {
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				MaxItems:    1,
				MinItems:    1,
				Description: "Defines how to style the login pages",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"logo_url": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"colors": {
							Type:     schema.TypeMap,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
					},
				},
			},
			"metadata": {
				Type:        schema.TypeMap,
				Optional:    true,
				Elem:        &schema.Schema{Type: schema.TypeString},
				Description: "Metadata associated with the organization, Maximum of 10 metadata properties allowed",
			},
		},
	}
}

func createOrganization(d *schema.ResourceData, m interface{}) error {
	o := expandOrganization(d)
	api := m.(*management.Management)
	if err := api.Organization.Create(o); err != nil {
		return err
	}
	d.SetId(o.GetID())
	return readOrganization(d, m)
}

func readOrganization(d *schema.ResourceData, m interface{}) error {
	api := m.(*management.Management)
	o, err := api.Organization.Read(d.Id())
	if err != nil {
		if mErr, ok := err.(management.Error); ok {
			if mErr.Status() == http.StatusNotFound {
				d.SetId("")
				return nil
			}
		}
		return err
	}

	d.SetId(o.GetID())
	d.Set("name", o.Name)
	d.Set("display_name", o.DisplayName)
	d.Set("branding", flattenOrganizationBranding(o.Branding))
	d.Set("metadata", o.Metadata)

	return nil
}

func updateOrganization(d *schema.ResourceData, m interface{}) error {
	o := expandOrganization(d)
	api := m.(*management.Management)
	err := api.Organization.Update(d.Id(), o)
	if err != nil {
		return err
	}
	return readOrganization(d, m)
}

func deleteOrganization(d *schema.ResourceData, m interface{}) error {
	api := m.(*management.Management)
	err := api.Organization.Delete(d.Id())
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

func expandOrganization(d *schema.ResourceData) *management.Organization {
	o := &management.Organization{
		Name:        String(d, "name"),
		DisplayName: String(d, "display_name"),
		Metadata:    Map(d, "metadata"),
	}
	List(d, "branding").Elem(func(d ResourceData) {
		o.Branding = &management.OrganizationBranding{
			LogoURL: String(d, "logo_url"),
			Colors:  Map(d, "colors"),
		}
	})
	return o
}

func flattenOrganizationBranding(b *management.OrganizationBranding) []interface{} {
	m := make(map[string]interface{})
	if b != nil {
		m["logo_url"] = b.LogoURL
		m["colors"] = b.Colors
	}
	return []interface{}{m}
}
