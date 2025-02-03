package cryptography

import (
	"crypto/aes"
	"encoding/hex"
	"log"
	"math/rand"
	"os"

	"golang.org/x/crypto/argon2"
)

const (
	blockSize = 16
	symbols = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	secretKeyFile = "secret.key"
	salt = "апапапап"
)

func GenerateKey() []byte {
	var key []byte
	
	for i := 0; i < blockSize; i++ {
		index := rand.Intn(len(symbols))
		key = append(key, symbols[index])
		
	}
	return key
}


// Шифровка данных
func Encrypt(key []byte, data string) string {
	dataBytes := []byte(data)
	result := make([]byte, len(dataBytes))
	
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Fatal("Не удалось создать шифровальщик ", err)
	}
	block.Encrypt(result, dataBytes)
	return hex.EncodeToString(result)
	
}

// Дешифрока данных
func Decrypt(key []byte, data string) string {
	dataBytes, err := hex.DecodeString(data)
	if err != nil {
		log.Fatal("Не удалось декодировать данные: ", err)
	}
	
	Result := make([]byte, len(dataBytes))
	
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Fatal("Не удалось создать шифровальщик ", err)
	}
	block.Decrypt(Result, dataBytes)
	return string(Result)
}

func ArgonMasterKey(key []byte) string {
	encryptMasterKey := argon2.Key(key, []byte(salt), 1, 32*1024, 2, 32)
	return hex.EncodeToString(encryptMasterKey)
}

func WriteEncryptKey(key string) {
	os.Remove(secretKeyFile)
	file, err := os.Create(secretKeyFile)
	if err != nil {
		log.Fatalf("Не удалось создать файл %v: %v", secretKeyFile, err)
	}
	defer file.Close()
	file.WriteString(key)
}
