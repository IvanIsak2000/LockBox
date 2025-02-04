package main

import (
	"client/color"
	"client/config"
	"client/connect"
	"client/cryptography"
	"client/ping"
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/helioloureiro/golorama"
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
				color.Print("Сервер доступен", golorama.GREEN)
			case 404:
				color.Print("Эндпоинт не найден, попробуйте URL_СЕРВЕРА без лишних `/`", golorama.RED)
			default:
				color.Print(
					fmt.Sprintf("Неопознаная ошибка: %v", strconv.Itoa(int(resp))),
					golorama.RED,
				)
		}
	}
	
	if *connectFlag != ""{
		
		config, err := config.GetConfig()
		if err != nil {
			color.Print(
				"Для подключения сначала установите API_KEY в файле .env", 
				golorama.RED)
		}
		
		resp := connect.SendConnectionRequest(*connectFlag + ConnectUrl, config.ApiKey)
		switch resp {
			case 200:
				color.Print("Авторизация успешно пройдена!", golorama.GREEN)
			case 400:
				color.Print("Требуется параметр api_key", golorama.RED)
			case 401:
			color.Print("Неверный api_key" , golorama.RED)
			case 404:
			color.Print("Эндпоинт не найден, попробуйте URL_СЕРВЕРА без лишних `/`", golorama.RED)
			default:
			color.Print(
				fmt.Sprintf("Неопознаная ошибка: %v", strconv.Itoa(int(resp))),
				golorama.RED,
			)		}
	} 
	
	if *loginFlag != ""{
		suggessMasterKey := cryptography.ArgonMasterKey([]byte(*loginFlag))
		fmt.Printf("suggessMasterKey: %v\n", suggessMasterKey)
		
		fileData, err := os.ReadFile("secret.key")
		if err != nil {
			color.Print(
				"Файл secret.key отсутствует, поэтому проверка мастер ключа пока невозможна. Создаю...",
				golorama.RED,
			)
			newKey := cryptography.GenerateKey()
			encryptKey := cryptography.ArgonMasterKey(newKey)
			fmt.Printf("encryptKey: %v\n", encryptKey)
			cryptography.WriteEncryptKey(encryptKey)

			keyStr := string(newKey)
			color.Print(
				fmt.Sprintf(`
ВАШ MASTER KEY УСПЕШНО СОЗДАН:
-----------------------------
%v
-----------------------------
СОХРАНИТЕ ЕГО В НАДЁЖНОМ МЕСТЕ. ПРОГРАММА
НЕ ЗНАЕТ ЕГО, ПОЭТОМУ ВЫ ЕГО БОЛЬШЕ НИКОГДА НЕ УВИДИТЕ.
ПРИ ЕГО УТРАТЕ ВЫ НЕ СМОЖЕТЕ ВОССТАНОВИТЬ ВАШИ ДАННЫЕ

`, keyStr),
				golorama.GREEN,
			)
			cryptography.WriteMasterKey(keyStr)
			
		} else {
			fmt.Printf("fileData: %v\n", string(fileData))
			if string(fileData) ==  suggessMasterKey{
				// если пользователь делает login, то в эту переменную
				// сохраняет исходынй ключ, если конечно он прошёл аргонирование и проверку на хэш
				var MasterKey string 
				MasterKey = *loginFlag
				
				color.Print("Пароль верный!", golorama.GREEN)
				fmt.Printf("MasterKey: %v\n", MasterKey)
			} else {
				color.Print("Пароль неверный!", golorama.RED)
			}
		}
		
	}
}