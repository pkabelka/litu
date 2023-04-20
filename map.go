package litu

func Keys[T ~map[K]V, K comparable, V any](a T) []K {
	res := make([]K, 0, len(a))
	for k := range a {
		res = append(res, k)
	}
	return res
}

func Values[T ~map[K]V, K comparable, V any](a T) []V {
	res := make([]V, 0, len(a))
	for _, v := range a {
		res = append(res, v)
	}
	return res
}

func EqualMap[A, B ~map[K]V, K, V comparable](a A, b B) bool {
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if e, ok := b[k]; !ok || v != e {
			return false
		}
	}
	return true
}

func KeysOf[K, V comparable](m map[K]V, v V) []K {
	res := make([]K, 0, len(m))
	for k, e := range m {
		if e == v {
			res = append(res, k)
		}
	}
	return res
}

func AnyKeyOf[K, V comparable](a map[K]V, v V) (K, bool) {
	for k, e := range a {
		if e == v {
			return k, true
		}
	}
	var empty K
	return empty, false
}

func ValInMap[K, V comparable](a map[K]V, v V) bool {
	_, found := AnyKeyOf(a, v)
	return found
}
