package random

import "testing"

func TestString(t *testing.T) {
	s := String(6)
	if len(s) != 6 {
		t.Errorf("unexpected string length")
	}
}

func TestTemplate(t *testing.T) {
	if s := Template(`{{.random}}`, "foo"); s != "foo" {
		t.Errorf("unexpected result from template")
	}
}
