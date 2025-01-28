package main

import (
	"client/config"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

func main() {
	pingFlag := flag.String("ping", "", "Ping a server.")
	// connectFlag := flag.String("connect", "http://localhost:8080/", "Connect to server")
	// unconnectlogin := flag.Bool("unconnect", false, "Unconnect from server")
	// syncFlag := flag.Bool("sync", false, "")
	
	flag.Parse()
	
	if *pingFlag != ""{
		config, err  := config.GetConfig()
		if err != nil {
			log.Fatal(err)
		}
		
		data := url.Values{}
		data.Add("api_key", config.ApiKey)
		fmt.Println("API Key:", config.ApiKey)

		
		response, err := http.PostForm(*pingFlag + "/ping", data)
		if err != nil {
			log.Fatal("Не удалось пингануть сервер: ", err)
		}
		fmt.Println("Ответ от сервера: ", response)
	}
	
	fmt.Println("args: ", flag.Args())
	
}