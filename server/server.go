package main

import (
	"fmt"
	"server/key"
	"server/structs"

	"github.com/gin-gonic/gin"
	"github.com/helioloureiro/golorama"
	"github.com/meteran/gnext"
)

type Server struct{
	ApiKey string
}

// Запуск сервера и обработка http запросов
func main() {
	apikey := key.GenerateApiKey()
	server := &Server{ApiKey: apikey}
	gin.SetMode(gin.ReleaseMode)
	
	fmt.Printf(golorama.GetCSI(golorama.RED) + "\nAPIKEY: %v\n", apikey + golorama.Reset())
	
	
	router := gnext.Router()
	router.GET("/", Ping)
	router.POST("/connect", server.Connection)
	
	_ = router.Run()
}
	
func Ping(request *structs.PingRequest) *structs.PingResponse {
	return &structs.PingResponse{Status: 200, Message: "Welcome!"}
	
}


func (s *Server) Connection(request *structs.ConnectRequest) *structs.ConnectResponse{
	if request == nil{
		return &structs.ConnectResponse{Status: 400, Message: "Требуется параметр api_key"}
	}
	
	if request.ApiKey == s.ApiKey{
		return &structs.ConnectResponse{Status: 200, Message: "Welcome!"}
	}
	return &structs.ConnectResponse{Status: 401, Message: "Неверный api_key"}
	
}