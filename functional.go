package litu

func Map[T, U any](a []T, f func(e T) U) []U {
    res := make([]U, len(a))
    for i, e := range a {
        res[i] = f(e)
    }
    return res
}
