package auth0

import (
	"testing"
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
