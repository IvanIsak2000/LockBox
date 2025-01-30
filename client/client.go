package main

import (
	"client/config"
	"client/connect"
	"client/ping"
	"flag"
	"fmt"
	"strconv"

	color "github.com/helioloureiro/golorama"
)

const (
	PingUrl = "/"
	ConnectUrl = "/connect"
)

func main() {
	pingFlag := flag.String("ping", "", "Сделать пинг сервера")
	connectFlag := flag.String("connect", "", "Соединение с сервером")
	// unconnectlogin := flag.Bool("unconnect", false, "Unconnect from server")
	// syncFlag := flag.Bool("sync", false, "")
	
	flag.Parse()

	if *pingFlag != ""{	
		resp := ping.PingServer(*pingFlag + PingUrl)	
		switch resp{
			case 200:
				fmt.Println(color.GetCSI(color.GREEN) + "Сервер доступен" + color.Reset())
			case 404:
				fmt.Println(color.GetCSI(color.RED) + "Эндпоинт не найден, попробуйте URL_СЕРВЕРА без лишних `/`" + color.Reset())
			default:
				fmt.Println(color.GetCSI(color.RED) + "Неопознаная ошибка: ", strconv.Itoa(int(resp)) + color.Reset())
		}
	}
	
	if *connectFlag != ""{
		
		config, err := config.GetConfig()
		if err != nil {
			fmt.Println(color.GetCSI(color.RED) + "Для подключения сначала установите API_KEY в файле .env" + color.Reset())
		}
		
		resp := connect.SendConnectionRequest(*connectFlag + ConnectUrl, config.ApiKey)
		switch resp {
			case 200:
				fmt.Println(color.GetCSI(color.GREEN) + "Авторизация успешно пройдена!" + color.Reset())
			case 400:
				fmt.Println(color.GetCSI(color.RED) + "Требуется параметр api_key" + color.Reset())
			case 401:
				fmt.Println(color.GetCSI(color.RED) + "Неверный api_key" + color.Reset())
			case 404:
				fmt.Println(color.GetCSI(color.RED) + "Эндпоинт не найден, попробуйте URL_СЕРВЕРА без лишних `/`" + color.Reset())
			default:
				fmt.Printf(color.GetCSI(color.RED) + "Неопознаная ошибка: ", strconv.Itoa(int(resp)) + color.Reset())
		}
	} 
}