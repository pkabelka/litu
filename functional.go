package litu

func Map[T, U any](a []T, f func(e T) U) []U {
    res := make([]U, len(a))
    for i, e := range a {
        res[i] = f(e)
    }
    return res
}

func Filter[T any](a []T, f func(e T) bool) []T {
    res := make([]T, 0)
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
