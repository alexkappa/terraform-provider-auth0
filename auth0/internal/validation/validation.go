package validation

import (
	"fmt"
	"net/url"
)

// IsURLWithNoFragment is a SchemaValidateFunc which tests if the provided value
// is of type string and a valid URL with no fragment.
func IsURLWithNoFragment(i interface{}, k string) (warnings []string, errors []error) {

	v, ok := i.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected type of %q to be string", k))
		return
	}

	if v == "" {
		errors = append(errors, fmt.Errorf("expected %q url to not be empty, got %v", k, i))
		return
	}

	u, err := url.Parse(v)
	if err != nil {
		errors = append(errors, fmt.Errorf("expected %q to be a valid url, got %v: %+v", k, v, err))
		return
	}

	if u.Host == "" {
		errors = append(errors, fmt.Errorf("expected %q to have a host, got %v", k, v))
		return
	}

	if u.Fragment != "" {
		errors = append(errors, fmt.Errorf("expected %q to have a url with an empty fragment. %s", k, v))
	}

	return
}
