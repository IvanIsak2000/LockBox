package main

import (
	"crypto/aes"
	"fmt"
	"log"
)

func Encrypt(data string) {
	dataBytes := []byte(data)
	cipher, err := aes.NewCipher(dataBytes)
	if err != nil {
		log.Fatal("Не удалось создать NewCipher: ", err)
	}
	fmt.Printf("cipher: %v\n", cipher)
}

func main() {
	data := "my_super_password"
	Encrypt(data)

}