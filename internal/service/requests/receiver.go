package requests

import (
	"net/http"
)

func NewGetReceiver(r *http.Request) []string {
	request := r.URL.Query()["to"]

	return request
}
