package util

import (
	"strings"

	"golang.org/x/exp/rand"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

// RandomInt generates a random integer between min and max
func RandomInt(min int64, max int64) int64 {
	return min + rand.Int63n(max - min + 1)
}

// RandomString generates a random string with length n
func RandomString(n int) string {
	var sb strings.Builder
	alphabetLength := len(alphabet)

	for i:=0; i<n; i++ {
		c := alphabet[rand.Intn(alphabetLength)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomOwner generates a random owner name
func RandomOwner() string {
	return RandomString(6)
}

// RandomMoney generates a random amount of money
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// RandomCurrency generates a random currency from "EUR", "USD", or "CAD"
func RandomCurrency() string {
	currency := []string{"EUR", "USD", "CAD"}
	return currency[rand.Intn(len(currency) - 1)]
}