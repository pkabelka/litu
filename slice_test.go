package litu

import (
	"testing"
)

func TestInsert(t *testing.T) {
	x := []string{"a", "b", "c"}

	res := Insert(x, 1, "d")

	expect := []string{"a", "d", "b", "c"}
	if !Equal(res, expect) {
		t.Errorf("res = %v, want %v", res, expect)
	}

	res = Insert(res, 1, "e", "f", "g")

	expect = []string{"a", "e", "f", "g", "d", "b", "c"}
	if !Equal(res, expect) {
		t.Errorf("res = %v, want %v", res, expect)
	}
}

func TestInsertFirst(t *testing.T) {
	x := []string{"a", "b", "c"}

	res := InsertFirst(x, "d")

	expect := []string{"d", "a", "b", "c"}
	if !Equal(res, expect) {
		t.Errorf("res = %v, want %v", res, expect)
	}
}

func TestPop(t *testing.T) {
	x := []string{"a", "b", "c"}

	last, res := Pop(x)

	if last != "c" {
		t.Errorf("res = %v, want %v", last, "c")
	}

	expect := []string{"a", "b"}
	if !Equal(res, expect) {
		t.Errorf("res = %v, want %v", res, expect)
	}
}

func TestPopFirst(t *testing.T) {
	x := []string{"a", "b", "c"}

	first, res := PopFirst(x)

	if first != "a" {
		t.Errorf("res = %v, want %v", first, "a")
	}

	expect := []string{"b", "c"}
	if !Equal(res, expect) {
		t.Errorf("res = %v, want %v", res, expect)
	}
}

func TestReverse(t *testing.T) {
	x := []string{"a", "b", "c"}

	res := Reverse(x)

	expect := []string{"a", "b", "c"}
	if !Equal(x, expect) {
		t.Errorf("x = %v, want %v", res, expect)
	}

	expect = []string{"c", "b", "a"}
	if !Equal(res, expect) {
		t.Errorf("res = %v, want %v", res, expect)
	}
}

func TestReverseInplace(t *testing.T) {
	x := []string{"a", "b", "c"}

	ReverseInplace(x)

	expect := []string{"c", "b", "a"}
	if !Equal(x, expect) {
		t.Errorf("res = %v, want %v", x, expect)
	}
}

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

	a = []string{"a", "b", "c", "b"}
	b = []string{"b", "c", "a", "c"}

	res = EqualUnordered(a, b)

	if res == true {
		t.Errorf("res = %v, want %v", res, true)
	}
}

func BenchmarkEqualUnordered(b *testing.B) {
	x := make([]uint64, 10000000)
	y := make([]uint64, 10000000)
	var i uint64
	for i = 0; i < 10000000; i++ {
		x[i] = i
		y[i] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		EqualUnordered(x, y)
	}
}

func TestTake(t *testing.T) {
	x := []string{"a", "b", "c", "d"}

	res, err := Take(x, []int{0, 2})
	if err != nil {
		t.Error(err.Error())
	}
	expect := []string{"a", "c"}
	if !Equal(res, expect) {
		t.Errorf("res = %v, want %v", res, expect)
	}

	_, err = Take(x, []int{0, 4})
	if err == nil {
		t.Error("Take() should cause error")
	}

	_, err = Take(x, []int{0, -1})
	if err == nil {
		t.Error("Take() should cause error")
	}
}

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

func TestInSlice(t *testing.T) {
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

func BenchmarkIndexOfFirst(b *testing.B) {
	x := make([]uint64, 100000000)
	var i uint64
	for i = 0; i < 100000000; i++ {
		x[i] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = IndexOfFirst(x, 99999998)
	}
}

func BenchmarkInSlice(b *testing.B) {
	x := make([]uint64, 100000000)
	var i uint64
	for i = 0; i < 100000000; i++ {
		x[i] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = InSlice(x, 99999998)
	}
}

func BenchmarkInSliceRawLoop(b *testing.B) {
	x := make([]uint64, 100000000)
	var i uint64
	for i = 0; i < 100000000; i++ {
		x[i] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, e := range x {
			if e == 99999998 {
				break
			}
		}
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

func BenchmarkAny(b *testing.B) {
	x := make([]uint64, 100000000)
	var i uint64
	for i = 0; i < 100000000; i++ {
		x[i] = i
	}
	y := []uint64{100000001, 100000002, 99999998}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = Any(x, y)
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

	idx, min := IndexOfMin(x)
	if min != 2 {
		t.Errorf("min = %v, want %v", min, 2)
	}
	if idx != 4 {
		t.Errorf("idx = %v, want %v", min, 4)
	}
}

func TestIndexOfMax(t *testing.T) {
	x := []int{3, 5, 7, 11, 2, 4}

	idx, max := IndexOfMax(x)
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

func TestDedupeSlice(t *testing.T) {
	x := []string{"a", "b", "c", "b", "c", "d"}

	res := DedupeSlice(x)

	expect := []string{"a", "b", "c", "d"}
	for _, e := range expect {
		occurences := 0
		for _, r := range res {
			if e == r {
				occurences++
			}
		}
		if occurences != 1 {
			t.Errorf("res = %q, want %q", res, expect)
		}
	}
}

func BenchmarkDedupeSlice(b *testing.B) {
	x := make([]uint64, 10000000)
	var i uint64
	var j uint64
	j = 0
	for i = 0; i < 10000000; i++ {
		x[i] = j
		j++
		if j == 100000 {
			j = 0
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		DedupeSlice(x)
	}
}

func BenchmarkDedupeSliceRaw(b *testing.B) {
	x := make([]uint64, 10000000)
	var i uint64
	var j uint64
	j = 0
	for i = 0; i < 10000000; i++ {
		x[i] = j
		j++
		if j == 100000 {
			j = 0
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		set := make(map[uint64]struct{}, len(x))
		res := make([]uint64, 0, len(x))
		for _, e := range x {
			if _, ok := set[e]; !ok {
				set[e] = struct{}{}
				res = append(res, e)
			}
		}
	}
}
