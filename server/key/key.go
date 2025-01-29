package key

import (
	"math/rand"
)


// Генерирует случайный apiKey
func GenerateApiKey() string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	const digits = "0123456789"
	const symbols = letters + digits
	
	var apiKeyLen = 100
	var apiKey string
	
	for i:=0; i < apiKeyLen; i++{
		randPos := rand.Intn(len(symbols))
		apiKey += string(symbols[randPos])
	}
	return apiKey
}