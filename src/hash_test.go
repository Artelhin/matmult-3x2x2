package src

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
	"testing"
)

func TestHash(t *testing.T) {
	cases := []map[string]int{
		{"111": 10, "222": 1},
		{"211": 1, "222": 1, "111": 9},
	}
	for i := range cases {
		fmt.Println(hash(cases[i]))
	}
}

func TestEncodeDecodeMatrix(t *testing.T) {
	m := mat.NewDense(3, 2, []float64{1, 0, 0, 1, 0, 0})
	s := matrixStrRepK(m)
	mm := matrixStrRepDecode(s)
	fmt.Println(matrixHashK(m, 2))
	display(m)
	display(mm)
}

func TestMatrixHash(t *testing.T) {
	m1 := mat.NewDense(3, 2, []float64{1, 0, 0, 1, 0, 0})
	m2 := mat.NewDense(2, 3, []float64{1, 0, 0, 1, 0, 0})
	fmt.Println(matrixHashK(m1, 2), matrixHashK(m2, 2))
}

func TestMatrixHashDecode(t *testing.T) {
	m := mat.NewDense(3, 2, []float64{1, 0, 0, 1, 0, 0})
	h := matrixStrHashK(m, 2)
	dm := matrixHashKDecode(h, 2)
	fmt.Println(h)
	display(m)
	display(dm)
}