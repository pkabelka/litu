package litu

import (
	"testing"
)

func TestWhere(t *testing.T) {
	x := []string{"a", "b", "c", "b"}

	indices := Where(x, func(e string) bool {
		return e == "b"
	})
	expect := []int{1, 3}
	if !Equal(indices, expect) {
		t.Errorf("indices = %v, want %v", indices, expect)
	}

	indices = Where(x, func(e string) bool {
		return e == "d"
	})
	expect = []int{}
	if !Equal(indices, expect) {
		t.Errorf("indices = %v, want %v", indices, expect)
	}
}

func TestWhereFirst(t *testing.T) {
	x := []string{"a", "b", "c", "b"}

	idx, res := WhereFirst(x, func(e string) bool {
		return e == "b"
	})
	if res != true {
		t.Errorf("res = %v, want %v", res, true)
	}
	if idx != 1 {
		t.Errorf("idx = %v, want %v", idx, 1)
	}

	_, res = WhereFirst(x, func(e string) bool {
		return e == "d"
	})
	if res != false {
		t.Errorf("res = %v, want %v", res, false)
	}
}

func TestWhereLast(t *testing.T) {
	x := []string{"a", "b", "c", "b"}

	idx, res := WhereLast(x, func(e string) bool {
		return e == "b"
	})
	if res != true {
		t.Errorf("res = %v, want %v", res, true)
	}
	if idx != 3 {
		t.Errorf("idx = %v, want %v", idx, 3)
	}

	idx, res = WhereLast(x, func(e string) bool {
		return e == "a"
	})
	if res != true {
		t.Errorf("res = %v, want %v", res, true)
	}
	if idx != 0 {
		t.Errorf("idx = %v, want %v", idx, 0)
	}

	_, res = WhereLast(x, func(e string) bool {
		return e == "d"
	})
	if res != false {
		t.Errorf("res = %v, want %v", res, false)
	}
}

func TestIndicesOf(t *testing.T) {
	x := []string{"a", "b", "c", "b"}

	indices := IndicesOf(x, "b")
	expect := []int{1, 3}
	if !Equal(indices, expect) {
		t.Errorf("indices = %v, want %v", indices, expect)
	}

	indices = IndicesOf(x, "d")
	expect = []int{}
	if !Equal(indices, expect) {
		t.Errorf("indices = %v, want %v", indices, expect)
	}
}

func TestIndexOfFirst(t *testing.T) {
	x := []string{"a", "b", "c", "b"}

	idx, res := IndexOfFirst(x, "b")
	if res != true {
		t.Errorf("res = %v, want %v", res, true)
	}
	if idx != 1 {
		t.Errorf("idx = %v, want %v", idx, 1)
	}

	_, res = IndexOfFirst(x, "d")
	if res != false {
		t.Errorf("res = %v, want %v", res, false)
	}
}

func TestIndexOfLast(t *testing.T) {
	x := []string{"a", "b", "c", "b"}

	idx, res := IndexOfLast(x, "b")
	if res != true {
		t.Errorf("res = %v, want %v", res, true)
	}
	if idx != 3 {
		t.Errorf("idx = %v, want %v", idx, 3)
	}

	idx, res = IndexOfLast(x, "a")
	if res != true {
		t.Errorf("res = %v, want %v", res, true)
	}
	if idx != 0 {
		t.Errorf("idx = %v, want %v", idx, 0)
	}

	_, res = IndexOfLast(x, "d")
	if res != false {
		t.Errorf("res = %v, want %v", res, false)
	}
}

func TestCount(t *testing.T) {
	x := []string{"a", "b", "c", "b"}

	count := Count(x, "b")
	if count != 2 {
		t.Errorf("count = %v, want %v", count, 2)
	}

	count = Count(x, "d")
	if count != 0 {
		t.Errorf("count = %v, want %v", count, 0)
	}
}

func TestCountWhere(t *testing.T) {
	x := []string{"a", "b", "c", "b"}

	count := CountWhere(x, func(e string) bool {
		return e == "b"
	})

	if count != 2 {
		t.Errorf("count = %v, want %v", count, 2)
	}

	count = CountWhere(x, func(e string) bool {
		return e == "d"
	})
	if count != 0 {
		t.Errorf("count = %v, want %v", count, 0)
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
