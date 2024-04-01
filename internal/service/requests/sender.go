package requests

import (
	"net/http"
)

func NewGetSender(r *http.Request) []string {
	request := r.URL.Query()["from"]

	return request
}
