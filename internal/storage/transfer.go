package storage

type TransferQ interface {
	InsertTransfer(Transfer) (*Transfer, error)

	Get() (*Transfer, error)
	Select() ([]Transfer, error)

	FilterBySender(from_address ...string) TransferQ
	FilterByReceiver(to_address ...string) TransferQ
}

type Transfer struct {
	ID     int64  `db:"id"`
	TxHash string `db:"txhash"`
	From   string `db:"from_address"`
	To     string `db:"to_address"`
	Value  string `db:"value_decimal"`
}
