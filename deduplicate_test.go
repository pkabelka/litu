package litu

import (
	"testing"
)

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
