package auth0

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func newConnectionSchema(optionsSchema *schema.Schema) map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"strategy": {
			Type:        schema.TypeString,
			Computed:    true,
			Description: "Type of the connection, which indicates the identity provider",
		},
		"name": {
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
			Description: "Name of the connection",
		},
		"display_name": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Name used in login screen",
		},
		"is_domain_connection": {
			Type:        schema.TypeBool,
			Optional:    true,
			Computed:    true,
			Description: "Indicates whether or not the connection is domain level",
		},
		"enabled_clients": {
			Type:        schema.TypeSet,
			Elem:        &schema.Schema{Type: schema.TypeString},
			Optional:    true,
			Computed:    true,
			Description: "IDs of the clients for which the connection is enabled",
		},
		"realms": {
			Type:        schema.TypeList,
			Elem:        &schema.Schema{Type: schema.TypeString},
			Optional:    true,
			Computed:    true,
			Description: "Defines the realms for which the connection will be used (i.e., email domains). If not specified, the connection name is added as the realm",
		},
		"options": optionsSchema,
	}
}
