package main

import (
	"crypto/aes"
	"encoding/hex"
	"fmt"
	"log"
)

func Encrypt(key, data string) string {
	keyBytes := []byte(key)
	dataBytes := []byte(data)
	Result := make([]byte, len(dataBytes))
	
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		log.Fatal("Не удалось создать NewCipher: ", err)
	}
	block.Encrypt(Result, dataBytes)
	return hex.EncodeToString(Result)
	
}

func main() {
	key := "my_super_passwor"
	data := "мой_секретная_ин"
	encryptData := Encrypt(key, data)
	fmt.Printf("encryptData: %v\n", encryptData)
}