package util

import (
	"fmt"
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
	for i := 0; i < n; i++ {
		randomStringBuilder.WriteByte(alphabet[rand.Intn(alphabetLength)])
	}
	return randomStringBuilder.String()
}

func RandomName() string {
	return RandomString(6)
}

func RandomBalance() int64 {
	return RandomInt(0, 1000)
}

func RandomCurrency() string {
	currencies := []string{EUR, USD, CAD}
	return currencies[rand.Intn(len(currencies))]
}

func RandomEmail() string {
	return fmt.Sprintf("%s@%s.com", RandomString(6), RandomString(6))
}
