package auth0

import (
	"fmt"
	"os"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/meta"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/terraform-providers/terraform-provider-auth0/version"

	"gopkg.in/auth0.v4"
	"gopkg.in/auth0.v4/management"
)

// Provider returns a terraform.ResourceProvider.
func Provider() terraform.ResourceProvider {
	provider := &schema.Provider{
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
			"auth0_client":                          newClient(),
			"auth0_global_client":                   newGlobalClient(),
			"auth0_client_grant":                    newClientGrant(),
			"auth0_connection":                      newConnection(),
			"auth0_connection_auth0":                newAuth0Connection(),
			"auth0_connection_ad":                   newADConnection(),
			"auth0_connection_adfs":                 newADFSConnection(),
			"auth0_connection_waad":                 newAzureADConnection(),
			"auth0_connection_google_oauth2":        newGoogleOAuth2Connection(),
			"auth0_connection_github":               newGitHubConnection(),
			"auth0_connection_email":                newEmailConnection(),
			"auth0_connection_sms":                  newSMSConnection(),
			"auth0_connection_salesforce":           newSalesforceConnection(management.ConnectionStrategySalesforce),
			"auth0_connection_salesforce_sandbox":   newSalesforceConnection(management.ConnectionStrategySalesforceSandbox),
			"auth0_connection_salesforce_community": newSalesforceConnection(management.ConnectionStrategySalesforceCommunity),
			"auth0_custom_domain":                   newCustomDomain(),
			"auth0_resource_server":                 newResourceServer(),
			"auth0_rule":                            newRule(),
			"auth0_rule_config":                     newRuleConfig(),
			"auth0_hook":                            newHook(),
			"auth0_prompt":                          newPrompt(),
			"auth0_email":                           newEmail(),
			"auth0_email_template":                  newEmailTemplate(),
			"auth0_user":                            newUser(),
			"auth0_tenant":                          newTenant(),
			"auth0_role":                            newRole(),
		},
	}

	provider.ConfigureFunc = func(data *schema.ResourceData) (interface{}, error) {
		domain := data.Get("domain").(string)
		id := data.Get("client_id").(string)
		secret := data.Get("client_secret").(string)
		debug := data.Get("debug").(bool)

		userAgent := fmt.Sprintf("Terraform-Provider-Auth0/%s (Go-Auth0-SDK/%s; Terraform-SDK/%s; Terraform/%s)",
			version.ProviderVersion,
			auth0.Version,
			meta.SDKVersionString(),
			provider.TerraformVersion)

		return management.New(domain, id, secret,
			management.WithDebug(debug),
			management.WithUserAgent(userAgent))
	}

	return provider
}
