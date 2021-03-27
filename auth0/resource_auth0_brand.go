package auth0

import (
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"gopkg.in/auth0.v5/management"
)

func newBrand() *schema.Resource {
	return &schema.Resource{

		Create: createBrand,
		Read:   readBrand,
		Update: updateBrand,
		Delete: deleteBrand,
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
		},
	}
}

func createBrand(d *schema.ResourceData, m interface{}) error {
	d.SetId(resource.UniqueId())
	return updateBrand(d, m)
}

func readBrand(d *schema.ResourceData, m interface{}) error {
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
	d.Set("colors", flattenBrandColors(b.Colors))
	d.Set("font", flattenBrandFont(b.Font))

	return nil
}

func updateBrand(d *schema.ResourceData, m interface{}) error {
	b := buildBrand(d)
	api := m.(*management.Management)
	err := api.Branding.Update(b)
	if err != nil {
		return err
	}
	return readBrand(d, m)
}

func deleteBrand(d *schema.ResourceData, m interface{}) error {
	d.SetId("")
	return nil
}

func buildBrand(d *schema.ResourceData) *management.Branding {
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

func flattenBrandColors(brandingColors *management.BrandingColors) []interface{} {
	m := make(map[string]interface{})
	if brandingColors != nil {
		m["page_background"] = brandingColors.PageBackground
		m["primary"] = brandingColors.Primary
	}
	return []interface{}{m}
}

func flattenBrandFont(brandingFont *management.BrandingFont) []interface{} {
	m := make(map[string]interface{})
	if brandingFont != nil {
		m["url"] = brandingFont.URL
	}
	return []interface{}{m}
}
