package requests

import (
	"net/http"
)

func NewGetSender(r *http.Request) []string {
	//var request string
	request := r.URL.Query()["from"]

	return request
}
