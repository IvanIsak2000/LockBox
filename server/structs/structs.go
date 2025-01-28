package structs

import "github.com/meteran/gnext"

type ConnectRequest struct {
	ApiKey string `json:"api_key"`
}


type ConnectResponse struct {
	Status gnext.Status
	Message string
}