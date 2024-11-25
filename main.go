package main

import (
	"math/rand"
	"strings"
	"time"
)

func randomPassword(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	randomSource := rand.NewSource(time.Now().UnixNano())
	random := rand.New(randomSource)

	var password strings.Builder
	for i := 0; i < length; i++ {
		randomIndex := random.Intn(len(charset))
		password.WriteByte(charset[randomIndex])
	}
	return password.String()
}

func main() {
	length := 12
	password := randomPassword(length)
	println("Generated Password:", password)
}
