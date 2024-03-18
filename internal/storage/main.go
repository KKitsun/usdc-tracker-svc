package storage

type TransferStorage interface {
	New() TransferStorage

	Transfer() TransferQ
}
