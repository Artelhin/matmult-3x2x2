package src

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
	"testing"
)

func TestParamMatrixesFieldK(t *testing.T) {
	// 4  4  2 -3  2 -1  0
	k, m, n, p, q, x, y, u := 13, 0,  0, -6, -1, -1,  0, -6
	mm, err := ParamMatrixesFieldK(k, float64(m), float64(n), float64(p), float64(q), float64(x), float64(y), float64(u))
	if err != nil {
		t.Fatal("invalid params")
	}
	b, ranks, err := CheckParamsFieldK(k, m, n, p, q, x, y, u)
	if err != nil {
		//fmt.Println(m, n, p, q, x, y, u, "invalid")
	}
	if !b {
		h := hash(ranks)
		fmt.Printf("%s give %v\n", h, ranks)
	} else {
		//Out.Log(fmt.Sprintln(m, n, p, q, x, y, u, "give same ranks"))
	}
	for i := range mm {
		fmt.Println(i, "-----------------")
		for j := range mm[i] {
			display(mm[i][j].(*mat.Dense))
		}
	}
}
