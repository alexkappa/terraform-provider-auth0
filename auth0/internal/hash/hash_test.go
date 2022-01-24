package hash

import "testing"

func TestStringKey(t *testing.T) {

	v := map[string]interface{}{
		"Foo": "Foo",
		"Bar": "Bar",
	}

	for key, expected := range map[string]int{
		"Foo": 3023971265,
		"Bar": 1320340042,
	} {
		t.Run(key, func(t *testing.T) {
			fn := StringKey(key)
			if fn(v) != expected {
				t.Errorf("expected %d to be %d", fn(v), expected)
			}
		})
	}
}
