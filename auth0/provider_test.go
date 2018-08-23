package auth0

import (
	"os"
	"testing"
)

func TestProvider(t *testing.T) {
	if err := Provider().InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_debugDefaults(t *testing.T) {
	testCases := map[string]struct {
		envSet   bool
		envValue string
		expected bool
	}{
		`should be true when AUTH0_DEBUG environment variable is set to "1"`: {
			envSet:   true,
			envValue: "1",
			expected: true,
		},
		`should be true when AUTH0_DEBUG environment variable is set to "true"`: {
			envSet:   true,
			envValue: "true",
			expected: true,
		},
		`should be true when AUTH0_DEBUG environment variable is set to "on"`: {
			envSet:   true,
			envValue: "on",
			expected: true,
		},
		`should be false when AUTH0_DEBUG environment variable is set to "false"`: {
			envSet:   true,
			envValue: "false",
			expected: false,
		},
		"should be false if no AUTH0_DEBUG environment variable is defined": {
			envSet:   false,
			expected: false,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			os.Unsetenv("AUTH0_DEBUG")
			if tc.envSet {
				os.Setenv("AUTH0_DEBUG", tc.envValue)
			}

			p := Provider()

			debug, err := p.Schema["debug"].DefaultValue()
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			if debug.(bool) != tc.expected {
				t.Fatalf("Expected debug to be %v, but got %v", tc.expected, debug)
			}
		})
	}
}
