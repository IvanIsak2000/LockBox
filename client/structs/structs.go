package structs

import "github.com/meteran/gnext"


type ConnectResponse struct {
	Status gnext.Status `json:"status"`
	Message string `json:"message"`
}