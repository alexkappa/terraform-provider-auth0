package auth0

import (
	"os"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/yieldr/go-auth0/management"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"domain": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("AUTH0_DOMAIN", nil),
			},
			"client_id": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("AUTH0_CLIENT_ID", nil),
			},
			"client_secret": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("AUTH0_CLIENT_SECRET", nil),
			},
			"debug": {
				Type:     schema.TypeBool,
				Optional: true,
				DefaultFunc: func() (interface{}, error) {
					v := os.Getenv("AUTH0_DEBUG")
					if v == "" {
						return false, nil
					}

					return v == "1" || v == "true" || v == "on", nil
				},
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"auth0_client":          newClient(),
			"auth0_client_grant":    newClientGrant(),
			"auth0_connection":      newConnection(),
			"auth0_custom_domain":   newCustomDomain(),
			"auth0_resource_server": newResourceServer(),
			"auth0_rule":            newRule(),
			"auth0_rule_config":     newRuleConfig(),
			"auth0_email":           newEmail(),
			"auth0_email_template":  newEmailTemplate(),
			"auth0_user":            newUser(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"auth0_client": newClientDataSource(),
		},

		ConfigureFunc: configure,
	}
}

func configure(data *schema.ResourceData) (interface{}, error) {

	domain := data.Get("domain").(string)
	id := data.Get("client_id").(string)
	secret := data.Get("client_secret").(string)
	debug := data.Get("debug").(bool)

	return management.New(domain, id, secret, management.WithDebug(debug))
}
