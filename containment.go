package litu

func IndexOf[T comparable](a []T, el T) (int, bool) {
    for i, e := range a {
        if e == el {
            return i, true
        }
    }
    return -1, false
}

func InSlice[T comparable](a []T, el T) bool {
    _, found := IndexOf(a, el)
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
