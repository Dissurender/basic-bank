package util

import (
	"math/rand"
	"strings"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

// RandomInt generates a random integer with constraints min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString generates a random string of length n
func RandomString(n int) string {
	var strB strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		character := alphabet[rand.Intn(k)]
		strB.WriteByte(character)
	}

	return strB.String()
}

// RandomOwner generates a random name of length 8
func RandomOwner() string {
	return RandomString(8)
}

// RandomMoney generate a random amount of currency
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// RandomCurrency generates a random currency code
func RadomCurrency() string {
	curr := []string{"USD", "EUR", "CAD"}
	n := len(curr)

	return curr[rand.Intn(n)]
}
