package main

import (
	"crypto/aes"
	"encoding/hex"
	"fmt"
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

func EncodeMasterKey(key []byte) string {
	encryptMasterKey := argon2.Key(key, []byte(salt), 1, 32*1024, 4, 32)
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

func main() {
	// Создём мастер ключ и выдаём его хэш - то есть для сравнения будет использовать хэш от мастер ключа 
	key := GenerateKey()
	fmt.Printf("Ключ: %v\n", hex.EncodeToString(key))
	encryptKey := EncodeMasterKey(key)
	fmt.Printf("argoned key: %v\n", encryptKey)
	WriteEncryptKey(encryptKey)
	
	
	data := "hello_worlds123!"
	fmt.Printf("Данные: %v\n", data)
	
	// TODO: добавить расшифровку из argon2 мастер ключа и передавать уже его!
	encryptData := Encrypt(key, data)
	fmt.Printf("Зашифрованные данные: %v\n", encryptData)
	
	decryptData := Decrypt(key, encryptData)
	fmt.Printf("decryptData: %v\n", decryptData)
}