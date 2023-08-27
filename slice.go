package litu

import (
	"errors"
	"sort"
)

func Insert[T ~[]U, U any](a T, idx int, e ...U) T {
	return append(a[:idx], append(e, a[idx:]...)...)
}

func InsertFirst[T ~[]U, U any](a T, e U) T {
	return append(T{e}, a...)
}

func Pop[T ~[]U, U any](a T) (U, T) {
	return a[len(a)-1], a[:len(a)-1]
}

func PopFirst[T ~[]U, U any](a T) (U, T) {
	return a[0], a[1:]
}

func ReverseInplace[T ~[]U, U any](a T) {
	for i := len(a)/2 - 1; i >= 0; i-- {
		j := len(a) - 1 - i
		a[i], a[j] = a[j], a[i]
	}
}

func Reverse[T ~[]U, U any](a T) T {
	rev := make(T, len(a))
	copy(rev, a)
	ReverseInplace(rev)
	return rev
}

func Equal[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func EqualUnordered[T Number | ~string](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	aCopy := make([]T, len(a))
	copy(aCopy, a)
	bCopy := make([]T, len(b))
	copy(bCopy, b)

	sort.Slice(aCopy, func(i, j int) bool {
		return aCopy[i] < aCopy[j]
	})
	sort.Slice(bCopy, func(i, j int) bool {
		return bCopy[i] < bCopy[j]
	})
	return Equal(aCopy, bCopy)
}

func Take[T any](a []T, indices []int) ([]T, error) {
	res := make([]T, 0, len(indices))
	for _, i := range indices {
		if i < 0 || i >= len(a) {
			return []T{}, errors.New("index out of range")
		}
		res = append(res, a[i])
	}
	return res, nil
}

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
	count := 0
	for _, e := range a {
		if e == el {
			count++
		}
	}
	return count
}

func CountWhere[T any](a []T, f func(e T) bool) int {
	count := 0
	for _, e := range a {
		if ok := f(e); ok {
			count++
		}
	}
	return count
}

func InSlice[T comparable](a []T, e T) bool {
	_, found := IndexOfFirst(a, e)
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

func IndexOfMin[T Number | ~string](el []T) (int, T) {
	var min T

	if len(el) == 0 {
		return -1, min
	}

	min = el[0]
	minIdx := 0

	for i, e := range el {
		if e < min {
			min = e
			minIdx = i
		}
	}
	return minIdx, min
}

func Min[T Number | ~string](e ...T) T {
	_, min := IndexOfMin(e)
	return min
}

func IndexOfMax[T Number | ~string](el []T) (int, T) {
	var max T

	if len(el) == 0 {
		return -1, max
	}

	max = el[0]
	maxIdx := 0

	for i, e := range el {
		if e > max {
			max = e
			maxIdx = i
		}
	}
	return maxIdx, max
}

func Max[T Number | ~string](e ...T) T {
	_, max := IndexOfMax(e)
	return max
}

func Deduplicate[T comparable](a []T) []T {
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
