package connect

import (
	"bytes"
	"client/structs"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/meteran/gnext"
)

func SendConnectionRequest(url string, apiKey string) gnext.Status {
	data := map[string]string{"api_key": apiKey}
	
	marshData, err := json.Marshal(data)
	if err != nil {
		log.Fatal("Не удалось маршализовать данные: ", err)
	}
	
	readData := bytes.NewReader(marshData)
	request, err := http.NewRequest("POST", url, readData)
	if err != nil {
		log.Fatal("Не удалось создать запрос: ", err)
	}
	request.Header.Set("Content-Type", "application/json")
	
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Fatal("Не удалось отправить запрос: ", err)
	}
	
	defer response.Body.Close()
	
	if response.StatusCode == http.StatusOK{
		bodyBytes, err := io.ReadAll(response.Body)
		if err != nil {
			log.Fatal("Не удалось прочитать тело ответа: ", err)
		}
		fmt.Println("Тело ответа: ", string(bodyBytes))
		
		connectResponse := &structs.ConnectResponse{}
		err = json.Unmarshal(bodyBytes, connectResponse)
		if err != nil {
			log.Fatal("Не удалось распаковать байты в структуру: ", err)
		}
		
		return connectResponse.Status
	} 
	return gnext.Status(response.StatusCode)

}
