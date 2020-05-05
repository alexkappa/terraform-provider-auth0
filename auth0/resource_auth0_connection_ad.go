package auth0

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"gopkg.in/auth0.v4/management"
)

func newADConnection() *schema.Resource {
	return &schema.Resource{
		Create: createConnectionWithStrategy(management.ConnectionStrategyAD),
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
					"ips": {
						Type:        schema.TypeSet,
						Elem:        &schema.Schema{Type: schema.TypeString},
						Optional:    true,
						Description: "",
					},
					"use_cert_auth": {
						Type:        schema.TypeBool,
						Optional:    true,
						Description: "",
					},
					"use_kerberos": {
						Type:        schema.TypeBool,
						Optional:    true,
						Description: "",
					},
					"disable_cache": {
						Type:        schema.TypeBool,
						Optional:    true,
						Description: "",
					},
					"brute_force_protection": {
						Type:        schema.TypeBool,
						Optional:    true,
						Description: "Indicates whether or not to enable brute force protection, which will limit the number of signups and failed logins from a suspicious IP address",
					},
				},
			},
			Description: "Configuration settings for connection options",
		}),
	}
}
