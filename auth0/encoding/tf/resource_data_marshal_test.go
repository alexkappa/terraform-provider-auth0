package tf

import (
	"reflect"
	"testing"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"gopkg.in/auth0.v1"
)

type Test struct {
	Bool      *bool                  `tf:"bool"`
	Int       int                    `tf:"int"`
	Float     float64                `tf:"float"`
	String    string                 `tf:"string"`
	List      []interface{}          `tf:"list"`
	Set       []interface{}          `tf:"set"`
	Map       map[string]interface{} `tf:"map"`
	Interface interface{}            `tf:"interface,list"`
}

func TestMarshal(t *testing.T) {

	for _, test := range []struct {
		Name   string
		Schema map[string]*schema.Schema
		State  *terraform.InstanceState
		Diff   *terraform.InstanceDiff
		Value  Test
	}{
		{
			Name: "#1",

			Schema: map[string]*schema.Schema{
				"bool":      &schema.Schema{Type: schema.TypeBool},
				"int":       &schema.Schema{Type: schema.TypeInt},
				"float":     &schema.Schema{Type: schema.TypeFloat},
				"string":    &schema.Schema{Type: schema.TypeString},
				"list":      &schema.Schema{Type: schema.TypeList},
				"set":       &schema.Schema{Type: schema.TypeSet, Elem: &schema.Schema{Type: schema.TypeString}},
				"map":       &schema.Schema{Type: schema.TypeMap},
				"interface": &schema.Schema{Type: schema.TypeList, Elem: &schema.Schema{Type: schema.TypeInt}},
			},

			State: nil,

			Diff: nil,

			Value: Test{
				Bool:      auth0.Bool(true),
				Int:       123,
				Float:     12.3,
				String:    "foo",
				Map:       map[string]interface{}{"x": "123"},
				Set:       new(schema.Set).List(),
				List:      []interface{}{},
				Interface: []interface{}{1, 2, 3},
			},
		},
	} {
		t.Run(test.Name, func(t *testing.T) {

			d, _ := schema.InternalMap(test.Schema).Data(test.State, test.Diff)

			err := Marshal(d, test.Value)
			if err != nil {
				t.Fatal(err)
			}

			assertEqual(t, auth0.BoolValue(test.Value.Bool), d.Get("bool"))
			assertEqual(t, test.Value.Int, d.Get("int"))
			assertEqual(t, test.Value.Float, d.Get("float"))
			assertEqual(t, test.Value.String, d.Get("string"))
			assertEqual(t, test.Value.Map, d.Get("map"))
			assertEqual(t, test.Value.Set, d.Get("set").(*schema.Set).List())
			assertEqual(t, test.Value.List, d.Get("list"))
			assertEqual(t, test.Value.Interface, d.Get("interface"))
		})
	}
}

func assertEqual(t *testing.T, a, b interface{}) {
	if !reflect.DeepEqual(a, b) {
		t.Errorf("Expected %#v, but have %#v instead", a, b)
	}
}
