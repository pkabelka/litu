package litu

import (
	"testing"
)

func TestMap(t *testing.T) {
    x := []string{"a", "b", "c"}

    res := Map(x, func(e string) string {
        return e + "1"
    })

    expect := []string{"a1", "b1", "c1"}
    if !Equal(res, expect) {
		t.Errorf("res = %q, want %q", res, expect)
    }
}

func TestFilter(t *testing.T) {
    x := []string{"a", "b", "c"}

    res := Filter(x, func(e string) bool {
        if e == "a" || e == "b" {
            return true
        }
        return false
    })
    expect := []string{"a", "b"}
    if !Equal(res, expect) {
		t.Errorf("res = %q, want %q", res, expect)
    }


    res = Filter(x, func(e string) bool {
        return true
    })
    if !Equal(res, x) {
		t.Errorf("res = %q, want %q", res, x)
    }


    res = Filter(x, func(e string) bool {
        return false
    })
    expect = []string{}
    if !Equal(res, expect) {
		t.Errorf("res = %q, want %q", res, expect)
    }
}
