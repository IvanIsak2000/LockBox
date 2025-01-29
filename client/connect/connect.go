package connect

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/meteran/gnext"
)

func SendConnectionRequest(url string, apiKey string) gnext.Status {
	// Создаем структуру данных с API ключом
	data := struct {
		APIKey string `json:"api_key"`
	}{
		APIKey: apiKey,
	}

	// Маршализуем структуру в JSON
	marshData, err := json.Marshal(data)
	if err != nil {
		log.Fatal("Не удалось маршализовать данные для connect: ", err)
	}

	// Создаем тело запроса
	body := bytes.NewReader(marshData)

	// Создаем новый HTTP запрос
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		log.Fatal("Не удалось создать запрос: ", err)
	}

	// Отправляем запрос
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Не удалось отправить запрос: ", err)
	}
	defer resp.Body.Close()

	// Читаем тело ответа
	var response struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}

	// Парсим тело ответа в структуру
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		log.Fatal("Не удалось распарсить ответ: ", err)
	}

	// Возвращаем статус из тела ответа
	return gnext.Status(response.Status)
}
