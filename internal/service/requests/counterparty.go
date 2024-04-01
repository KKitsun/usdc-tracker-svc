package requests

import (
	"net/http"
)

func NewGetCounterparty(r *http.Request) []string {
	//var request string
	request := r.URL.Query()["counterparty"]

	return request
}
