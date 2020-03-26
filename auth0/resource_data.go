package auth0

import (
	"reflect"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/structure"

	"gopkg.in/auth0.v4"
)

// Data generalises schema.ResourceData so that we can reuse the accessor
// methods defined below.
type Data interface {

	// IsNewResource reports whether or not the resource is seen for the first
	// time. If so, checks for change won't be carried out.
	IsNewResource() bool

	// HasChange reports whether or not the given key has been changed.
	HasChange(key string) bool

	// HasChanges reports whether or not any of the given keys have been changed.
	HasChanges(keys ...string) bool

	// GetChange returns the old and new value for a given key.
	GetChange(key string) (interface{}, interface{})

	// Get returns the data for the given key, or nil if the key doesn't exist
	// in the schema.
	Get(key string) interface{}

	// GetOkExists returns the data for a given key and whether or not the key
	// has been set to a non-zero value. This is only useful for determining
	// if boolean attributes have been set, if they are Optional but do not
	// have a Default value.
	GetOkExists(key string) (interface{}, bool)
}

type data struct {
	prefix string
	Data
}

func dataAtKey(key string, d Data) Data { return &data{key, d} }
func dataAtIndex(i int, d Data) Data    { return &data{strconv.Itoa(i), d} }

func (d *data) IsNewResource() bool {
	return d.Data.IsNewResource()
}

func (d *data) HasChange(key string) bool {
	return d.Data.HasChange(d.prefix + "." + key)
}

func (d *data) HasChanges(keys ...string) bool {
	for _, key := range keys {
		if d.HasChange(key) {
			return true
		}
	}
	return false
}

func (d *data) GetChange(key string) (interface{}, interface{}) {
	return d.Data.GetChange(d.prefix + "." + key)
}

func (d *data) Get(key string) interface{} {
	return d.Data.Get(d.prefix + "." + key)
}

func (d *data) GetOkExists(key string) (interface{}, bool) {
	return d.Data.GetOkExists(d.prefix + "." + key)
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

func (md MapData) GetChange(key string) (interface{}, interface{}) {
	return md[key], md[key]
}

func (md MapData) Get(key string) interface{} {
	return md[key]
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
		s = StringIfExists(d, key)
	}
	return
}

// Returns a pointer to the value of the given key cast to a string or nil (ignores if changes exist or not)
func StringIfExists(d Data, key string) (s *string) {
	v, ok := d.GetOkExists(key)
	if ok {
		s = auth0.String(v.(string))
	}
	return
}

// Int accesses the value held by key and type asserts it to a pointer to a
// int.
func Int(d Data, key string) (i *int) {
	if d.IsNewResource() || d.HasChange(key) {
		i = IntIfExists(d, key)
	}
	return
}

// IntIfExists a pointer to the value of the given key cast to an integer or nil (ignores if changes exist or not)
func IntIfExists(d Data, key string) (i *int) {
	v, ok := d.GetOkExists(key)
	if ok {
		i = auth0.Int(v.(int))
	}
	return
}

// Bool accesses the value held by key and type asserts it to a pointer to a
// bool.
func Bool(d Data, key string) (b *bool) {
	if d.IsNewResource() || d.HasChange(key) {
		b = BoolIfExists(d, key)
	}
	return
}

