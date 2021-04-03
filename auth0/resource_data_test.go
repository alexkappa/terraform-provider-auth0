package auth0

import (
	"errors"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestMapData(t *testing.T) {
	d := MapData{
		"one":  1,
		"zero": 0,
		"nil":  nil,
	}

	t.Run("GetOk", func(t *testing.T) {
		for key, shouldBeOk := range map[string]bool{
			"one":       true,
			"zero":      false,
			"nil":       false,
			"undefined": false,
		} {
			if _, ok := d.GetOk(key); ok != shouldBeOk {
				t.Errorf("d.GetOk(%s) should report ok == %t", key, shouldBeOk)
			}
		}
	})

	t.Run("GetOkExists", func(t *testing.T) {
		for key, shouldBeOk := range map[string]bool{
			"one":       true,
			"zero":      true,
			"nil":       false,
			"undefined": false,
		} {
			if _, ok := d.GetOkExists(key); ok != shouldBeOk {
				t.Errorf("d.GetOkExists(%s) should report ok == %t", key, shouldBeOk)
			}
		}
	})

}

func TestJSON(t *testing.T) {
	d := MapData{"json": `{"foo": 123}`}
	v, err := JSON(d, "json")
	if err != nil {
		t.Error(err)
	}
	j, ok := v["foo"]
	if !ok {
		t.Errorf("Expected result to be a int, instead it was %T\n", j)
	}
}

func TestIsNil(t *testing.T) {

	for _, v := range []interface{}{
		nil,
		(*bool)(nil),
		(*string)(nil),
		(*int)(nil),
		(*int64)(nil),
		(*float32)(nil),
		(*float64)(nil),
		(*struct{})(nil),
		(*interface{})(nil),
		(interface{})(nil),
		([]interface{})(nil),
	} {
		if !isNil(v) {
			t.Errorf("Expected isNil(%#v) to return true", v)
		}
	}
}

func checkDataSourceStateMatchesResourceState(dataSourceName, resourceName string, includeFields []string) func(*terraform.State) error {
	return func(s *terraform.State) error {
		ds, ok := s.RootModule().Resources[dataSourceName]
		if !ok {
			return fmt.Errorf("can't find %s in state", dataSourceName)
		}

		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("can't find %s in state", resourceName)
		}

		dsAttr := ds.Primary.Attributes
		rsAttr := rs.Primary.Attributes

		errMsg := ""

		for _, k := range includeFields {
			if k == "%" {
				continue
			}
			if dsAttr[k] != rsAttr[k] {
				// ignore data sources where an empty list is being compared against a null list.
				if k[len(k)-1:] == "#" && (dsAttr[k] == "" || dsAttr[k] == "0") && (rsAttr[k] == "" || rsAttr[k] == "0") {
					continue
				}
				errMsg += fmt.Sprintf("%s is %s; want %s\n", k, dsAttr[k], rsAttr[k])
			}
		}

		if errMsg != "" {
			return errors.New(errMsg)
		}

		return nil
	}
}
