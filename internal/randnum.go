package internal

import (
	"math/rand"
	"time"
)

// RandNumber 随机数
func RandNumber(max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Int() % max
}
