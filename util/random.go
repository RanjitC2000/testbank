package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())
}

// RandInt generates a random number between min and max
func RandInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandString generates a random string of length n
func RandString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandOwner generates a random owner name
func RandOwner() string {
	return RandString(6)
}

// RandMoney generates a random amount of money
func RandMoney() int64 {
	return RandInt(0, 1000)
}

// RandCurrency generates a random currency code
func RandCurrency() string {
	currencies := []string{"USD", "EUR", "CAD"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}
