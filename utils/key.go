package utils

import (
	"crypto/rand"
	"math/big"
)

const KeyLength = 6

// Define the character set
var characters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func GenerateKey(length int) (string, error) {
	key := make([]rune, length)
	for i := range key {
		index, err := rand.Int(rand.Reader, big.NewInt(int64(len(characters))))
		if err != nil {
			return "", err
		}
		key[i] = characters[index.Int64()]
	}
	return string(key), nil
}

func GetShortUrl(key string) string {
	return "http://localhost:8080/" + key
}