// BoolIfExists returns a pointer to the value of the key cast to bool (ignores if changes exist or not)
func BoolIfExists(d Data, key string) (b *bool) {
	v, ok := d.GetOkExists(key)
	if ok {
		b = auth0.Bool(v.(bool))
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

// SliceIfExists returns a pointer to the value of the key cast to []interface{} (ignores if changes exist or not)
func SliceIfExists(d Data, key string) (s []interface{}) {
	v, ok := d.GetOkExists(key)
	if ok {
		s = v.([]interface{})
	}
	return
}

// Map accesses the value held by key and type asserts it to a map.
func Map(d Data, key string) (m map[string]interface{}) {
	if d.IsNewResource() || d.HasChange(key) {
		m = MapIfExists(d, key)
	}
	return
}

// Returns a pointer to the value of the key cast to map[string]interface{} (ignores if changes exist or not)
func MapIfExists(d Data, key string) (m map[string]interface{}) {
	v, ok := d.GetOkExists(key)
	if ok {
		m = v.(map[string]interface{})
	}
	return
}

// Map accesses the value held by key and fills a new map[string]string with the contents if the object is new or has changed.
func StringMap(d Data, key string) (m map[string]string) {
	values := Map(d, key)
	if values != nil {
		m = make(map[string]string)
		for key, value := range values {
			m[key] = value.(string)
		}
	}
	return
}

// Map accesses the value held by key and fills a new map[string]string with the contents. (ignores if changes exist or not)
func StringMapIfExists(d Data, key string) (m map[string]string) {
	values := MapIfExists(d, key)
	if values != nil {
		m = make(map[string]string)
		for key, value := range values {
			m[key] = value.(string)
		}
	}
	return
}

// List accesses the value held by key and returns an iterator able to go over
// its elements only if the resource is new or the key has changed.
// Returns an empty list if the resource is not new and the key is unchanged
func List(d Data, key string) Iterator {
	if d.IsNewResource() || d.HasChange(key) {
		v, ok := d.GetOkExists(key)
		if ok {
			return &list{dataAtKey(key, d), v.([]interface{})}
		}
	}
	return &list{}
}

// List accesses the value held by key and returns an iterator able to go over
// its elements. (ignores if the resource is new or the key has changed).
// Ignores if the resource is new or the key has changed.
func ListIfExists(d Data, key string) Iterator {
	v, ok := d.GetOkExists(key)
	if ok {
		return &list{dataAtKey(key, d), v.([]interface{})}
	}
	return &list{}
}

// Set accesses the value held by key, type asserts it to a set and returns an
// iterator able to go over its elements. Returns empty iterator if the resource is not new and the key is unchanged
func Set(d Data, key string) Iterator {
	if d.IsNewResource() || d.HasChange(key) {
		return SetIfExists(d, key)
	}
	return &set{nil, &schema.Set{}}
}

// Set accesses the value held by key, type asserts it to a set and returns an
// iterator able to go over its elements. Ignores if the resource is new or the key has changed.
func SetIfExists(d Data, key string) Iterator {
	v, ok := d.GetOkExists(key)
	if ok {
		if s, ok := v.(*schema.Set); ok {
			return &set{dataAtKey(key, d), s}
		}
	}
	return &set{nil, &schema.Set{}}
}

type Iterator interface {

	// Elem iterates over all elements of the list or set, calling fn with each
	// iteration.
	//
	// The callback takes a Data interface as argument which is prefixed with
	// its parents key, making nested data access more convenient.
	//
	// The operation
	//
	// 	bar = d.Get("foo.0.bar").(string)
	//
	// can be expressed as
	//
	// 	List(d, "foo").Elem(func (d Data) {
	//		bar = String(d, "bar")
	// 	})
	//
	// making data access more intuitive for nested structures.
	Elem(func(d Data))

	// Range iterates over all elements of the list, calling fn in each iteration.
	Range(func(k int, v interface{}))

	// List returns the underlying list as a Go slice.
	List() []interface{}
}

type list struct {
	d Data
	v []interface{}
}

func (l *list) Range(fn func(key int, value interface{})) {
	for key, value := range l.v {
		fn(key, value)
	}
}

func (l *list) Elem(fn func(Data)) {
	for idx := range l.v {
		fn(dataAtIndex(idx, l.d))
	}
}

func (l *list) List() []interface{} {
	return l.v
}

type set struct {
	d Data
	s *schema.Set
}

func (s *set) hash(item interface{}) string {
	code := s.s.F(item)
	if code < 0 {
		code = -code
	}
	return strconv.Itoa(code)
}

func (s *set) Range(fn func(key int, value interface{})) {
	for key, value := range s.s.List() {
		fn(key, value)
	}
}

func (s *set) Elem(fn func(Data)) {
	for _, v := range s.s.List() {
		fn(dataAtKey(s.hash(v), s.d))
	}
}

func (s *set) List() []interface{} {
	return s.s.List()
}

// Diff accesses the value held by key and type asserts it to a set. It then
// compares it's changes if any and returns what needs to be added and what
// needs to be removed.
func Diff(d Data, key string) (add []interface{}, rm []interface{}) {
	if d.IsNewResource() {
		add = Set(d, key).List()
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
