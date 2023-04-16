package litu

import (
	"fmt"
	"testing"
)

func TestSetFromSlice(t *testing.T) {
	x := []string{"a", "b", "c", "b", "c", "d"}

	set := NewSetFromSlice(x)

	expect := map[string]struct{}{"a": {}, "b": {}, "c": {}, "d": {}}
	if fmt.Sprint(set) != fmt.Sprint(expect) {
		t.Errorf("res = %q, want %q", set, expect)
	}
}

func TestSetAddRemove(t *testing.T) {
	set := NewSet[string](0)

	set.Add("a")

	expect := map[string]struct{}{"a": {}}
	if fmt.Sprint(set) != fmt.Sprint(expect) {
		t.Errorf("res = %q, want %q", set, expect)
	}

	set.Add("b")

	expect = map[string]struct{}{"a": {}, "b": {}}
	if fmt.Sprint(set) != fmt.Sprint(expect) {
		t.Errorf("res = %q, want %q", set, expect)
	}

	set.Remove("a")

	expect = map[string]struct{}{"b": {}}
	if fmt.Sprint(set) != fmt.Sprint(expect) {
		t.Errorf("res = %q, want %q", set, expect)
	}

	set.Remove("b")

	expect = map[string]struct{}{}
	if fmt.Sprint(set) != fmt.Sprint(expect) {
		t.Errorf("res = %q, want %q", set, expect)
	}
	if len(set) != 0 {
		t.Errorf("len() = %v, want %v", len(set), 0)
	}
}

func TestSetContains(t *testing.T) {
	x := []string{"a", "b", "c", "b", "c", "d"}

	set := NewSetFromSlice(x)

	if !set.Contains("b") {
		t.Errorf("set should contain %q", "b")
	}

	set.Remove("c")

	if set.Contains("c") {
		t.Errorf("set should not contain %q", "c")
	}
}

func TestSetToSlice(t *testing.T) {
	x := []string{"a", "b", "c", "b", "c", "d"}

	set := NewSetFromSlice(x)
	res := set.ToSlice()

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
