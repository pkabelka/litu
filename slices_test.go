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
