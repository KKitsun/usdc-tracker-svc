package postgres

import (
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/kit/pgdb"

	"github.com/KKitsun/usdc-tracker-svc/internal/storage"
	_ "gitlab.com/distributed_lab/logan/v3/errors"
)

const transferTableName = "transfer"

func newTransferQ(db *pgdb.DB) storage.TransferQ {
	return &transferQ{
		db:  db,
		sql: sq.StatementBuilder,
	}
}

type transferQ struct {
	db  *pgdb.DB
	sql sq.StatementBuilderType
}

func (q *transferQ) InsertTransfer(value storage.Transfer) (*storage.Transfer, error) {
	var result storage.Transfer
	stmt := sq.Insert(transferTableName).
		Columns("txhash", "from_address", "to_address", "value_decimal").
		Values(value.TxHash, value.From, value.To, value.Value).
		Suffix("returning id")
	err := q.db.Get(&result, stmt)
	if err != nil {
		return nil, errors.Wrap(err, "failed to insert transfer to db")
	}
	return &result, nil
}

func (q *transferQ) Get() (*storage.Transfer, error) {
	var result storage.Transfer
	err := q.db.Get(&result, q.sql.Select("*").From(transferTableName))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, errors.Wrap(err, "failed to get transfer from db")
	}
	return &result, nil
}

func (q *transferQ) Select() ([]storage.Transfer, error) {
	var result []storage.Transfer
	err := q.db.Select(&result, q.sql.Select("*").From(transferTableName))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, errors.Wrap(err, "failed to select transfers from db")
	}
	return result, nil
}

func (q *transferQ) FilterBySender(from_address ...string) storage.TransferQ {
	pred := sq.Eq{"from_address": from_address}
	q.sql = q.sql.Where(pred)
	return q
}

func (q *transferQ) FilterByReceiver(to_address ...string) storage.TransferQ {
	pred := sq.Eq{"to_address": to_address}
	q.sql = q.sql.Where(pred)
	return q
}

func (q *transferQ) FilterByCounterparty(address ...string) storage.TransferQ {
	pred1 := sq.Eq{"from_address": address}
	pred2 := sq.Eq{"to_address": address}
	pred := sq.Or{pred1, pred2}
	q.sql = q.sql.Where(pred)
	return q
}
