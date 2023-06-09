package litu

func Map[T, U any](a []T, f func(e T) U) []U {
	res := make([]U, len(a))
	for i, e := range a {
		res[i] = f(e)
	}
	return res
}

func Filter[T any](a []T, f func(e T) bool) []T {
	res := make([]T, 0, len(a))
	for _, e := range a {
		if val := f(e); val {
			res = append(res, e)
		}
	}
	return res
}

func Reduce[T, A any](a []T, f func(acc A, e T) A, acc A) A {
	for _, e := range a {
		acc = f(acc, e)
	}
	return acc
}

func ReduceRight[T, A any](a []T, f func(acc A, e T) A, acc A) A {
	for i := len(a) - 1; i >= 0; i-- {
		acc = f(acc, a[i])
	}
	return acc
}

func ScanLeft[T, A any](a []T, f func(acc A, e T) A, acc A) []A {
	res := make([]A, len(a))
	for i, e := range a {
		acc = f(acc, e)
		res[i] = acc
	}
	return res
}

func ScanRight[T, A any](a []T, f func(acc A, e T) A, acc A) []A {
	res := make([]A, len(a))
	for i := len(a) - 1; i >= 0; i-- {
		acc = f(acc, a[i])
		res[len(a)-1-i] = acc
	}
	return res
}

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 |
		~uint32 | ~uint64 | ~uintptr | ~float32 | ~float64
}

func Sum[T Number](a []T) T {
	return Reduce(a, func(acc T, e T) T {
		return acc + e
	}, 0)
}

func CumSum[T Number](a []T) []T {
	return ScanLeft(a, func(acc T, e T) T {
		return acc + e
	}, 0)
}
