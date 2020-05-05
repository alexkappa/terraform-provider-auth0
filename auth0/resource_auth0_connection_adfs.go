package auth0

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func newADFSConnection() *schema.Resource {
	return &schema.Resource{
		Create: createConnectionWithStrategy("adfs"),
		Read:   readConnection,
		Update: updateConnection,
		Delete: deleteConnection,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: newConnectionSchema(&schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			MaxItems: 1,
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"tenant_domain": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "",
					},
					"domain_aliases": {
						Type:        schema.TypeSet,
						Elem:        &schema.Schema{Type: schema.TypeString},
						Optional:    true,
						Description: "",
					},
					"icon_url": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "",
					},
					"api_enable_users": {
						Type:     schema.TypeBool,
						Optional: true,
					},
					"adfs_server": {
						Type:     schema.TypeString,
						Optional: true,
					},
				},
			},
			Description: "Configuration settings for connection options",
		}),
	}
}
