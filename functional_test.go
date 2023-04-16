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
