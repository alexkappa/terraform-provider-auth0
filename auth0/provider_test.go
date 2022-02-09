package auth0

import (
	"errors"
	"os"
	"sort"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"gopkg.in/auth0.v5/management"
)

const wiremockHost = "localhost:8080"

func providerWithTestingConfiguration() *schema.Provider {
	provider := Provider()
	provider.ConfigureFunc = func(data *schema.ResourceData) (interface{}, error) {
		return management.New(
			wiremockHost,
			management.WithInsecure(),
			management.WithDebug(true),
		)
	}
	return provider
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

func TestProvider_configValidation(t *testing.T) {
	testCases := []struct {
		name           string
		environment    map[string]string
		resourceConfig map[string]interface{}
		expectedErrors []error
	}{
		{
			name:           "missing client id",
			environment:    map[string]string{"AUTH0_DOMAIN": "test", "AUTH0_CLIENT_SECRET": "test"},
			expectedErrors: []error{errors.New("\"client_secret\": all of `client_id,client_secret` must be specified")},
		},
		{
			name:           "missing client secret",
			environment:    map[string]string{"AUTH0_DOMAIN": "test", "AUTH0_CLIENT_ID": "test"},
			expectedErrors: []error{errors.New("\"client_id\": all of `client_id,client_secret` must be specified")},
		},
		{
			name:           "conflicting auth0 client and management token without domain",
			resourceConfig: map[string]interface{}{"client_id": "test", "client_secret": "test", "api_token": "test"},
			environment:    map[string]string{},
			expectedErrors: []error{
				errors.New("\"domain\": required field is not set"),
				errors.New("\"client_id\": conflicts with api_token"),
				errors.New("\"client_secret\": conflicts with api_token"),
				errors.New("\"api_token\": conflicts with client_id"),
			},
		},
		{
			name:           "valid auth0 client",
			resourceConfig: map[string]interface{}{"domain": "valid_domain", "client_id": "test", "client_secret": "test"},
			environment:    map[string]string{},
			expectedErrors: nil,
		},
		{
			name:           "valid auth0 token",
			resourceConfig: map[string]interface{}{"domain": "valid_domain", "api_token": "test"},
			environment:    map[string]string{},
			expectedErrors: nil,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			for k, v := range test.environment {
				os.Unsetenv(k)
				os.Setenv(k, v)
			}

			c := terraform.NewResourceConfigRaw(test.resourceConfig)
			p := Provider()

			_, errs := p.Validate(c)
			assertErrorsSliceEqual(t, test.expectedErrors, errs)

			for k := range test.environment {
				os.Unsetenv(k)
			}
		})
	}
}

func sortErrors(errs []error) {
	sort.Slice(errs, func(i, j int) bool {
		return errs[i].Error() < errs[j].Error()
	})
}

func assertErrorsSliceEqual(t *testing.T, expected, actual []error) {
	if len(expected) != len(actual) {
		t.Fatalf("actual did not match expected. len(expected) != len(actual). expected: %v, actual: %v", expected, actual)
	}

	sortErrors(expected)
	sortErrors(actual)

	for i := range expected {
		if expected[i].Error() != actual[i].Error() {
			t.Fatalf("actual did not match expected. expected[%d] != actual[%d]. expected: %v, actual: %v", i, i, expected, actual)
		}
	}
}
