package hash

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/hashcode"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// StringKey returns a schema.SchemaSetFunc able to hash a string value
// from map accessed by k.
func StringKey(k string) schema.SchemaSetFunc {
	return func(v interface{}) int {
		m, ok := v.(map[string]interface{})
		if !ok {
			return 0
		}
		if v, ok := m[k].(string); ok {
			return hashcode.String(v)
		}
		return 0
	}
}
