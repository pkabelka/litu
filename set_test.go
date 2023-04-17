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

	set.Add("c", "d")

	expect = map[string]struct{}{"a": {}, "b": {}, "c": {}, "d": {}}
	if fmt.Sprint(set) != fmt.Sprint(expect) {
		t.Errorf("res = %q, want %q", set, expect)
	}

	set.Remove("c", "d")

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

func TestSetAddFrom(t *testing.T) {
	x := []string{"a", "b", "c", "d"}
	y := []string{"c", "d", "e", "f"}

	set1 := NewSetFromSlice(x)
	set2 := NewSetFromSlice(y)

	set1.AddFrom(set2)

	expect := map[string]struct{}{"a": {}, "b": {}, "c": {}, "d": {}, "e": {}, "f": {}}
	if !EqualMap(set1, expect) {
		t.Errorf("res = %q, want %q", set1, expect)
	}

	y = []string{"f", "g", "h"}
	z := []string{"i", "j", "k"}

	set2 = NewSetFromSlice(y)
	set3 := NewSetFromSlice(z)

	set1.AddFrom(set2, set3)

	expect = map[string]struct{}{"a": {}, "b": {}, "c": {}, "d": {}, "e": {},
		"f": {}, "g": {}, "h": {}, "i": {}, "j": {}, "k": {}}
	if !EqualMap(set1, expect) {
		t.Errorf("res = %q, want %q", set1, expect)
	}
}

func TestSetIntersect(t *testing.T) {
	x := []string{"a", "b", "c", "d"}
	y := []string{"c", "d", "e", "f"}

	set1 := NewSetFromSlice(x)
	set2 := NewSetFromSlice(y)

	res := set1.Intersect(set2)

	expect := map[string]struct{}{"c": {}, "d": {}}
	if !EqualMap(res, expect) {
		t.Errorf("res = %q, want %q", res, expect)
	}
}

func TestSetUnion(t *testing.T) {
	x := []string{"a", "b", "c", "d"}
	y := []string{"c", "d", "e", "f"}

	set1 := NewSetFromSlice(x)
	set2 := NewSetFromSlice(y)

	res := set1.Union(set2)

	expect := map[string]struct{}{"a": {}, "b": {}, "c": {}, "d": {}, "e": {}, "f": {}}
	if !EqualMap(res, expect) {
		t.Errorf("res = %q, want %q", res, expect)
	}
}

func TestSetLeftJoin(t *testing.T) {
	x := []string{"a", "b", "c", "d"}
	y := []string{"c", "d", "e", "f"}

	set1 := NewSetFromSlice(x)
	set2 := NewSetFromSlice(y)

	res := set1.LeftJoin(set2)

	expect := map[string]struct{}{"a": {}, "b": {}, "c": {}, "d": {}}
	if !EqualMap(res, expect) {
		t.Errorf("res = %q, want %q", res, expect)
	}
}

func TestSetEqual(t *testing.T) {
	x := []string{"a", "b", "c", "d"}
	y := []string{"d", "c", "a", "b"}

	set1 := NewSetFromSlice(x)
	set2 := NewSetFromSlice(y)

	if !set1.Equal(set2) {
		t.Errorf("sets should be equal")
	}

	y = []string{"d", "c", "a", "a", "b"}

	set1 = NewSetFromSlice(x)
	set2 = NewSetFromSlice(y)

	if !set1.Equal(set2) {
		t.Errorf("sets should be equal")
	}

	y = []string{"d", "c", "a", "a"}

	set1 = NewSetFromSlice(x)
	set2 = NewSetFromSlice(y)

	if set1.Equal(set2) {
		t.Errorf("sets should not be equal")
	}

	y = []string{}

	set1 = NewSetFromSlice(x)
	set2 = NewSetFromSlice(y)

	if set1.Equal(set2) {
		t.Errorf("sets should not be equal")
	}

	x = []string{}

	set1 = NewSetFromSlice(x)

	if !set1.Equal(set2) {
		t.Errorf("sets should be equal")
	}
}

func TestSetSubsetOf(t *testing.T) {
	x := []string{"a", "b", "c", "d"}
	y := []string{"d", "c", "a", "b"}

	set1 := NewSetFromSlice(x)

	if !set1.SubsetOf(set1) {
		t.Errorf("a set should be a subset of itself")
	}

	set2 := NewSetFromSlice(y)

	if !set1.SubsetOf(set2) {
		t.Errorf("set1 should be a subset of set2")
	}

	y = []string{"c", "a", "a", "b"}

	set2 = NewSetFromSlice(y)

	if !set2.SubsetOf(set1) {
		t.Errorf("set1 should be a subset of set2")
	}

	x = []string{"a", "a", "b"}

	set1 = NewSetFromSlice(x)

	if set2.SubsetOf(set1) {
		t.Errorf("set1 should not be a subset of set2")
	}

	set1 = NewSet[string](0)

	if !set1.SubsetOf(set2) {
		t.Errorf("set1 should be a subset of set2")
	}

	set2 = NewSet[string](0)

	if !set1.SubsetOf(set2) {
		t.Errorf("set1 should be a subset of set2")
	}
}
