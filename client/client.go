package main

import (
	"client/config"
	"client/connect"
	"client/cryptography"
	"client/ping"
	"flag"
	"fmt"
	"os"
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
	loginFlag := flag.String("login", "", "Авторизироваться в системе и предьявить MasterKey")
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
	
	if *loginFlag != ""{
		suggessMasterKey := cryptography.ArgonMasterKey([]byte(*loginFlag))
		fmt.Printf("suggessMasterKey: %v\n", suggessMasterKey)
		
		fileData, err := os.ReadFile("secret.key")
		if err != nil {
			println("Файл secret.key отсутствует, поэтому проверка мастер ключа невозможна. Создаю...")
			newKey := cryptography.GenerateKey()
			encryptKey := cryptography.ArgonMasterKey(newKey)
			fmt.Printf("encryptKey: %v\n", encryptKey)
			keyStr := string(newKey)
			fmt.Printf("изнчальный ключ: %v\n", keyStr)
			cryptography.WriteEncryptKey(encryptKey)
			
		}
		fmt.Printf("fileData: %v\n", string(fileData))
		// suggessMasterKeyHex := hex.EncodeToString([]byte(suggessMasterKey))
		// fmt.Printf("suggessMasterKeyHex: %v\n", suggessMasterKeyHex)
		if string(fileData) ==  suggessMasterKey{
			fmt.Println("Пароль верный!")
		} else {
			fmt.Println("Пароль неверный!")
		}
		// var MasterKey string 
		// MasterKey = *loginFlag
		// если пользователь делает login, то в эту переменную
		// сохраняет исходынй ключ, если конечно он прошёл аргонирование и проверку на хэш
		
	}
}