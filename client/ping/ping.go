package ping

import (
	"log"
	"net/http"

	"github.com/meteran/gnext"
)

func PingServer(url string) gnext.Status{
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("Не удалось пингануть сервер: ", err)
	}
	return gnext.Status(resp.StatusCode)

}