package utility

import (
	"math/rand"
	"strconv"
	"time"
)

func UID(length int) string {
	if length <= 0 {
		panic("UUID length must be greater than 0")
	}

	maxValue := intPow(10, length) - 1
	uid := strconv.Itoa(rand.Intn(maxValue + 1))
	ustr := generateRandomString(10)
	return ustr + uid

}

func intPow(base, exponent int) int {
	result := 1
	for i := 0; i < exponent; i++ {
		result *= base
	}
	return result
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func generateRandomString(length int) string {
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(result)
}
