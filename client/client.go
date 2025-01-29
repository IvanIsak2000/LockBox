package main

import (
	"client/connect"
	"client/ping"
	"flag"
	"fmt"

	color "github.com/helioloureiro/golorama"
)

func main() {
	pingFlag := flag.String("ping", "", "Сделать пинг сервера")
	connectFlag := flag.String("connect", "", "Соединение с сервером")
	apiKeyFlag := flag.String("apiKey", "", "api_key параметр для подключения")
	// unconnectlogin := flag.Bool("unconnect", false, "Unconnect from server")
	// syncFlag := flag.Bool("sync", false, "")
	
	flag.Parse()
	
	if *pingFlag != ""{	
		resp := ping.PingServer(*pingFlag + "/")	
		switch resp{
			case 200:
				fmt.Println(color.GetCSI(color.GREEN) + "Сервер доступен" + color.Reset())
			case 404:
				fmt.Println(color.GetCSI(color.RED) + "Эндпоинт не найден, попробуйте URL_СЕРВЕРА без лишних `/`" + color.Reset())
			default:
				fmt.Println(color.GetCSI(color.RED) + "Неопознаная ошибка: ", string(resp) + color.Reset())
		}
	}
	
	if *connectFlag != ""{
		if *apiKeyFlag != "" {
			resp := connect.SendConnectionRequest(*connectFlag + "/connect", *apiKeyFlag)
			switch resp {
				case 200:
					fmt.Println(color.GetCSI(color.GREEN) + "Авторизация успешно пройдена!" + color.Reset())
				case 400:
					fmt.Println(color.GetCSI(color.RED) + "Требуется параметр api_key" + color.Reset())
				case 401:
					fmt.Println(color.GetCSI(color.RED) + "Неверный api_key" + color.Reset())
				default:
					fmt.Printf(color.GetCSI(color.RED) + "Неопознаная ошибка: ", string(resp) + color.Reset())
			}
		} else {
			fmt.Println(color.GetCSI(color.RED) + "Для подключения к серверу необходимо передать параметр apiKey: ... --apiKey <API_KEY>" + color.Reset())
		}
	}
}