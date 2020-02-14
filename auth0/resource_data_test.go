package auth0

import (
	"testing"
)

func TestMapData(t *testing.T) {
	d := MapData{
		"one":  1,
		"zero": 0,
	}

	for key, shouldBeOk := range map[string]bool{
		"one":  true,
		"zero": false,
	} {
		if _, ok := d.GetOkExists(key); ok != shouldBeOk {
			t.Errorf("d.GetOkExists(%s) should retport ok == %t", key, shouldBeOk)
		}
	}
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
