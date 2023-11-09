package utils

import (
	"math/rand"
	"time"
)

func RandomString(length int) string {
	var domain = []string{"@gmail.com"}
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	deref := rand.NewSource(time.Now().UnixNano())
	randSrc := rand.New(deref)
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[randSrc.Intn(len(charset))]
	}
	return string(result) + domain[randSrc.Intn(len(domain))]
}

func RandomPassword() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	firstChar := charset[rand.Intn(26)+26]
	password := make([]byte, 7)
	for i := 0; i < 7; i++ {
		password[i] = charset[rand.Intn(len(charset))]
	}
	password[rand.Intn(7)] = charset[rand.Intn(10)+52]
	return string(firstChar) + string(password)
}
