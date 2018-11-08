package auth0

import (
	auth0 "github.com/yieldr/go-auth0"
)

// Data generalises schema.ResourceData so that we can reuse the accessor
// methods we define below.
type Data interface {

	// HasChange reports whether or not the given key has been changed.
	HasChange(key string) bool

	// GetOkExists returns the data for a given key and whether or not the key
	// has been set to a non-zero value. This is only useful for determining
	// if boolean attributes have been set, if they are Optional but do not
	// have a Default value.
	GetOkExists(key string) (interface{}, bool)
}

func MapData(m map[string]interface{}) Data {
	return mapData(m)
}

type mapData map[string]interface{}

func (md mapData) HasChange(key string) bool {
	_, ok := md[key]
	return ok
}

func (md mapData) GetOkExists(key string) (interface{}, bool) {
	v, ok := md[key]
	return v, ok
}

func String(d Data, key string) (s *string) {
	if d.HasChange(key) {
		v, ok := d.GetOkExists(key)
		if ok {
			s = auth0.String(v.(string))
		}
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

func Int(d Data, key string) (i *int) {
	if d.HasChange(key) {
		v, ok := d.GetOkExists(key)
		if ok {
			i = auth0.Int(v.(int))
		}
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

func Bool(d Data, key string) (b *bool) {
	if d.HasChange(key) {
		v, ok := d.GetOkExists(key)
		if ok {
			b = auth0.Bool(v.(bool))
		}
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

func Slice(d Data, key string) (s []interface{}) {
	if d.HasChange(key) {
		v, ok := d.GetOkExists(key)
		if ok {
			s = v.([]interface{})
		}
	}
	return
}

func Map(d Data, key string) (m map[string]interface{}) {
	if d.HasChange(key) {
		v, ok := d.GetOkExists(key)
		if ok {
			m = v.(map[string]interface{})
		}
	}
	return
}

type Iterator struct {
	i []interface{}
}

func (i *Iterator) All(f func(key int, value interface{})) {
	for key, value := range i.i {
		f(key, value)
	}
}

func (i *Iterator) First(f func(value interface{})) {
	for _, value := range i.i {
		f(value)
		return
	}
}

func List(d Data, key string) *Iterator {
	if d.HasChange(key) {
		v, ok := d.GetOkExists(key)
		if ok {
			return &Iterator{v.([]interface{})}
		}
	}
	return &Iterator{}
}
