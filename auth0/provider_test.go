package auth0

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"

	"gopkg.in/auth0.v5/management"
)

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
