package validation

import "testing"

func TestIsURLWithNoFragment(t *testing.T) {
	for url, valid := range map[string]bool{
		"http://example.com":      true,
		"http://example.com/foo":  true,
		"http://example.com#foo":  false,
		"https://example.com/foo": true,
		"https://example.com#foo": false,
	} {
		_, err := IsURLWithNoFragment(url, "url")
		if err != nil && valid {
			t.Errorf("IsURLWithNoFragment(%s) produced an unexpected error", url)
		}
	}
}
