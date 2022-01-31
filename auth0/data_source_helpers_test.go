package auth0

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

var newMockResourceSchema = map[string]*schema.Schema{
	"string_prop": {
		Type:        schema.TypeString,
		Description: "Some string property passed into mock schema",
		Required:    true,
	},
	"map_prop": {
		Type:     schema.TypeMap,
		Optional: true,
	},
	"bool_prop": {
		Type:     schema.TypeBool,
		Optional: true,
		Computed: false,
	},
	"list_prop": {
		Type:     schema.TypeList,
		Elem:     &schema.Schema{Type: schema.TypeString},
		Optional: true,
	},
	"float_prop": {
		Type:     schema.TypeFloat,
		Optional: true,
		Computed: false,
	},
	"set_prop": {
		Type:        schema.TypeSet,
		Optional:    true,
		Description: "Some set property passed into mock schema",
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"set_prop_child": {
					Type:     schema.TypeString,
					Required: true,
				},
			},
		},
	},
}

func TestDatasourceSchemaFromResourceSchema(t *testing.T) {
	dsSchema := datasourceSchemaFromResourceSchema(newMockResourceSchema)

	if len(dsSchema) != len(newMockResourceSchema) {
		t.Errorf("Unexpected number of properties in schema: got %v want %v", len(dsSchema), len(newMockResourceSchema))
	}

	for k, v := range dsSchema {
		if v.Optional == true {
			t.Errorf("Expected %v schema property to be required", k)
		}

		if v.Computed == false {
			t.Errorf("Expected %v schema property to be computed", k)
		}

		if v.Description != newMockResourceSchema[k].Description {
			t.Errorf("Description not being passed correctly, got %v want %v", v.Description, newMockResourceSchema[k].Description)
		}

		if v.Type != newMockResourceSchema[k].Type {
			t.Errorf("Unexpected number of properties in schema: got %v want %v", len(dsSchema), len(newMockResourceSchema))
		}

		if (k == "list_prop" || k == "set_prop") && v.Elem == nil {
			t.Errorf("Non-nil elements passed into list or set type properties")
		}
	}
}
