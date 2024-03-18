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

	sender := requests.NewGetSender(r)
	receiver := requests.NewGetReceiver(r)

	if sender != nil && receiver != nil {
		transfer, err := db.Transfer().FilterBySender(sender...).FilterByReceiver(receiver...).Select()
		if err != nil {
			Log(r).WithError(err).Error("failed to query db")
			ape.RenderErr(w, problems.InternalError())
			return
		}

		ape.Render(w, models.NewTransfersListModel(transfer, "", "", "", "", "http://"+r.Host+r.URL.Path+"?"+r.URL.RawQuery))
		return

	} else if sender != nil {
		transfer, err := db.Transfer().FilterBySender(sender...).Select()
		if err != nil {
			Log(r).WithError(err).Error("failed to query db")
			ape.RenderErr(w, problems.InternalError())
			return
		}

		ape.Render(w, models.NewTransfersListModel(transfer, "", "", "", "", "http://"+r.Host+r.URL.Path+"?"+r.URL.RawQuery))
		return

	} else if receiver != nil {
		transfer, err := db.Transfer().FilterByReceiver(receiver...).Select()
		if err != nil {
			Log(r).WithError(err).Error("failed to query db")
			ape.RenderErr(w, problems.InternalError())
			return
		}

		ape.Render(w, models.NewTransfersListModel(transfer, "", "", "", "", "http://"+r.Host+r.URL.Path+"?"+r.URL.RawQuery))
		return

	} else {
		transfer, err := db.Transfer().Select()
		if err != nil {
			Log(r).WithError(err).Error("failed to query db")
			ape.RenderErr(w, problems.InternalError())
			return
		}

		ape.Render(w, models.NewTransfersListModel(transfer, "", "", "", "", "http://"+r.Host+r.URL.Path))
		return
	}

}
