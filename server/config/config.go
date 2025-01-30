package config

import (
	"fmt"
	"log"
	"os"
	"server/color"

	"github.com/helioloureiro/golorama"
	"github.com/joho/godotenv"
)

const (
	PrivateKeyFile = ".env"
)

type Config struct {
	ApiKey string
}

func PrivateKeyIsExist() bool {
	_, err := os.ReadFile(PrivateKeyFile)
	if err != nil {
		fmt.Printf("Файл %s ещё не создан. Создаю...\n", PrivateKeyFile)
		return false
	}
	color.Print(fmt.Sprintf("Файл %s уже существует.", PrivateKeyFile), golorama.GREEN)
	return true
}


func DoPrivateKeyFile(apiKey string)  {
	data := fmt.Sprintf("API_KEY: %s", apiKey)

	err := os.WriteFile(PrivateKeyFile, []byte(data), 0644)
	if err != nil {
		log.Fatal("Не удалось записать новый ключ: ", err)
	}
	color.Print("Файл .env успешно создан. Значение успешно записано", golorama.GREEN)
	
}

func GetConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Не удалось прочитать .env: ", err)
	}
	key := os.Getenv("API_KEY")
	return &Config{ApiKey: key}
} 