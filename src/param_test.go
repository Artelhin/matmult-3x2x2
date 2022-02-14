package src

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
	"testing"
)

func TestParamMatrixesFieldK(t *testing.T) {
	// 4  4  2 -3  2 -1  0
	mm, err := ParamMatrixesFieldK(5, 4, 4, 2,-3, 2, -1, 0)
	if err != nil {
		t.Fatal("invalid params")
	}
	for i := range mm {
		fmt.Println(i, "-----------------")
		for j := range mm[i] {
			display(mm[i][j].(*mat.Dense))
		}
	}
}
