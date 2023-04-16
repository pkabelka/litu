package litu

import (
	"testing"
)

func TestEqual(t *testing.T) {
    a := []string{"a", "b", "c"}
    b := []string{"a", "b", "c"}

    res := Equal(a, b)

    if res != true {
		t.Errorf("res = %v, want %v", res, true)
    }

    b = []string{"a", "b", "d"}

    res = Equal(a, b)

    if res != false {
		t.Errorf("res = %v, want %v", res, false)
    }

    b = []string{}

    res = Equal(a, b)

    if res != false {
		t.Errorf("res = %v, want %v", res, false)
    }

    a = []string{}

    res = Equal(a, b)

    if res != true {
		t.Errorf("res = %v, want %v", res, true)
    }
}

func TestEqualUnordered(t *testing.T) {
    a := []string{"a", "b", "c"}
    b := []string{"b", "c", "a"}

    res := EqualUnordered(a, b)

    if res != true {
		t.Errorf("res = %v, want %v", res, true)
    }

    b = []string{"a", "b", "d"}

    res = EqualUnordered(a, b)

    if res != false {
		t.Errorf("res = %v, want %v", res, false)
    }

    b = []string{}

    res = EqualUnordered(a, b)

    if res != false {
		t.Errorf("res = %v, want %v", res, false)
    }

    a = []string{}

    res = EqualUnordered(a, b)

    if res != true {
		t.Errorf("res = %v, want %v", res, true)
    }
}
