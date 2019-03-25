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
