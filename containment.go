package litu

func InSlice[T comparable](a []T, el T) bool {
    for _, e := range a {
        if e == el {
            return true
        }
    }
    return false
}

func InMap[T comparable, U comparable](a map[U]T, el T) bool {
    for _, e := range a {
        if e == el {
            return true
        }
    }
    return false
}
