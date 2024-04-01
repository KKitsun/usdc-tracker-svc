package requests

import (
	"net/http"
)

func NewGetCounterparty(r *http.Request) []string {
	request := r.URL.Query()["counterparty"]

	return request
}
