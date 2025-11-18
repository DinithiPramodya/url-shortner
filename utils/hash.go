package utils

import (
	"crypto/sha1"
	"math/rand"
	"strings"
	"time"
)

var chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func GenerateShortCode(n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, n)
	for i := range b {
		b[i] = chars[rand.Intn(len(chars))]
	}
	return string(b)
}

func GenerateShortCodeWithSalt(url string) string {
	salt := time.Now().String()
	hash := sha1.Sum([]byte(url + salt))

	var code strings.Builder
	for _, b := range hash[:6] {
		code.WriteByte(chars[int(b)%len(chars)])
	}
	return code.String()
}
