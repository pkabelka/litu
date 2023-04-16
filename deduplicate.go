package litu

func DedupeSlice[T comparable](a []T) []T {
	set := NewSetFromSlice(a)
	return set.ToSlice()
}
