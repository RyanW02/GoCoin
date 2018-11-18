package utils

import (
	"math/rand"
)

var(
	letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

func Must(err error) {
	if err != nil {
		panic(err)
	}
}

func RandomString(length int) string {
	array := make([]rune, length)

	for i := range array {
		array[i] = letters[rand.Intn(len(letters))]
	}

	return string(array)
}
