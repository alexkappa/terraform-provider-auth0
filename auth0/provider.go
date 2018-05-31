package auth0

import (
	"github.com/90poe/go-auth0/management"
	"github.com/hashicorp/terraform/helper/schema"
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
		},
		ResourcesMap: map[string]*schema.Resource{
			"auth0_client":          newClient(),
			"auth0_client_grant":    newClientGrant(),
			"auth0_connection":      newConnection(),
			"auth0_custom_domain":   newCustomDomain(),
			"auth0_resource_server": newResourceServer(),
			"auth0_rule":            newRule(),
			"auth0_rule_config":     newRuleConfig(),
		},
		ConfigureFunc: configure,
	}
}

func configure(data *schema.ResourceData) (interface{}, error) {

	domain := data.Get("domain").(string)
	id := data.Get("client_id").(string)
	secret := data.Get("client_secret").(string)

	return management.New(domain, id, secret)
}
