package litu

type Set[T comparable] map[T]struct{}

func NewSet[T comparable](size int) Set[T] {
    return make(Set[T], size)
}

func NewSetFromSlice[T comparable](a []T) Set[T] {
    res := make(Set[T], len(a))
    for _, e := range a {
        res[e] = struct{}{}
    }
    return res
}

func (s Set[T]) Contains(e T) bool {
    _, ok := s[e]
    return ok
}

func (s Set[T]) Add(e T) {
    s[e] = struct{}{}
}

func (s Set[T]) Remove(e T) {
    delete(s, e)
}

func (s Set[T]) ToSlice() []T {
    res := make([]T, 0, len(s))
    for e := range s {
        res = append(res, e)
    }
    return res
}
