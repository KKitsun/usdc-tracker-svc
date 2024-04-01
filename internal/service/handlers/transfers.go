package handlers

import (
	"net/http"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"

	"github.com/KKitsun/usdc-tracker-svc/internal/service/models"
	"github.com/KKitsun/usdc-tracker-svc/internal/service/requests"
)

func GetTransfer(w http.ResponseWriter, r *http.Request) {
	db := DB(r)

	// requestFilter, err := requests.NewGetFilter(r)
	// if err != nil {
	// 	ape.RenderErr(w, problems.BadRequest(err)...)
	// 	return
	// }

	from := requests.NewGetSender(r)
	to := requests.NewGetReceiver(r)
	counterparty := requests.NewGetCounterparty(r)

	query := db.Transfer()
	if from != nil {
		query = query.FilterBySender(from...)
	}
	if to != nil {
		query = query.FilterByReceiver(to...)
	}
	if counterparty != nil {
		query = query.FilterByCounterparty(counterparty...)
	}

	transfer, err := query.Select()
	if err != nil {
		Log(r).WithError(err).Error("failed to query db")
		ape.RenderErr(w, problems.InternalError())
		return
	}
	ape.Render(w, models.NewTransfersListModel(transfer, "", "", "", "", "http://"+r.Host+r.URL.Path+"?"+r.URL.RawQuery))

}
