package auth0

import (
	"reflect"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/structure"

	"gopkg.in/auth0.v4"
)

// ResourceData generalises schema.ResourceData so that we can reuse the
// accessor methods defined below.
type ResourceData interface {

	// IsNewResource reports whether or not the resource is seen for the first
	// time. If so, checks for change won't be carried out.
	IsNewResource() bool

	// HasChange reports whether or not the given key has been changed.
	HasChange(key string) bool

	// GetChange returns the old and new value for a given key.
	GetChange(key string) (interface{}, interface{})

	// Get returns the data for the given key, or nil if the key doesn't exist
	// in the schema.
	Get(key string) interface{}

	// GetOk returns the data for the given key and whether or not the key
	// has been set to a non-zero value at some point.
	//
	// The first result will not necessarilly be nil if the value doesn't exist.
	// The second result should be checked to determine this information.
	GetOk(key string) (interface{}, bool)

	// GetOkExists can check if TypeBool attributes that are Optional with
	// no Default value have been set.
	//
	// Deprecated: usage is discouraged due to undefined behaviors and may be
	// removed in a future version of the SDK
	GetOkExists(key string) (interface{}, bool)
}

type resourceData struct {
	ResourceData
	prefix string
}

func newResourceDataAtKey(key string, d ResourceData) ResourceData {
	return &resourceData{d, key}
}

func newResourceDataAtIndex(i int, d ResourceData) ResourceData {
	return &resourceData{d, strconv.Itoa(i)}
}

func (d *resourceData) IsNewResource() bool {
	return d.ResourceData.IsNewResource()
}

func (d *resourceData) HasChange(key string) bool {
	return d.ResourceData.HasChange(d.prefix + "." + key)
}

func (d *resourceData) GetChange(key string) (interface{}, interface{}) {
	return d.ResourceData.GetChange(d.prefix + "." + key)
}

func (d *resourceData) Get(key string) interface{} {
	return d.ResourceData.Get(d.prefix + "." + key)
}

func (d *resourceData) GetOk(key string) (interface{}, bool) {
	return d.ResourceData.GetOk(d.prefix + "." + key)
}

