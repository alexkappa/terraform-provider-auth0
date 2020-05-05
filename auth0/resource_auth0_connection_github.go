package auth0

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"gopkg.in/auth0.v4/management"
)

func newGitHubConnection() *schema.Resource {
	return &schema.Resource{
		Create: createConnectionWithStrategy(management.ConnectionStrategyGitHub),
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
