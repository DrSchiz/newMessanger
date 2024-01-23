package math

import "math/rand"

func GenerateCode() string {
	var letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

	code := make([]byte, 6)
	for i := range code {
		code[i] = letters[rand.Intn(len(letters))]
	}
	return string(code)
}
