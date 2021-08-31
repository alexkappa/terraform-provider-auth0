package auth0

import (
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"gopkg.in/auth0.v5/management"
)

func newBranding() *schema.Resource {
	return &schema.Resource{

		Create: createBranding,
		Read:   readBranding,
		Update: updateBranding,
		Delete: deleteBranding,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"colors": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"primary": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"page_background": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"favicon_url": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"logo_url": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"font": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"url": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"universal_login": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"body": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func createBranding(d *schema.ResourceData, m interface{}) error {
	d.SetId(resource.UniqueId())
	return updateBranding(d, m)
}

func readBranding(d *schema.ResourceData, m interface{}) error {
	api := m.(*management.Management)
	b, err := api.Branding.Read()
	if err != nil {
		if mErr, ok := err.(management.Error); ok {
			if mErr.Status() == http.StatusNotFound {
				d.SetId("")
				return nil
			}
		}
		return err
	}

	d.Set("favicon_url", b.FaviconURL)
	d.Set("logo_url", b.LogoURL)
	d.Set("colors", flattenBrandingColors(b.Colors))
	d.Set("font", flattenBrandingFont(b.Font))

	t, err := api.Tenant.Read()
	if err != nil {
		return err
	}

	if t.Flags.EnableCustomDomainInEmails != nil && *t.Flags.EnableCustomDomainInEmails {
		if err := assignUniversalLogin(d, m); err != nil {
			d.SetId("")
			return err
		}
	}

	return nil
}

func updateBranding(d *schema.ResourceData, m interface{}) error {
	b := buildBranding(d)
	ul := buildBrandingUniversalLogin(d)
	api := m.(*management.Management)
	err := api.Branding.Update(b)
	if err != nil {
		return err
	}

	if ul.GetBody() != "" {
		err = api.Branding.SetUniversalLogin(ul)
		if err != nil {
			return err
		}
	}
	return readBranding(d, m)
}

func deleteBranding(d *schema.ResourceData, m interface{}) error {
	api := m.(*management.Management)
	t, err := api.Tenant.Read()
	if err != nil {
		return err
	}

	if t.Flags.EnableCustomDomainInEmails != nil && *t.Flags.EnableCustomDomainInEmails {
		err = api.Branding.DeleteUniversalLogin()
		if err != nil {
			if mErr, ok := err.(management.Error); ok {
				if mErr.Status() == http.StatusNotFound {
					d.SetId("")
					return nil
				}
			}
		}
	}

	return err
}

func buildBranding(d *schema.ResourceData) *management.Branding {
	b := &management.Branding{
		FaviconURL: String(d, "favicon_url"),
		LogoURL:    String(d, "logo_url"),
	}

	List(d, "colors").Elem(func(d ResourceData) {
		b.Colors = &management.BrandingColors{
			PageBackground: String(d, "page_background"),
			Primary:        String(d, "primary"),
		}
	})

	List(d, "font").Elem(func(d ResourceData) {
		b.Font = &management.BrandingFont{
			URL: String(d, "url"),
		}
	})

	return b
}

func buildBrandingUniversalLogin(d *schema.ResourceData) *management.BrandingUniversalLogin {
	b := &management.BrandingUniversalLogin{}

	List(d, "universal_login").Elem(func(d ResourceData) {
		b.Body = String(d, "body")
	})

	return b
}

func assignUniversalLogin(d *schema.ResourceData, m interface{}) error {
	api := m.(*management.Management)
	ul, err := api.Branding.UniversalLogin()
	if err != nil {
		if mErr, ok := err.(management.Error); ok {
			if mErr.Status() == http.StatusNotFound {
				return nil
			}
		}
		return err
	}

	d.Set("universal_login", flattenBrandingUniversalLogin(ul))
	return nil
}

func flattenBrandingColors(brandingColors *management.BrandingColors) []interface{} {
	m := make(map[string]interface{})
	if brandingColors != nil {
		m["page_background"] = brandingColors.PageBackground
		m["primary"] = brandingColors.Primary
	}
	return []interface{}{m}
}

func flattenBrandingUniversalLogin(brandingUniversalLogin *management.BrandingUniversalLogin) []interface{} {
	m := make(map[string]interface{})
	if brandingUniversalLogin != nil {
		m["body"] = brandingUniversalLogin.Body
	}
	return []interface{}{m}
}

func flattenBrandingFont(brandingFont *management.BrandingFont) []interface{} {
	m := make(map[string]interface{})
	if brandingFont != nil {
		m["url"] = brandingFont.URL
	}
	return []interface{}{m}
}
