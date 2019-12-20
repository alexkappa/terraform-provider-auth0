package auth0

import (
	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/structure"
	"gopkg.in/auth0.v2"
)

// Data generalises schema.ResourceData so that we can reuse the accessor
// methods defined below.
type Data interface {

	// IsNewResource reports whether or not the resource is seen for the first
	// time. If so, checks for change won't be carried out.
	IsNewResource() bool

	// HasChange reports whether or not the given key has been changed.
	HasChange(key string) bool

	// GetOkExists returns the data for a given key and whether or not the key
	// has been set to a non-zero value. This is only useful for determining
	// if boolean attributes have been set, if they are Optional but do not
	// have a Default value.
	GetOkExists(key string) (interface{}, bool)
}

// MapData wraps a map satisfying the Data interface, so it can be used in the
// accessor methods defined below.
type MapData map[string]interface{}

func (md MapData) IsNewResource() bool {
	return false
}

func (md MapData) HasChange(key string) bool {
	_, ok := md[key]
	return ok
}

func (md MapData) GetOkExists(key string) (interface{}, bool) {
	v, ok := md[key]
	return v, ok && !isNil(v) && !isZero(v)
}

func isNil(v interface{}) bool {
	return v == nil
}

func isZero(v interface{}) bool {
	return reflect.DeepEqual(v, reflect.Zero(reflect.TypeOf(v)).Interface())
}

var _ Data = (*schema.ResourceData)(nil)

// String accesses the value held by key and type asserts it to a pointer to a
// string.
func String(d Data, key string) (s *string) {
	if d.IsNewResource() || d.HasChange(key) {
		v, ok := d.GetOkExists(key)
		if ok {
			s = auth0.String(v.(string))
		}
	}
	return
}

// Int accesses the value held by key and type asserts it to a pointer to a
// int.
func Int(d Data, key string) (i *int) {
	if d.IsNewResource() || d.HasChange(key) {
		v, ok := d.GetOkExists(key)
		if ok {
			i = auth0.Int(v.(int))
		}
	}
	return
}

// Bool accesses the value held by key and type asserts it to a pointer to a
// bool.
func Bool(d Data, key string) (b *bool) {
	if d.IsNewResource() || d.HasChange(key) {
		v, ok := d.GetOkExists(key)
		if ok {
			b = auth0.Bool(v.(bool))
		}
	}
	return
}

// Slice accesses the value held by key and type asserts it to a slice.
func Slice(d Data, key string) (s []interface{}) {
	if d.IsNewResource() || d.HasChange(key) {
		v, ok := d.GetOkExists(key)
		if ok {
			s = v.([]interface{})
		}
	}
	return
}

// Map accesses the value held by key and type asserts it to a map.
func Map(d Data, key string) (m map[string]interface{}) {
	if d.IsNewResource() || d.HasChange(key) {
		v, ok := d.GetOkExists(key)
		if ok {
			m = v.(map[string]interface{})
		}
	}
	return
}

// List accesses the value held by key and returns an iterator able to go over
// the items of the list.
//
// The iterator can go over all the items in the list or just the first one,
// which is a common use case for defining nested schemas in Terraform.
func List(d Data, key string) *iterator {
	if d.IsNewResource() || d.HasChange(key) {
		v, ok := d.GetOkExists(key)
		if ok {
			return &iterator{v.([]interface{})}
		}
	}
	return &iterator{}
}

// Set accesses the value held by key, type asserts it to a set and returns an
// iterator able to go over the items of the list.
func Set(d Data, key string) *iterator {
	if d.IsNewResource() || d.HasChange(key) {
		v, ok := d.GetOkExists(key)
		if ok {
			if s, ok := v.(*schema.Set); ok {
				return &iterator{s.List()}
			}
		}
	}
	return &iterator{}
}

type iterator struct {
	i []interface{}
}

// All iterates over all elements of the list, calling f in each iteration.
func (i *iterator) All(f func(key int, value interface{})) {
	for key, value := range i.i {
		f(key, value)
	}
}

// First iterates over the first element of the list, calling f with the value
// at the first key.
//
// The function f will be called at most one time, as the list may be empty.
func (i *iterator) First(f func(value interface{})) {
	for _, value := range i.i {
		f(value)
		return
	}
}

// Slice returns the underlying list as a raw slice.
func (i *iterator) Slice() []interface{} {
	return i.i
}

// Set accesses the value held by key, type asserts it to a set. It then
// compares it's changes if any and returns what needs to be added (created) and
// what needs to be removed (delete).
func Diff(d *schema.ResourceData, key string) (add []interface{}, rm []interface{}) {
	if d.IsNewResource() {
		add = Set(d, key).Slice()
	}
	if d.HasChange(key) {
		o, n := d.GetChange(key)
		add = n.(*schema.Set).Difference(o.(*schema.Set)).List()
		rm = o.(*schema.Set).Difference(n.(*schema.Set)).List()
	}
	return
}

// JSON accesses the value held by key and unmarshals it into a map.
func JSON(d Data, key string) (m map[string]interface{}, err error) {
	if d.IsNewResource() || d.HasChange(key) {
		v, ok := d.GetOkExists(key)
		if ok {
			m, err = structure.ExpandJsonFromString(v.(string))
			if err != nil {
				return
			}
		}
	}
	return
}
