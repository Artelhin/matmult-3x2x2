package src

import "fmt"

type MatrixCache struct {
	RankCache map[string]int
	RankHits,
	RankMisses int
}

var cache = NewMatrixCache()

func (c *MatrixCache) CheckRank(rep string) (int, bool) {
	v, ok := c.RankCache[rep]
	if ok {
		c.RankHits++
	} else {
		c.RankMisses++
	}
	return v, ok
}

func (c *MatrixCache) PutRank(rep string, r int) {
	c.RankCache[rep] = r
}

func (c *MatrixCache) Stats() string {
	return fmt.Sprintf("---Cache stats---\nRank hits: %d\nRank misses: %d\n-----------------", c.RankHits, c.RankMisses)
}

func NewMatrixCache() *MatrixCache {
	c := new(MatrixCache)
	c.RankCache = make(map[string]int)
	c.RankHits = 0
	c.RankMisses = 0
	return c
}
