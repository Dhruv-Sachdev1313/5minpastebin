package pkg

import (
	"math/rand"
	"time"
)

var randSeed = rand.New(rand.NewSource(time.Now().UnixNano()))

func RandomStringGenerator(length int) string {
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := range result {
		result[i] = chars[randSeed.Intn(len(chars))]
	}
	return string(result)
}
