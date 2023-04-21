package litu

func DedupeSlice[T comparable](a []T) []T {
	set := NewSet[T](len(a))
	res := make([]T, 0, len(a))
	for _, e := range a {
		if !set.Contains(e) {
			set.Add(e)
			res = append(res, e)
		}
	}
	return res
}
