package services

import (
	"math/rand"
	"time"
)

func GeneratePassword() string {
	rand.Seed(time.Now().UnixNano())
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZÅÄÖ" +
		"abcdefghijklmnopqrstuvwxyzåäö" +
		"0123456789" +
		"!@#$%^&*()-_=+,.?/:;{}[]`~")
	length := 20
	var b []rune
	for i := 0; i < length; i++ {
		b = append(b, chars[rand.Intn(len(chars))])
	}
	return string(b)
}
