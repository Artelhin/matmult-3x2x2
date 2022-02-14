package src

import (
	"strconv"
	"strings"

	"gonum.org/v1/gonum/mat"
)

func hash(m map[string]int) string {
	s := ""
	s += strconv.FormatInt(int64(m["111"]), 16)
	s += strconv.FormatInt(int64(m["112"]), 16)
	s += strconv.FormatInt(int64(m["121"]), 16)
	s += strconv.FormatInt(int64(m["211"]), 16)
	s += strconv.FormatInt(int64(m["122"]), 16)
	s += strconv.FormatInt(int64(m["212"]), 16)
	s += strconv.FormatInt(int64(m["221"]), 16)
	s += strconv.FormatInt(int64(m["222"]), 16)
	return s
}

func Decode(h string) string {
	s := ""
	templ := []string{"111", "112", "121", "211", "122", "212", "221", "222"}
	for i, c := range h {
		s += templ[i] + ": " + string(c) + "; "
	}
	return strings.TrimRight(s, "; ")
}

// k is field size
func matrixHashK(x mat.Matrix, k int) int {
	h := float64(0)
	p := float64(k)
	r, c := x.Dims()
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			h = p*h + x.At(i, j)
		}
	}
	return int(h)
}

func matrixStrHashK(x mat.Matrix, k int) string {
	r, c := x.Dims()
	s := strconv.Itoa(r) + "." + strconv.Itoa(c) + "."
	s += strconv.Itoa(matrixHashK(x, k))
	return s
}

// representation is 'r.c.hash' where r, c, and hash is int and are size and hash respectively
func matrixHashKDecode(s string, k int) *mat.Dense {
	ss := strings.Split(s, ".")
	r, _ := strconv.ParseInt(ss[0], 10, 64)
	c, _ := strconv.ParseInt(ss[1], 10, 64)
	h, _ := strconv.ParseInt(ss[2], 10, 64)
	x := mat.NewDense(int(r), int(c), nil)
	t := int(h)
	for i := int(r) - 1; i >= 0; i-- {
		for j := int(c) - 1; j >= 0; j-- {
			e := t % k
			x.Set(i, j, float64(e))
			t = (t - e) / k
		}
	}
	return x
}

// string representation in field with size k
func matrixStrRepK(x mat.Matrix) string {
	r, c := x.Dims()
	s := strconv.Itoa(r) + "." + strconv.Itoa(c) + "."
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			s += strconv.Itoa(int(x.At(i, j))) + "."
		}
	}
	return strings.TrimRight(s, ".")
}

// does not check string for being valid representation
func matrixStrRepDecode(s string) *mat.Dense {
	ss := strings.Split(s, ".")
	r, _ := strconv.ParseInt(ss[0], 10, 64)
	c, _ := strconv.ParseInt(ss[1], 10, 64)
	data := make([]float64, 0, len(ss)-2)
	for _, v := range ss[2:] {
		n, _ := strconv.ParseFloat(v, 64)
		data = append(data, n)
	}
	return mat.NewDense(int(r), int(c), data)
}