func (d *resourceData) GetOkExists(key string) (interface{}, bool) {
	return d.ResourceData.GetOkExists(d.prefix + "." + key)
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

func (md MapData) GetOk(key string) (interface{}, bool) {
	v, ok := md[key]
	return v, ok && !isNil(v) && !isZero(v)
}

func (md MapData) GetOkExists(key string) (interface{}, bool) {
	v, ok := md[key]
	return v, ok && !isNil(v)
}

func isNil(v interface{}) bool {
	return v == nil
}

func isZero(v interface{}) bool {
	return reflect.DeepEqual(v, reflect.Zero(reflect.TypeOf(v)).Interface())
}

var _ ResourceData = (*schema.ResourceData)(nil)

// Condition is a function that checks whether a condition holds true for a
// value being accessed.
//
// It is used with accessor functions such as Int, String, etc to only retrieve
// the value if the conditions hold true.
type Condition func(d ResourceData, key string) bool

// Eval performs the evaluation of the condition.
func (c Condition) Eval(d ResourceData, key string) bool {
	return c(d, key)
}

// IsNewResource is a condition that evaluates to true if the resource access is
// new.
func IsNewResource() Condition {
	return func(d ResourceData, key string) bool {
		return d.IsNewResource()
	}
}

// HasChange is a condition that evaluates to true if the value accessed has
// changed.
func HasChange() Condition {
	return func(d ResourceData, key string) bool {
		return d.HasChange(key)
	}
}

// Any is a condition that evaluates to true if any of its child conditions
// evaluate to true. If it has no child conditions it will evaluate to true.
func Any(conditions ...Condition) Condition {
	return func(d ResourceData, key string) bool {
		if len(conditions) == 0 {
			return true
		}
		for _, condition := range conditions {
			if condition.Eval(d, key) {
				return true
			}
		}
		return false
	}
}

// All is a condition that evaluates to true if all of its child conditions
// evaluate to true.
func All(conditions ...Condition) Condition {
	return func(d ResourceData, key string) bool {
		for _, condition := range conditions {
			if !condition.Eval(d, key) {
				return false
			}
		}
		return true
	}
}

// String accesses the value held by key and type asserts it to a pointer to a
// string.
func String(d ResourceData, key string, conditions ...Condition) (s *string) {
	v, ok := d.GetOkExists(key)
	if ok && Any(conditions...).Eval(d, key) {
		s = auth0.String(v.(string))
	}
	return
}

// Int accesses the value held by key and type asserts it to a pointer to a
// int.
func Int(d ResourceData, key string, conditions ...Condition) (i *int) {
	v, ok := d.GetOkExists(key)
	if ok && Any(conditions...).Eval(d, key) {
		i = auth0.Int(v.(int))
	}
	return
}

// Bool accesses the value held by key and type asserts it to a pointer to a
// bool.
func Bool(d ResourceData, key string, conditions ...Condition) (b *bool) {
	v, ok := d.GetOkExists(key)
	if ok && Any(conditions...).Eval(d, key) {
		b = auth0.Bool(v.(bool))
	}
	return
}

// Slice accesses the value held by key and type asserts it to a slice.
func Slice(d ResourceData, key string, conditions ...Condition) (s []interface{}) {
	v, ok := d.GetOkExists(key)
	if ok && Any(conditions...).Eval(d, key) {
		s = v.([]interface{})
	}
	return
}

// Map accesses the value held by key and type asserts it to a map.
func Map(d ResourceData, key string, conditions ...Condition) (m map[string]interface{}) {
	v, ok := d.GetOkExists(key)
	if ok && Any(conditions...).Eval(d, key) {
		m = v.(map[string]interface{})
	}
	return
}

// List accesses the value held by key and returns an iterator able to go over
// its elements.
func List(d ResourceData, key string, conditions ...Condition) Iterator {
	v, ok := d.GetOkExists(key)
	if ok && Any(conditions...).Eval(d, key) {
		return &list{newResourceDataAtKey(key, d), v.([]interface{})}
	}
	return &list{}
}

// Set accesses the value held by key, type asserts it to a set and returns an
// iterator able to go over its elements.
func Set(d ResourceData, key string, conditions ...Condition) Iterator {
	v, ok := d.GetOkExists(key)
	if ok && Any(conditions...).Eval(d, key) {
		if s, ok := v.(*schema.Set); ok {
			return &set{newResourceDataAtKey(key, d), s}
		}
	}
	return &set{nil, &schema.Set{}}
}

// Iterator is used to iterate over a list or set.
//
// Elem iterates over all elements of the list or set, calling fn with each
// iteration. The callback takes a Data interface as argument which is prefixed
// with its parents key, allowing for convenient nested data access.
//
// Range iterates over all elements of the list, calling fn in each iteration.
//
// List returns the underlying list as a Go slice.
type Iterator interface {
	Elem(func(d ResourceData))
	Range(func(k int, v interface{}))
	List() []interface{}
}

type list struct {
	d ResourceData
	v []interface{}
}

func (l *list) Range(fn func(key int, value interface{})) {
	for key, value := range l.v {
		fn(key, value)
	}
}

func (l *list) Elem(fn func(ResourceData)) {
	for idx := range l.v {
		fn(newResourceDataAtIndex(idx, l.d))
	}
}

func (l *list) List() []interface{} {
	return l.v
}

type set struct {
	d ResourceData
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

func (s *set) Elem(fn func(ResourceData)) {
	for _, v := range s.s.List() {
		fn(newResourceDataAtKey(s.hash(v), s.d))
	}
}

func (s *set) List() []interface{} {
	return s.s.List()
}

// Diff accesses the value held by key and type asserts it to a set. It then
// compares it's changes if any and returns what needs to be added and what
// needs to be removed.
func Diff(d ResourceData, key string) (add []interface{}, rm []interface{}) {
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
func JSON(d ResourceData, key string, conditions ...Condition) (m map[string]interface{}, err error) {
	v, ok := d.GetOkExists(key)
	if ok && Any(conditions...).Eval(d, key) {
		m, err = structure.ExpandJsonFromString(v.(string))
		if err != nil {
			return
		}
	}
	return
}
