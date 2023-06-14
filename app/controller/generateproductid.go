package controller

import (
	"math/rand"
	"strconv"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func GenerateRandomProductID() string {
	rand.Seed(time.Now().UnixNano())

	// Generate angka acak antara 0 dan 9999
	randomNumber := rand.Intn(9999)

	// Generate karakter acak
	k := make([]byte, 4)
	for i := range k {
		index := rand.Intn(len(charset))
		k[i] = charset[index]
	}

	return strconv.Itoa(randomNumber) + string(k)
}
