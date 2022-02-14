package src

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

func CheckFieldK(k int) {
	c := 0
	all := 0
	family := make(map[string]int)
	for m := -k + 1; m < k; m++ {
		for n := -k + 1; n < k; n++ {
			for p := -k + 1; p < k; p++ {
				for q := -k + 1; q < k; q++ {
					for x := -k + 1; x < k; x++ {
						for y := -k + 1; y < k; y++ {
							for u := -k + 1; u < k; u++ {
								all++
								t, ranks, err := CheckParamsFieldK(k, m, n, p, q, x, y, u)
								if err != nil {
									//fmt.Println(m, n, p, q, x, y, u, "invalid")
									continue
								}
								if !t {
									h := hash(ranks)
									Out.Log(fmt.Sprintf("%s %2d %2d %2d %2d %2d %2d %2d give %v\n",
										h, m, n, p, q, x, y, u, ranks))
									c++
									if v, ok := family[h]; ok {
										family[h] = v + 1
									} else {
										family[h] = 1
									}
								} else {
									//Out.Log(fmt.Sprintln(m, n, p, q, x, y, u, "give same ranks"))
								}
							}
						}
					}
				}
			}
		}
	}
	Out.Log(fmt.Sprintf("%d different params found, %d processed\n", c, all))
	Out.Log(cache.Stats())
	Out.Report(fmt.Sprintf("field MOD %d rank families (%d found):\n", k, len(family)))
	for k, v := range family {
		Out.Report(fmt.Sprintln(k, Decode(k), "-", v))
	}
}

// true is same ranks, false as different
func CheckParams(m, n, p, q, x, y, u int) (bool, map[string]int, error) {
	pm, err := ParamMatrixes(float64(m), float64(n), float64(p), float64(q), float64(x), float64(y), float64(u))
	if err != nil {
		return false, nil, err
	}
	t, ranks := HasSameRanksAsStrassen(pm)
	return t, ranks, nil
}

// true is same ranks, false as different
func CheckParamsFieldK(k, m, n, p, q, x, y, u int) (bool, map[string]int, error) {
	pm, err := ParamMatrixesFieldK(k, float64(m), float64(n), float64(p), float64(q), float64(x), float64(y), float64(u))
	if err != nil {
		return false, nil, err
	}
	t, ranks := HasSameRanksAsStrassenFieldK(pm, k)
	return t, ranks, nil
}

func HasSameCounters(m1, m2 map[string]int) bool {
	for k, v := range m1 {
		if t, ok := m2[k]; !ok || v != t {
			return false
		}
	}
	for k, v := range m2 {
		if t, ok := m1[k]; !ok || v != t {
			return false
		}
	}
	return true
}

func HasSameRanks(m1, m2 [][]mat.Matrix) bool {
	mm1 := Matrix3NumberWithEachRank(m1)
	mm2 := Matrix3NumberWithEachRank(m2)
	return HasSameCounters(mm1, mm2)
}

func HasSameRanksAsStrassen(m [][]mat.Matrix) (bool, map[string]int) {
	mm := Matrix3NumberWithEachRank(m)
	sm := map[string]int{
		"222": 1,
		"111": 10,
	}
	return HasSameCounters(mm, sm), mm
}

func HasSameRanksFieldK(m1, m2 [][]mat.Matrix, k int) bool {
	mm1 := Matrix3NumberWithEachRankFieldK(m1, k)
	mm2 := Matrix3NumberWithEachRankFieldK(m2, k)
	return HasSameCounters(mm1, mm2)
}

func HasSameRanksAsStrassenFieldK(m [][]mat.Matrix, k int) (bool, map[string]int) {
	mm := Matrix3NumberWithEachRankFieldK(m, k)
	sm := map[string]int{
		"222": 1,
		"111": 10,
	}
	return HasSameCounters(mm, sm), mm
}
