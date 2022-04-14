package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {
	var randomStringBuilder strings.Builder
	alphabetLength := len(alphabet)
	for i := 0; i < len(alphabet); i++ {
		randomStringBuilder.WriteByte(alphabet[rand.Intn(alphabetLength)])
	}
	return randomStringBuilder.String()
}

func RandomOwner() string {
	return RandomString(6)
}

func RandomBalance() int64 {
	return RandomInt(0, 1000)
}

func RandomCurrency() string {
	currencies := []string{"USD", "EUR", "CAD"}
	return currencies[rand.Intn(len(currencies))]
}
