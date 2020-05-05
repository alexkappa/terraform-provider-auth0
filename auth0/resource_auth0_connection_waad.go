package auth0

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"gopkg.in/auth0.v4/management"
)

func newAzureADConnection() *schema.Resource {
	return &schema.Resource{
		Create: createAzureADConnection,
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
					"app_id": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "",
					},
					"domain": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "",
					},
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
					"max_groups_to_retrieve": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "",
					},
					"use_wsfed": {
						Type:        schema.TypeBool,
						Optional:    true,
						Description: "",
					},
					"waad_protocol": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "",
					},
					"waad_common_endpoint": {
						Type:        schema.TypeBool,
						Optional:    true,
						Description: "",
					},
					"api_enable_users": {
						Type:     schema.TypeBool,
						Optional: true,
					},
					"identity_api": {
						Type:        schema.TypeString,
						Optional:    true,
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

func createAzureADConnection(d *schema.ResourceData, m interface{}) error {
	d.Set("strategy", management.ConnectionStrategyAzureAD)
	return createConnection(d, m)
}
