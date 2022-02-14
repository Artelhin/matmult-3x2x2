package src

import (
	"fmt"
	"testing"
)

func TestInvFieldK(t *testing.T) {
	cases := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for _, i := range cases {
		fmt.Println(i, "->", InvFieldK(i, 11))
	}
}

func TestModK(t *testing.T) {
	cases := []struct{
		x, k, res int
	}{
		{1, 2, 1},
		{3, 2, 1},
		{-1, 2, 1},
		{8, 5, 3},
		{-3, 5, 2},
		{-8, 5, 2},
	}
	for i := range cases {
		c := &cases[i]
		x := ModK(float64(c.x), c.k)
		if x != float64(c.res) {
			t.Errorf("%d mod %d: expected %d, got %f", c.x, c.k, c.res, x)
		}
	}
}