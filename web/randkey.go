package web

import (
	"math/rand"
)

const (
	validChars = "abcdefghijklmnoprstuwxyzABCDEFGHIJKLMNOPRSTUWXYZ0123456789"
)

func genKey(length int) string {
	chars := make([]byte, 0, length)
	for i := 0; i < length; i++ {
		c := validChars[rand.Intn(len(validChars))]
		chars = append(chars, c)
	}
	return string(chars)
}
