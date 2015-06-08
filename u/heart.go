package u

import (
	"math/rand"
	"time"
)

var (
	e9 = 1e9
)

func Heart() time.Duration {
	rander := rand.New(rand.NewSource(time.Now().Unix()))
	n := float64(rander.Int63n(30)+10) * e9
	return time.Duration(n)
}
