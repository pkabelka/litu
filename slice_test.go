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
