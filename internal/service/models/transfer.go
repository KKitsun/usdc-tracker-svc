package models

import (
	"github.com/KKitsun/usdc-tracker-svc/internal/storage"
	"github.com/KKitsun/usdc-tracker-svc/resources"
)

func NewTransferModel(t storage.Transfer) resources.TransferResponse {
	response := resources.TransferResponse{
		Data: resources.Transfer{
			Key: resources.Key{Type: resources.TRANSFER},
			Attributes: &resources.TransferAttributes{
				Txhash:       t.TxHash,
				FromAddress:  t.From,
				ToAddress:    t.To,
				ValueDecimal: t.Value,
			},
		},
	}
	return response
}

func TransfersList(ts []storage.Transfer) []resources.Transfer {
	var transfersList []resources.Transfer
	for _, t := range ts {
		transfersList = append(transfersList, resources.Transfer{
			Key: resources.NewKeyInt64(t.ID, resources.TRANSFER),
			Attributes: &resources.TransferAttributes{
				Txhash:       t.TxHash,
				FromAddress:  t.From,
				ToAddress:    t.To,
				ValueDecimal: t.Value,
			},
		})
	}
	return transfersList
}

func NewTransfersListModel(ts []storage.Transfer, firstLink, lastLink, nextLink, prevLink, selfLink string) resources.TransferListResponse {
	response := resources.TransferListResponse{
		Data:     TransfersList(ts),
		Included: resources.Included{},
		Links: &resources.Links{
			First: firstLink,
			Last:  lastLink,
			Next:  nextLink,
			Prev:  prevLink,
			Self:  selfLink,
		},
	}
	return response
}
