package requests

import (
	"net/http"
)

func NewGetReceiver(r *http.Request) []string {
	//var request string
	request := r.URL.Query()["to"]

	return request
}
