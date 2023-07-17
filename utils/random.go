package utils

import (
	"math/rand"
	"time"
)

var bytes = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890")

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandStr(n int) string {
	result := make([]byte, n)
	for i := 0; i < n; i++ {
		result[i] = bytes[rand.Int31()%62]
	}
	return string(result)
}
