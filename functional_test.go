package litu

import (
	"testing"
)

func TestMap(t *testing.T) {
	x := []string{"a", "b", "c"}

	res := Map(x, func(e string) string {
		return e + "1"
	})

	expect := []string{"a1", "b1", "c1"}
	if !Equal(res, expect) {
		t.Errorf("res = %q, want %q", res, expect)
	}
}

func TestFilter(t *testing.T) {
	x := []string{"a", "b", "c"}

	res := Filter(x, func(e string) bool {
		if e == "a" || e == "b" {
			return true
		}
		return false
	})
	expect := []string{"a", "b"}
	if !Equal(res, expect) {
		t.Errorf("res = %q, want %q", res, expect)
	}

	res = Filter(x, func(e string) bool {
		return true
	})
	if !Equal(res, x) {
		t.Errorf("res = %q, want %q", res, x)
	}

	res = Filter(x, func(e string) bool {
		return false
	})
	expect = []string{}
	if !Equal(res, expect) {
		t.Errorf("res = %q, want %q", res, expect)
	}
}

func TestReduceInt(t *testing.T) {
	x := []int{1, 2, 3}

	res := Reduce(x, func(acc int, e int) int {
		return acc + e
	}, 0)
	if res != 6 {
		t.Errorf("res = %v, want %v", res, 6)
	}

	res = Reduce(x, func(acc int, e int) int {
		return acc * e
	}, 1)
	if res != 6 {
		t.Errorf("res = %v, want %v", res, 6)
	}

	res = Reduce(x, func(acc int, e int) int {
		return acc * e
	}, 0)
	if res != 0 {
		t.Errorf("res = %v, want %v", res, 0)
	}
}

func TestReduceFloat(t *testing.T) {
	x := []float64{1, 2, 3}

	res := Reduce(x, func(acc float64, e float64) float64 {
		return acc + e
	}, 0)
	if res != 6 {
		t.Errorf("res = %v, want %v", res, 6)
	}

	res = Reduce(x, func(acc float64, e float64) float64 {
		return acc * e
	}, 1)
	if res != 6 {
		t.Errorf("res = %v, want %v", res, 6)
	}

	res = Reduce(x, func(acc float64, e float64) float64 {
		return acc * e
	}, 0)
	if res != 0 {
		t.Errorf("res = %v, want %v", res, 0)
	}
}

func TestScanLeftInt(t *testing.T) {
	x := []int{1, 2, 3}

	res := ScanLeft(x, func(acc int, e int) int {
		return acc + e
	}, 0)

	expect := []int{1, 3, 6}
	if !Equal(res, expect) {
		t.Errorf("res = %v, want %v", res, expect)
	}

	res = ScanLeft(x, func(acc int, e int) int {
		return acc * e
	}, 0)

	expect = []int{0, 0, 0}
	if !Equal(res, expect) {
		t.Errorf("res = %v, want %v", res, expect)
	}

	res = ScanLeft(x, func(acc int, e int) int {
		return acc * e
	}, 1)

	expect = []int{1, 2, 6}
	if !Equal(res, expect) {
		t.Errorf("res = %v, want %v", res, expect)
	}
}

func TestScanLeftFloat(t *testing.T) {
	x := []float64{1, 2, 3}

	res := ScanLeft(x, func(acc float64, e float64) float64 {
		return acc + e
	}, 0)

	expect := []float64{1, 3, 6}
	if !Equal(res, expect) {
		t.Errorf("res = %v, want %v", res, expect)
	}

	res = ScanLeft(x, func(acc float64, e float64) float64 {
		return acc * e
	}, 0)

	expect = []float64{0, 0, 0}
	if !Equal(res, expect) {
		t.Errorf("res = %v, want %v", res, expect)
	}

	res = ScanLeft(x, func(acc float64, e float64) float64 {
		return acc * e
	}, 1)

	expect = []float64{1, 2, 6}
	if !Equal(res, expect) {
		t.Errorf("res = %v, want %v", res, expect)
	}
}

func TestScanRightInt(t *testing.T) {
	x := []int{1, 2, 3}

	res := ScanRight(x, func(acc int, e int) int {
		return acc + e
	}, 0)

	expect := []int{3, 5, 6}
	if !Equal(res, expect) {
		t.Errorf("res = %v, want %v", res, expect)
	}

	res = ScanRight(x, func(acc int, e int) int {
		return acc * e
	}, 0)

	expect = []int{0, 0, 0}
	if !Equal(res, expect) {
		t.Errorf("res = %v, want %v", res, expect)
	}

	res = ScanRight(x, func(acc int, e int) int {
		return acc * e
	}, 1)

	expect = []int{3, 6, 6}
	if !Equal(res, expect) {
		t.Errorf("res = %v, want %v", res, expect)
	}
}

func TestScanRightFloat(t *testing.T) {
	x := []float64{1, 2, 3}

	res := ScanRight(x, func(acc float64, e float64) float64 {
		return acc + e
	}, 0)

	expect := []float64{3, 5, 6}
	if !Equal(res, expect) {
		t.Errorf("res = %v, want %v", res, expect)
	}

	res = ScanRight(x, func(acc float64, e float64) float64 {
		return acc * e
	}, 0)

	expect = []float64{0, 0, 0}
	if !Equal(res, expect) {
		t.Errorf("res = %v, want %v", res, expect)
	}

	res = ScanRight(x, func(acc float64, e float64) float64 {
		return acc * e
	}, 1)

	expect = []float64{3, 6, 6}
	if !Equal(res, expect) {
		t.Errorf("res = %v, want %v", res, expect)
	}
}

func TestSumInt(t *testing.T) {
	x := []int{1, 2, 3}

	res := Sum(x)
	if res != 6 {
		t.Errorf("res = %v, want %v", res, 6)
	}
}

func TestSumFloat(t *testing.T) {
	x := []float64{1, 2, 3}

	res := Sum(x)
	if res != 6 {
		t.Errorf("res = %v, want %v", res, 6)
	}
}

func TestCumSumInt(t *testing.T) {
	x := []int{1, 2, 3}

	res := CumSum(x)
	expect := []int{1, 3, 6}
	if !Equal(res, expect) {
		t.Errorf("res = %v, want %v", res, expect)
	}
}

func TestCumSumFloat(t *testing.T) {
	x := []float64{1, 2, 3}

	res := CumSum(x)
	expect := []float64{1, 3, 6}
	if !Equal(res, expect) {
		t.Errorf("res = %v, want %v", res, expect)
	}
}
