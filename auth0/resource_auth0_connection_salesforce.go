package auth0

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func newSalesforceConnection(strategy string) *schema.Resource {
	return &schema.Resource{
		Create: createConnectionWithStrategy(strategy),
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
					"client_id": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "",
					},
					"client_secret": {
						Type:        schema.TypeString,
						Optional:    true,
						Sensitive:   true,
						Description: "",
					},
					"community_base_url": {
						Type:     schema.TypeString,
						Optional: true,
					},
					"scopes": {
						Type:     schema.TypeSet,
						Optional: true,
						Elem:     &schema.Schema{Type: schema.TypeString},
					},
				},
			},
			Description: "Configuration settings for connection options",
		}),
	}
}
