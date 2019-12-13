package auth0

import (
	"testing"
)

func TestMapData(t *testing.T) {
	d := MapData{
		"one":  1,
		"zero": 0,
	}

	if _, okExists := d.GetOkExists("one"); okExists != true {
		t.Error("unexpected value should return true")
	}

	if _, okExists := d.GetOkExists("zero"); okExists != false {
		t.Error("unexpected value should return false")
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
