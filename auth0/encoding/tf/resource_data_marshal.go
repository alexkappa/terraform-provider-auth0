package tf

import (
	"reflect"
	"strings"

	"github.com/alexkappa/errors"
	"github.com/hashicorp/terraform/helper/schema"
)

type Marshaler interface {
	MarshalTF(d *schema.ResourceData) error
}

type Type string

const (
	TypeBool   Type = "bool"
	TypeInt         = "int"
	TypeFloat       = "float"
	TypeString      = "string"
	TypeList        = "list"
	TypeSet         = "set"
	TypeMap         = "map"
)

// Marshal reads the contents of `v` recursively and populates `d`. It relies on
// `tf` annotations defined in the type of `v`.
//
// Annotations can be in the form `tf:"key,type"`
//
func Marshal(d *schema.ResourceData, v interface{}) error {

	rv := reflect.ValueOf(v)

	if rv.Kind() != reflect.Struct {
		return errors.Errorf("Expected a struct type but instead got %s", rv.Kind())
	}

	for i := 0; i < rv.NumField(); i++ {

		field := rv.Type().Field(i)

		if tag, ok := field.Tag.Lookup("tf"); ok {

			key, typ, err := parseTag(tag)
			if err != nil {
				return errors.Wrap(err, "unable to parse tag")
			}

			value := castValue(typ, rv.Field(i))

			if err := d.Set(key, value); err != nil {
				return errors.Wrapf(err, "cannot set %q to %v", key, value)
			}
		}
	}

	return nil
}

func parseTag(tag string) (key string, typ Type, err error) {

	parts := strings.Split(tag, ",")

	if len(parts) == 0 {
		err = errors.New("empty tag")
	}

	if len(parts) > 0 {
		key = parts[0]
	}

	if len(parts) > 1 {
		typ = Type(parts[1])
	}

	return
}

var marshaler = reflect.TypeOf(new(Marshaler)).Elem()

func castValue(t Type, v reflect.Value) interface{} {
	switch t {
	case TypeString:
		return v.String()
	case TypeInt:
		return v.Int()
	case TypeFloat:
		return v.Float()
	case TypeBool:
		return v.Bool()
	case TypeSet:
		return v.Interface()
	case TypeList:
		return v.Interface()
	case TypeMap:
		return v.Interface().(map[string]interface{})
	}

	if v.Type().Implements(marshaler) {
		v.Interface().(Marshaler).MarshalTF(d)
	}

	return v.Interface()
}

func isList(v reflect.Value) bool {
	return v.Kind() == reflect.Slice || v.Kind() == reflect.Array
}

func Unmarshal(d *schema.ResourceData, v interface{}) error {

	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return errors.New(reflect.TypeOf(v).String())
	}

	return nil
}
