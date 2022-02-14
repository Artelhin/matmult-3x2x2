package src

import (
	"strconv"

	"gonum.org/v1/gonum/mat"
)

func Matrix3NumberWithEachRank(m [][]mat.Matrix) map[string]int {
	nums := make(map[string]int)
	for i := range m {
		s := CalculateRanksFor3(m[i])
		if v, ok := nums[s]; ok {
			nums[s] = v + 1
		} else {
			nums[s] = 1
		}
	}
	return nums
}

//string format is xxx where each x is the rank of respective matrix
func CalculateRanksFor3(m []mat.Matrix) string {
	s := ""
	for i := range m {
		s += strconv.Itoa(Rank(m[i]))
	}
	return s
}

func Matrix3NumberWithEachRankFieldK(m [][]mat.Matrix, k int) map[string]int {
	nums := make(map[string]int)
	for i := range m {
		s := CalculateRanksFor3FieldK(m[i], k)
		if v, ok := nums[s]; ok {
			nums[s] = v + 1
		} else {
			nums[s] = 1
		}
	}
	return nums
}

//string format is xxx where each x is the rank of respective matrix
func CalculateRanksFor3FieldK(m []mat.Matrix, k int) string {
	s := ""
	for i := range m {
		s += strconv.Itoa(RankFieldK(m[i], k))
	}
	return s
}

func RankFieldK(matrix mat.Matrix, k int) int {
	// check if it was already calculated
	rep := matrixStrHashK(matrix, k)
	if rank, ok := cache.CheckRank(rep); ok {
		return rank
	}

	r, c := matrix.Dims()
	x := mat.NewDense(r, c, nil)
	x.CloneFrom(matrix)
	rank := c

	for row := 0; row < rank; row++ {
		//display(x)
		if row >= r {
			rank = r
			break
		}
		if x.At(row, row) != 0 {
			for col := 0; col < r; col++ {
				if col != row {
					// mult := x.At(col, row) / x.At(row, row)
					mult := InvFieldK(x.At(row, row), k) * x.At(col, row)
					for i := 0; i < rank; i++ {
						x.Set(col, i, ModK(x.At(col, i)-mult*x.At(row, i), k))
					}
				}
			}
		} else {
			reduce := true
			for i := row + 1; i < r; i++ {
				if x.At(i, row) != 0 {
					swap(x, row, i, rank)
					reduce = false
					break
				}
			}
			if reduce {
				rank--
				for i := 0; i < r; i++ {
					x.Set(i, row, x.At(i, rank))
				}
			}
			row--
		}
	}
	// memorize
	cache.PutRank(rep, rank)
	return rank
}

func Rank(matrix mat.Matrix) int {
	r, c := matrix.Dims()
	x := mat.NewDense(r, c, nil)
	x.CloneFrom(matrix)
	rank := c

	for row := 0; row < rank; row++ {
		//display(x)
		if row >= r {
			rank = r
			break
		}
		if x.At(row, row) != 0 {
			for col := 0; col < r; col++ {
				if col != row {
					mult := x.At(col, row) / x.At(row, row)
					for i := 0; i < rank; i++ {
						x.Set(col, i, x.At(col, i)-mult*x.At(row, i))
					}
				}
			}
		} else {
			reduce := true
			for i := row + 1; i < r; i++ {
				if x.At(i, row) != 0 {
					swap(x, row, i, rank)
					reduce = false
					break
				}
			}
			if reduce {
				rank--
				for i := 0; i < r; i++ {
					x.Set(i, row, x.At(i, rank))
				}
			}
			row--
		}
	}
	return rank
}

func swap(x *mat.Dense, row1, row2, col int) {
	for i := 0; i < col; i++ {
		temp := x.At(row1, i)
		x.Set(row1, i, x.At(row2, i))
		x.Set(row2, i, temp)
	}
}
