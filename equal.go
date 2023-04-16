package litu

func Equal[T comparable](a, b []T) bool {
    if len(a) != len(b) {
        return false
    }
    for i, e := range a {
        if e != b[i] {
            return false
        }
    }
    return true
}
