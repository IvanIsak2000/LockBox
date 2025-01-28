package structs

type ConnectRequest struct {
	ApiKey string `json:"api_key"`
}


type ConnectResponse struct {
	Success bool
	Message string
}