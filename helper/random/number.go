package random

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func Int(min, max int) int {
	return min + rand.Int()%(max-min+1)
}
