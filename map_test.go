package litu

import (
	"testing"
)

func TestKeys(t *testing.T) {
	x := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	res := Keys(x)

	expect := []string{"a", "b", "c"}
	if !EqualUnordered(res, expect) {
		t.Errorf("res = %v, want %v", res, expect)
	}
}

func TestValues(t *testing.T) {
	x := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	res := Values(x)

	expect := []int{1, 2, 3}
	if !EqualUnordered(res, expect) {
		t.Errorf("res = %v, want %v", res, expect)
	}
}

func TestEqualMap(t *testing.T) {
	x := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	y := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	if !EqualMap(x, y) {
		t.Errorf("EqualMap should return true")
	}

	w := map[string]int{
		"a": 1,
		"b": 2,
		"d": 3,
	}

	if EqualMap(x, w) {
		t.Errorf("EqualMap should return false")
	}

	z := map[string]int{
		"a": 1,
		"b": 2,
		"c": 4,
	}

	if EqualMap(x, z) {
		t.Errorf("EqualMap should return false")
	}
}

func TestAnyKeyOf(t *testing.T) {
	x := map[string]int{
		"foo": 1,
		"bar": 2,
	}

	key, res := AnyKeyOf(x, 1)
	if res != true {
		t.Errorf("res = %v, want %v", res, true)
	}
	if key != "foo" {
		t.Errorf("res = %v, want %v", key, "foo")
	}

	_, res = AnyKeyOf(x, 3)
	if res != false {
		t.Errorf("res = %v, want %v", res, false)
	}
}

func TestKeysOf(t *testing.T) {
	x := map[string]int{
		"foo": 1,
		"bar": 2,
		"baz": 1,
	}

	res := KeysOf(x, 1)
	expect := []string{"foo", "baz"}
	if !EqualUnordered(res, expect) {
		t.Errorf("res = %v, want %v", res, expect)
	}

	res = KeysOf(x, 3)
	expect = []string{}
	if !Equal(res, expect) {
		t.Errorf("res = %v, want %v", res, expect)
	}
}

func TestValInMap(t *testing.T) {
	x := map[string]int{
		"foo": 1,
		"bar": 2,
	}

	res := ValInMap(x, 1)
	if res != true {
		t.Errorf("res = %v, want %v", res, true)
	}

	res = ValInMap(x, 3)
	if res != false {
		t.Errorf("res = %v, want %v", res, false)
	}
}
