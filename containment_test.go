package litu

import (
    "testing"
)

func TestSliceContains(t *testing.T) {
    x := []string{"a", "b", "c"}

    res := InSlice(x, "b")
    if res != true {
		t.Errorf("res = %v, want %v", res, true)
    }

    res = InSlice(x, "d")
    if res != false {
		t.Errorf("res = %v, want %v", res, false)
    }
}

func TestMapContains(t *testing.T) {
    x := map[string]int{
        "foo": 1,
        "bar": 2,
    }

    res := InMap(x, 1)
    if res != true {
		t.Errorf("res = %v, want %v", res, true)
    }

    res = InMap(x, 3)
    if res != false {
		t.Errorf("res = %v, want %v", res, false)
    }
}
