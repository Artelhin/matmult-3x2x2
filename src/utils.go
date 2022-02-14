package src

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
)

func IsZ(x float64) bool {
	return float64(int(x)) == x
}

func display(x *mat.Dense) {
	s := mat.Formatted(x)
	fmt.Printf("%v\n\n", s)
}

func ModK(x float64, k int) float64 {
	if x < 0 {
		return float64(k - (int(-x) % k))
	}
	return float64((k + int(x)) % k)
}

func InvFieldK(x float64, k int) float64 {
	return BinPow(x, k-2)
}

func BinPow(x float64, k int) float64 {
	switch k {
	case 0:
		return 1
	case 1:
		return x
	case 2:
		return x*x
	default:
		if k % 2 == 1 {
			return BinPow(x, k - 1) * x
		} else {
			b := BinPow(x, k / 2)
			return b * b
		}
	}
}