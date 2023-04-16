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
