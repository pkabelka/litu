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
