package postgres

import (
	"github.com/KKitsun/usdc-tracker-svc/internal/storage"
	"gitlab.com/distributed_lab/kit/pgdb"
)

func NewTransferStorage(db *pgdb.DB) storage.TransferStorage {
	return &transferStorage{
		db: db.Clone(),
	}
}

type transferStorage struct {
	db *pgdb.DB
}

func (ts *transferStorage) New() storage.TransferStorage {
	return NewTransferStorage(ts.db)
}

func (ts *transferStorage) Transfer() storage.TransferQ {
	return newTransferQ(ts.db)
}
