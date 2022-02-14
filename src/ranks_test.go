package src

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
	"testing"
)

func TestRanks(t *testing.T) {
	m := StrassenMatrixes()
	for i := range m {
		fmt.Print(Rank(m[i][0]))
		fmt.Print(Rank(m[i][1]))
		fmt.Println(Rank(m[i][2]))
	}
}

func TestRanks2(t *testing.T) {
	m := mat.NewDense(3, 2, []float64{0, 1, 0, -1, 0, 0})
	display(m)
	fmt.Println(Rank(m))
}

func TestRankFieldK(t *testing.T) {
	m := StrassenMatrixes()
	for i := range m {
		fmt.Print(RankFieldK(m[i][0], 2))
		fmt.Print(RankFieldK(m[i][1], 2))
		fmt.Println(RankFieldK(m[i][2], 2))
	}
}

func TestRanksFieldK2(t *testing.T) {
	m := mat.NewDense(3, 2, []float64{1, 0, 0, 1, 0, 0})
	display(m)
	fmt.Println(RankFieldK(m, 2))
}