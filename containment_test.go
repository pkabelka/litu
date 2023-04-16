package litu

import (
	"testing"
)

func TestFindInSlice(t *testing.T) {
	x := []string{"a", "b", "c"}

	idx, res := IndexOf(x, "b")
	if res != true {
		t.Errorf("res = %v, want %v", res, true)
	}
	if idx != 1 {
		t.Errorf("res = %v, want %v", idx, 1)
	}

	_, res = IndexOf(x, "d")
	if res != false {
		t.Errorf("res = %v, want %v", res, false)
	}
}

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

func TestFindInMap(t *testing.T) {
	x := map[string]int{
		"foo": 1,
		"bar": 2,
	}

	key, res := KeyOf(x, 1)
	if res != true {
		t.Errorf("res = %v, want %v", res, true)
	}
	if key != "foo" {
		t.Errorf("res = %v, want %v", key, "foo")
	}

	_, res = KeyOf(x, 3)
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

func TestAny(t *testing.T) {
	x := []string{"a", "b", "c"}

	res := Any(x, []string{"d", "c"})
	if res != true {
		t.Errorf("res = %v, want %v", res, true)
	}

	res = Any(x, []string{"d", "e"})
	if res != false {
		t.Errorf("res = %v, want %v", res, true)
	}
}

func TestAll(t *testing.T) {
	x := []string{"a", "b", "c"}

	res := All(x, []string{"c", "a"})
	if res != true {
		t.Errorf("res = %v, want %v", res, true)
	}

	res = All(x, []string{"b", "d"})
	if res != false {
		t.Errorf("res = %v, want %v", res, true)
	}
}

func TestIndexOfMin(t *testing.T) {
	x := []int{3, 5, 7, 11, 2, 4}

	min, idx := IndexOfMin(x)
	if min != 2 {
		t.Errorf("min = %v, want %v", min, 2)
	}
	if idx != 4 {
		t.Errorf("idx = %v, want %v", min, 4)
	}
}

func TestIndexOfMax(t *testing.T) {
	x := []int{3, 5, 7, 11, 2, 4}

	max, idx := IndexOfMax(x)
	if max != 11 {
		t.Errorf("max = %v, want %v", max, 11)
	}
	if idx != 3 {
		t.Errorf("idx = %v, want %v", max, 3)
	}
}

func TestMin(t *testing.T) {
	x := []int{3, 5, 7, 11, 2, 4}

	min := Min(x)
	if min != 2 {
		t.Errorf("min = %v, want %v", min, 2)
	}
}

func TestMax(t *testing.T) {
	x := []int{3, 5, 7, 11, 2, 4}

	max := Max(x)
	if max != 11 {
		t.Errorf("max = %v, want %v", max, 11)
	}
}
