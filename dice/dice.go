package dice

import (
	"math/rand"
	"time"
)

func NdM(n int, m int) []int {
	rolls := make([]int, n)
	for i := 0; i < n; i++ {
		rolls[i] = Roll(m)
	}
	return rolls
}

func Roll(max int) int {
	seed := time.Now().UnixNano()
	rand.Seed(seed)
	// instead of returning [0,max), we want to return [1,max]
	return rand.Intn(max) + 1
}
