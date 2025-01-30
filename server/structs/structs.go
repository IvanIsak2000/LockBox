package structs

import "github.com/meteran/gnext"

type PingRequest struct {}

type EmptyRequest struct {}

type PingResponse struct {
	Status gnext.Status
	Message string
}

type ConnectRequest struct{
	ApiKey string `json:"api_key"`
}

type ConnectResponse struct {
	Status gnext.Status `json:"status"`
	Message string `json:"message"`
}
