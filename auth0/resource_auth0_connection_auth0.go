package auth0

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"gopkg.in/auth0.v4/management"
)

func newAuth0Connection() *schema.Resource {
	return &schema.Resource{
		Create: createConnectionWithStrategy(management.ConnectionStrategyAuth0),
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
					"validation": {
						Type:     schema.TypeMap,
						Elem:     &schema.Schema{Type: schema.TypeString},
						Optional: true,
					},
					"password_policy": {
						Type:     schema.TypeString,
						Optional: true,
						Computed: true,
						ValidateFunc: validation.StringInSlice([]string{
							"none", "low", "fair", "good", "excellent",
						}, false),
						Description: "Indicates level of password strength to enforce during authentication. A strong password policy will make it difficult, if not improbable, for someone to guess a password through either manual or automated means. Options include `none`, `low`, `fair`, `good`, `excellent`",
					},
					"password_history": {
						Type:     schema.TypeList,
						Optional: true,
						Computed: true,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"enable": {
									Type:     schema.TypeBool,
									Optional: true,
								},
								"size": {
									Type:     schema.TypeInt,
									Optional: true,
								},
							},
						},
						Description: "Configuration settings for the password history that is maintained for each user to prevent the reuse of passwords",
					},
					"password_no_personal_info": {
						Type:     schema.TypeList,
						Optional: true,
						MaxItems: 1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"enable": {
									Type:     schema.TypeBool,
									Optional: true,
								},
							},
						},
						Description: "Configuration settings for the password personal info check, which does not allow passwords that contain any part of the user's personal data, including user's name, username, nickname, user_metadata.name, user_metadata.first, user_metadata.last, user's email, or firstpart of the user's email",
					},
					"password_dictionary": {
						Type:     schema.TypeList,
						Optional: true,
						MaxItems: 1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"enable": {
									Type:     schema.TypeBool,
									Optional: true,
								},
								"dictionary": {
									Type:     schema.TypeSet,
									Elem:     &schema.Schema{Type: schema.TypeString},
									Optional: true,
								},
							},
						},
						Description: "Configuration settings for the password dictionary check, which does not allow passwords that are part of the password dictionary",
					},
					"password_complexity_options": {
						Type:     schema.TypeList,
						Optional: true,
						MaxItems: 1,
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"min_length": {
									Type:         schema.TypeInt,
									Optional:     true,
									ValidateFunc: validation.IntAtLeast(1),
								},
							},
						},
						Description: "Configuration settings for password complexity",
					},
					"enabled_database_customization": {
						Type:        schema.TypeBool,
						Optional:    true,
						Description: "",
					},
					"brute_force_protection": {
						Type:        schema.TypeBool,
						Optional:    true,
						Description: "Indicates whether or not to enable brute force protection, which will limit the number of signups and failed logins from a suspicious IP address",
					},
					"import_mode": {
						Type:        schema.TypeBool,
						Optional:    true,
						Description: "Indicates whether or not you have a legacy user store and want to gradually migrate those users to the Auth0 user store",
					},
					"disable_signup": {
						Type:        schema.TypeBool,
						Optional:    true,
						Description: "Indicates whether or not to allow user sign-ups to your application",
					},
					"requires_username": {
						Type:        schema.TypeBool,
						Optional:    true,
						Description: "Indicates whether or not the user is required to provide a username in addition to an email address",
					},
					"custom_scripts": {
						Type:        schema.TypeMap,
						Elem:        &schema.Schema{Type: schema.TypeString},
						Optional:    true,
						Description: "",
					},
					"configuration": {
						Type:        schema.TypeMap,
						Elem:        &schema.Schema{Type: schema.TypeString},
						Sensitive:   true,
						Optional:    true,
						Description: "",
					},
				},
			},
			Description: "Configuration settings for connection options",
		}),
	}
}
