package requests

import (
	"net/http"

	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/urlval"
)

type FilterRequest struct {
	From         string `url:"from"`
	To           string `url:"to"`
	Counterparty string `url:"counterparty"`
}

func NewGetFilter(r *http.Request) (*FilterRequest, error) {
	var request FilterRequest
	if err := urlval.Decode(r.URL.Query(), &request); err != nil {
		return nil, errors.Wrap(err, "failed decode query")
	}
	return &request, nil
}
