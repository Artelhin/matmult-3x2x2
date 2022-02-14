package src

import (
	"errors"

	"gonum.org/v1/gonum/mat"
)

var ErrInvalidParam = errors.New("invalid param")

func ParamMatrixesFieldK(k int, m, n, p, q, x, y, u float64) ([][]mat.Matrix, error) {
	if ModK(q*x, k) == ModK(p*y, k) ||
		(!IsZ((n*y+1)/(q*x-p*y)) ||
			!IsZ(y/(p*y-q*x)) ||
			!IsZ((n*x+p)/(p*y-q*x)) ||
			!IsZ(x/(q*x-p*y))) {
		return nil, ErrInvalidParam
	}
	data := [][][]float64{
		{ // 1
			{1, -1, 0, ModK(m+n, k), 1, -1},
			{1, 0, 0, 0},
			{0, 0, 1, 0, 0, 1},
		},
		{ // 2
			{0, 0, 0, 0, -1, 0},
			{1, -1, 0, 0},
			{0, 0, 0, 0, -u, 1},
		},
		{ // 3
			{0, 0, 1, 0, u, 0},
			{1, -1, 0, 0},
			{0, 0, 0, 0, -1, 0},
		},
		{ // 4
			{0, 0, 0, ModK((n*y+1)/(q*x-p*y), k), 0, ModK(y/(p*y-q*x), k)},
			{0, 0, 1, -1},
			{-p, 0, p, ModK(-m*x, k), -x, ModK(-n*x, k)},
		},
		{ // 5
			{0, 0, 0, ModK((n*x+p)/(p*y-q*x), k), 0, ModK(x/(q*x-p*y), k)},
			{0, 0, 1, -1},
			{-q, 0, q, ModK(-m*y, k), -y, ModK(-n*y, k)},
		},
		{ // 6
			{0, 1, 0, -m - n, 0, 1},
			{1, 0, 1, 0},
			{1, 0, 0, 0, 0, 0},
		},
		{ // 7
			{0, 0, 1, 0, 0, 0},
			{1, 0, 0, 0},
			{0, 1, 0, 0, 1, 0},
		},
		{ // 8
			{1, -1, 0, ModK(m+n, k), 0, -1},
			{1, 0, 0, 1},
			{1, 0, -1, 0, 0, -1},
		},
		{ // 9
			{0, 0, 0, 1, 0, 0},
			{0, 0, 1, 0},
			{m, 1, n, m, 1, n},
		},
		{ // 10
			{1, -1, 0, m, 0, 0},
			{0, 0, 0, 1},
			{-1, 0, 1, -1, 0, 1},
		},
		{ // 11
			{1, 0, 0, 0, 0, 0},
			{0, 1, 0, 1},
			{0, 0, 0, 1, 0, 0},
		},
	}
	res := make([][]mat.Matrix, 0, 11)
	for i := range data {
		el := make([]mat.Matrix, 0, 3)
		el = append(el, mat.NewDense(3, 2, data[i][0]))
		el = append(el, mat.NewDense(2, 2, data[i][1]))
		el = append(el, mat.NewDense(2, 3, data[i][2]))
		res = append(res, el)
	}
	return res, nil
}

func ParamMatrixes(m, n, p, q, x, y, u float64) ([][]mat.Matrix, error) {
	if q*x == p*y ||
		!(IsZ((n*y+1)/(q*x-p*y)) ||
			!IsZ(y/(p*y-q*x)) ||
			!IsZ((n*x+p)/(p*y-q*x)) ||
			!IsZ(x/(q*x-p*y))) {
		return nil, ErrInvalidParam
	}
	data := [][][]float64{
		{ // 1
			{1, -1, 0, m + n, 1, -1},
			{1, 0, 0, 0},
			{0, 0, 1, 0, 0, 1},
		},
		{ // 2
			{0, 0, 0, 0, -1, 0},
			{1, -1, 0, 0},
			{0, 0, 0, 0, -u, 1},
		},
		{ // 3
			{0, 0, 1, 0, u, 0},
			{1, -1, 0, 0},
			{0, 0, 0, 0, -1, 0},
		},
		{ // 4
			{0, 0, 0, (n*y + 1) / (q*x - p*y), 0, y / (p*y - q*x)},
			{0, 0, 1, -1},
			{-p, 0, p, -m * x, -x, -n * x},
		},
		{ // 5
			{0, 0, 0, (n*x + p) / (p*y - q*x), 0, x / (q*x - p*y)},
			{0, 0, 1, -1},
			{-q, 0, q, -m * y, -y, -n * y},
		},
		{ // 6
			{0, 1, 0, -m - n, 0, 1},
			{1, 0, 1, 0},
			{1, 0, 0, 0, 0, 0},
		},
		{ // 7
			{0, 0, 1, 0, 0, 0},
			{1, 0, 0, 0},
			{0, 1, 0, 0, 1, 0},
		},
		{ // 8
			{1, -1, 0, m + n, 0, -1},
			{1, 0, 0, 1},
			{1, 0, -1, 0, 0, -1},
		},
		{ // 9
			{0, 0, 0, 1, 0, 0},
			{0, 0, 1, 0},
			{m, 1, n, m, 1, n},
		},
		{ // 10
			{1, -1, 0, m, 0, 0},
			{0, 0, 0, 1},
			{-1, 0, 1, -1, 0, 1},
		},
		{ // 11
			{1, 0, 0, 0, 0, 0},
			{0, 1, 0, 1},
			{0, 0, 0, 1, 0, 0},
		},
	}
	res := make([][]mat.Matrix, 0, 11)
	for i := range data {
		el := make([]mat.Matrix, 0, 3)
		el = append(el, mat.NewDense(3, 2, data[i][0]))
		el = append(el, mat.NewDense(2, 2, data[i][1]))
		el = append(el, mat.NewDense(2, 3, data[i][2]))
		res = append(res, el)
	}
	return res, nil
}
