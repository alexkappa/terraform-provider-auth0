package auth0

import (
	"github.com/hashicorp/terraform/helper/schema"
	auth0 "github.com/yieldr/go-auth0"
)

func String(d *schema.ResourceData, key string) (s *string) {
	v, ok := d.GetOk(key)
	if ok {
		s = auth0.String(v.(string))
	}
	return
}

func MapString(m map[string]interface{}, key string) (s *string) {
	v, ok := m[key]
	if ok && v != "" {
		s = auth0.String(v.(string))
	}
	return
}

func Int(d *schema.ResourceData, key string) (i *int) {
	v, ok := d.GetOk(key)
	if ok {
		i = auth0.Int(v.(int))
	}
	return
}

func MapInt(m map[string]interface{}, key string) (i *int) {
	v, ok := m[key]
	if ok && v != 0 {
		i = auth0.Int(v.(int))
	}
	return
}

func Bool(d *schema.ResourceData, key string) (b *bool) {
	v, ok := d.GetOk(key)
	if ok {
		b = auth0.Bool(v.(bool))
	}
	return
}

func MapBool(m map[string]interface{}, key string) (b *bool) {
	v, ok := m[key]
	if ok && v != false {
		b = auth0.Bool(v.(bool))
	}
	return
}

func Slice(d *schema.ResourceData, key string) (s []interface{}) {
	v, ok := d.GetOk(key)
	if ok {
		s = v.([]interface{})
	}
	return
}

func Map(d *schema.ResourceData, key string) (m map[string]interface{}) {
	v, ok := d.GetOk(key)
	if ok {
		m = v.(map[string]interface{})
	}
	return
}
