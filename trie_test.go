package litu

import (
	"testing"
)

func TestTrieFromSlice(t *testing.T) {
	x := []string{"a", "b"}

	trie := Trie[string, string]{}
	trie.Insert("foo", x...)

	if len(trie.next) != 1 || trie.next["a"] == nil {
		t.Errorf("root is supposed to have 1 node in it")
	}

	if len(trie.next["a"].next) != 1 || trie.next["a"].next["b"] == nil {
		t.Errorf("node %q is supposed to have 1 node in it", "a")
	}

	if trie.next["a"].next["b"].Value != "foo" {
		t.Errorf("value of node %q = %q, want %q", "a->b",
			trie.next["a"].next["b"].Value, "foo")
	}
}

func TestTrieGet(t *testing.T) {
	x := []string{"a", "b", "c", "b", "c", "d"}

	trie := Trie[string, string]{}
	trie.Insert("foo", x...)

	expect := "foo"
	res := trie.Get(x...).Value
	if res != expect {
		t.Errorf("res = %q, want %q", res, expect)
	}
}

func TestTrieRemove(t *testing.T) {
	x := []string{"a", "b", "c", "b", "c", "d"}

	trie := Trie[string, string]{}
	trie.Insert("foo", x...)
	trie.Insert("bar", "a", "b", "c", "c")

	res := trie.Remove(x...)
	expect := true
	if res != expect {
		t.Errorf("res = %v, want %v", res, expect)
	}

	res = trie.Get("a").Remove("b", "c", "c")
	expect = true
	if res != expect {
		t.Errorf("res = %v, want %v", res, expect)
	}
}

func TestTrieEach(t *testing.T) {
	x := []string{"a", "b"}

	trie := Trie[string, string]{}
	trie.Insert("foo", x...)
	trie.Insert("bar", "c", "d")
	trie.Insert("baz", "a", "e")

	expect := []string{"a", "b", "c", "d", "e"}
	keys := make([]string, 0, 4)
	trie.Each(func(key string, node *Trie[string, string]) {
		keys = append(keys, key)
	})

	if !EqualUnordered(keys, expect) {
		t.Errorf("res = %v, want %v", keys, expect)
	}
}
