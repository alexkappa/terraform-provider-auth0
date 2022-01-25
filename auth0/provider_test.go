package auth0

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"gopkg.in/auth0.v5/management"
)

func providerWithWiremock() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"auth0_client":                     newClient(),
			"auth0_global_client":              newGlobalClient(),
			"auth0_client_grant":               newClientGrant(),
			"auth0_connection":                 newConnection(),
			"auth0_custom_domain":              newCustomDomain(),
			"auth0_custom_domain_verification": newCustomDomainVerification(),
			"auth0_resource_server":            newResourceServer(),
			"auth0_rule":                       newRule(),
			"auth0_rule_config":                newRuleConfig(),
			"auth0_hook":                       newHook(),
			"auth0_prompt":                     newPrompt(),
			"auth0_prompt_custom_text":         newPromptCustomText(),
			"auth0_email":                      newEmail(),
			"auth0_email_template":             newEmailTemplate(),
			"auth0_user":                       newUser(),
			"auth0_tenant":                     newTenant(),
			"auth0_role":                       newRole(),
			"auth0_log_stream":                 newLogStream(),
			"auth0_branding":                   newBranding(),
			"auth0_guardian":                   newGuardian(),
			"auth0_organization":               newOrganization(),
			"auth0_action":                     newAction(),
			"auth0_trigger_binding":            newTriggerBinding(),
		},
		ConfigureFunc: func(data *schema.ResourceData) (interface{}, error) {
			return management.New(
				"localhost:8080",
				management.WithInsecure(),
				management.WithDebug(true),
			)
		},
	}
}

func Auth0() (*management.Management, error) {
	c := terraform.NewResourceConfigRaw(nil)
	p := Provider()
	if err := p.Configure(c); err != nil {
		return nil, err
	}
	return p.Meta().(*management.Management), nil
}

func TestMain(m *testing.M) {
	resource.TestMain(m)
}

func TestProvider(t *testing.T) {
	if err := Provider().InternalValidate(); err != nil {
		t.Fatal(err)
	}
}

func TestProvider_debugDefaults(t *testing.T) {
	for value, expected := range map[string]bool{
		"1":     true,
		"true":  true,
		"on":    true,
		"0":     false,
		"off":   false,
		"false": false,
		"foo":   false,
		"":      false,
	} {
		os.Unsetenv("AUTH0_DEBUG")
		if value != "" {
			os.Setenv("AUTH0_DEBUG", value)
		}

		p := Provider()
		debug, err := p.Schema["debug"].DefaultValue()
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		if debug.(bool) != expected {
			t.Fatalf("Expected debug to be %v, but got %v", expected, debug)
		}
	}
}
