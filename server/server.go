package main

import (
	"fmt"
	"server/color"
	"server/config"
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
	if !config.PrivateKeyIsExist(){
		newApikey := key.GenerateApiKey()
		config.DoPrivateKeyFile(newApikey)
		
	}
	config := config.GetConfig()
	server := &Server{ApiKey: config.ApiKey}
	color.Print(fmt.Sprint("API_KEY: ", config.ApiKey), golorama.RED)
	
	gin.SetMode(gin.ReleaseMode)
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