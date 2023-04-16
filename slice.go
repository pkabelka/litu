package litu

func Insert[T ~[]U, U any](a T, idx int, e U) T {
    return append(a[:idx], append(T{e}, a[idx:]...)...)
}

func InsertFirst[T ~[]U, U any](a T, e U)  T {
    return append(T{e}, a...)
}

func Pop[T ~[]U, U any](a T) (U, T) {
    return a[len(a)-1], a[:len(a)-1]
}

func PopFirst[T ~[]U, U any](a T)  (U, T) {
    return a[0], a[1:]
}

func ReverseInplace[T ~[]U, U any](a T) {
    for i := len(a)/2-1; i >= 0; i-- {
        opp := len(a)-1-i
        a[i], a[opp] = a[opp], a[i]
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
    for i, e := range a {
        if e != b[i] {
            return false
        }
    }
    return true
}

func EqualUnordered[T comparable](a, b []T) bool {
    if len(a) != len(b) {
        return false
    }
    return All(a, b)
}
