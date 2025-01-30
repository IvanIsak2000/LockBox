package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct{
	ApiKey string
}

func GetConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil{
		return nil, fmt.Errorf("Не удалось загрузить конфиг: %e", err)
	}
	apiKey := os.Getenv("API_KEY")
	if apiKey == ""{
		return nil, fmt.Errorf("API_KEY не обнаружен в файле .env. Пожалуйста, получите его на сервере и вставьте в файл")
	}
	
	return &Config{ApiKey: apiKey}, nil
}