package litu

func Where[T any](a []T, f func(e T) bool) []int {
	indices := make([]int, 0, len(a))
	for i, e := range a {
		if ok := f(e); ok {
			indices = append(indices, i)
		}
	}
	return indices
}

func WhereFirst[T any](a []T, f func(e T) bool) (int, bool) {
	for i, e := range a {
		if ok := f(e); ok {
			return i, true
		}
	}
	return -1, false
}

func WhereLast[T any](a []T, f func(e T) bool) (int, bool) {
	for i := len(a) - 1; i >= 0; i-- {
		if ok := f(a[i]); ok {
			return i, true
		}
	}
	return -1, false
}

func IndicesOf[T comparable](a []T, el T) []int {
	indices := make([]int, 0, len(a))
	for i, e := range a {
		if e == el {
			indices = append(indices, i)
		}
	}
	return indices
}

func IndexOfFirst[T comparable](a []T, el T) (int, bool) {
	for i, e := range a {
		if e == el {
			return i, true
		}
	}
	return -1, false
}

func IndexOfLast[T comparable](a []T, el T) (int, bool) {
	for i := len(a) - 1; i >= 0; i-- {
		if a[i] == el {
			return i, true
		}
	}
	return -1, false
}

func Count[T comparable](a []T, el T) int {
	return len(IndicesOf(a, el))
}

func CountWhere[T comparable](a []T, f func(e T) bool) int {
	return len(Where(a, f))
}

func InSlice[T comparable](a []T, el T) bool {
	_, found := IndexOfFirst(a, el)
	return found
}

func KeyOf[T comparable, U comparable](a map[U]T, el T) (U, bool) {
	for i, e := range a {
		if e == el {
			return i, true
		}
	}
	var empty U
	return empty, false
}

func InMap[T comparable, U comparable](a map[U]T, el T) bool {
	_, found := KeyOf(a, el)
	return found
}

func Any[T comparable](a, b []T) bool {
	for _, e := range b {
		if InSlice(a, e) {
			return true
		}
	}
	return false
}

func All[T comparable](a, b []T) bool {
	for _, e := range b {
		if !InSlice(a, e) {
			return false
		}
	}
	return true
}

func IndexOfMin[T Number | ~string](a []T) (T, int) {
	var min T

	if len(a) == 0 {
		return min, -1
	}

	min = a[0]
	minIdx := 0

	for i, e := range a {
		if e < min {
			min = e
			minIdx = i
		}
	}
	return min, minIdx
}

func Min[T Number | ~string](a []T) T {
	min, _ := IndexOfMin(a)
	return min
}

func IndexOfMax[T Number | ~string](a []T) (T, int) {
	var max T

	if len(a) == 0 {
		return max, -1
	}

	max = a[0]
	maxIdx := 0

	for i, e := range a {
		if e > max {
			max = e
			maxIdx = i
		}
	}
	return max, maxIdx
}

func Max[T Number | ~string](a []T) T {
	max, _ := IndexOfMax(a)
	return max
}
