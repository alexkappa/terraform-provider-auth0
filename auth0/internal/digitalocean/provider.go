package digitalocean

import (
	"github.com/digitalocean/godo"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// Provider returns a schema.Provider for a minimal version of the DigitalOcean
// provider used for testing.
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"token": {
				Type:     schema.TypeString,
				Optional: true,
				DefaultFunc: schema.MultiEnvDefaultFunc([]string{
					"DIGITALOCEAN_TOKEN",
					"DIGITALOCEAN_ACCESS_TOKEN",
				}, nil),
				Description: "The token key for API operations.",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"digitalocean_record": resourceDigitalOceanRecord(),
		},
		ConfigureFunc: func(d *schema.ResourceData) (interface{}, error) {
			return godo.NewFromToken(d.Get("token").(string)), nil
		},
	}
}
