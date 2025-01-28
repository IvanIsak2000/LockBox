package main

import (
	"fmt"
	"server/key"
	"server/structs"
	"github.com/meteran/gnext"
)

type Server struct{
	ApiKey string
}

// Запуск сервера и обработка http запросов
func main() {
	apikey := key.GenerateApiKey()
	server := &Server{ApiKey: apikey}
	
	fmt.Printf("\nApikey: %v\n", apikey)
	
	router := gnext.Router()
	router.POST("/ping", server.Ping)
	_ = router.Run()
}
	
func (s *Server) Ping(request *structs.ConnectRequest) *structs.ConnectResponse {
	fmt.Printf("\nRequest: %v\n", request)
	if request == nil{
		return &structs.ConnectResponse{Status: 400, Message: "Request need contains api_key"}
	}
	
	if request.ApiKey != s.ApiKey{
		return &structs.ConnectResponse{Status: 401, Message: "Incorrect ApiKey"}
	}
	
	return &structs.ConnectResponse{Status: 200, Message: "Welcome!"}
	
}