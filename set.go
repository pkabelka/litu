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

func (s Set[T]) Add(e ...T) {
	for _, v := range e {
		s[v] = struct{}{}
	}
}

func (s Set[T]) Remove(e ...T) {
	for _, v := range e {
		if !s.Contains(v) {
			continue
		}
		delete(s, v)
	}
}

func (s Set[T]) ToSlice() []T {
	res := make([]T, 0, len(s))
	for e := range s {
		res = append(res, e)
	}
	return res
}

func (s Set[T]) AddFrom(s2 ...Set[T]) {
	for _, s2_ := range s2 {
		for e := range s2_ {
			s.Add(e)
		}
	}
}

func (s Set[T]) Intersect(s2 Set[T]) Set[T] {
	res := NewSet[T](len(s) + len(s2))
	for e := range s {
		if s2.Contains(e) {
			res.Add(e)
		}
	}
	return res
}

func (s Set[T]) Union(s2 Set[T]) Set[T] {
	res := NewSet[T](len(s) + len(s2))
	for e := range s {
		res.Add(e)
	}
	for e := range s2 {
		res.Add(e)
	}
	return res
}

func (s Set[T]) LeftJoin(s2 Set[T]) Set[T] {
	res := NewSet[T](len(s) + len(s2))
	for e := range s {
		res.Add(e)
	}
	for e := range s2 {
		if s.Contains(e) {
			res.Add(e)
		}
	}
	return res
}

func (s Set[T]) Diff(s2 Set[T]) Set[T] {
	res := NewSet[T](len(s) + len(s2))
	for e := range s {
		if !s2.Contains(e) {
			res.Add(e)
		}
	}
	for e := range s2 {
		if !s.Contains(e) {
			res.Add(e)
		}
	}
	return res
}

func (s Set[T]) Equal(s2 Set[T]) bool {
	if len(s) != len(s2) {
		return false
	}
	for e := range s {
		if !s2.Contains(e) {
			return false
		}
	}
	return true
}

func (s Set[T]) SubsetOf(s2 Set[T]) bool {
	if len(s) > len(s2) {
		return false
	}
	for e := range s {
		if !s2.Contains(e) {
			return false
		}
	}
	return true
}

func (s Set[T]) Any(e ...T) bool {
	for _, v := range e {
		if s.Contains(v) {
			return true
		}
	}
	return false
}

func (s Set[T]) All(e ...T) bool {
	for _, v := range e {
		if !s.Contains(v) {
			return false
		}
	}
	return true
}
