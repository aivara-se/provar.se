package random

import (
	"math/rand"
)

const (
	chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

// String generates a new base62 encoded random string of given length
func String(length int) string {
	id := make([]byte, length)
	for i := 0; i < length; i++ {
		id[i] = chars[rand.Intn(len(chars))]
	}
	return string(id)
}

// WithLength returns a function that generates a random with given length
func WithLength(length int) func() string {
	return func() string {
		return String(length)
	}
}
