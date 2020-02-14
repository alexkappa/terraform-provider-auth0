package random

import (
	"bytes"
	"text/template"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

// String generates a random alphanumeric string of the length specified.
func String(strlen int) string {
	return acctest.RandString(strlen)
}

// Template renders templates defined with {{.random}} placeholders. This is
// useful for acceptance tests. By introducing entropy to resources generated
// during the tests, we can run tests concurrently.
func Template(tpl, rand string) string {
	var buf bytes.Buffer
	t := template.Must(template.New("tpl").Parse(tpl))
	err := t.Execute(&buf, map[string]string{"random": rand})
	if err != nil {
		panic(err)
	}
	return buf.String()
}

// TestCheckResourceAttr is a TestCheckFunc which validates the value in state
// for the given name/key combination after applying a template over the value.
func TestCheckResourceAttr(name, key, value, rand string) resource.TestCheckFunc {
	return resource.TestCheckResourceAttr(name, key, Template(value, rand))
}
