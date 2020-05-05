package auth0

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"gopkg.in/auth0.v4/management"
)

func newSMSConnection() *schema.Resource {
	return &schema.Resource{
		Create: createConnectionWithStrategy(management.ConnectionStrategySMS),
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
					"name": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "",
					},
					"twilio_sid": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "",
					},
					"twilio_token": {
						Type:        schema.TypeString,
						Optional:    true,
						Sensitive:   true,
						DefaultFunc: schema.EnvDefaultFunc("TWILIO_TOKEN", nil),
						Description: "",
					},
					"from": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "",
					},
					"syntax": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "",
					},
					"template": {
						Type:        schema.TypeString,
						Optional:    true,
						Description: "",
					},
					"messaging_service_sid": {
						Type:     schema.TypeString,
						Optional: true,
					},
					"disable_signup": {
						Type:        schema.TypeBool,
						Optional:    true,
						Description: "Indicates whether or not to allow user sign-ups to your application",
					},
					"brute_force_protection": {
						Type:        schema.TypeBool,
						Optional:    true,
						Description: "Indicates whether or not to enable brute force protection, which will limit the number of signups and failed logins from a suspicious IP address",
					},
					"totp": {
						Type:     schema.TypeList,
						Optional: true,
						MaxItems: 1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"time_step": {
									Type:     schema.TypeInt,
									Optional: true,
								},
								"length": {
									Type:     schema.TypeInt,
									Optional: true,
								},
							},
						},
						Description: "",
					},
				},
			},
			Description: "Configuration settings for connection options",
		}),
	}
}
