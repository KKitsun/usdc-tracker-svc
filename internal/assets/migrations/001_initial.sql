-- +migrate Up

CREATE TABLE transfer
(
    id BIGSERIAL PRIMARY KEY,
	txhash BYTEA NOT NULL,
	from_address BYTEA NOT NULL,
    to_address BYTEA NOT NULL,
    value_decimal DECIMAL NOT NULL
);
CREATE INDEX transfer_from ON transfer(from_address);
CREATE INDEX transfer_to ON transfer(to_address);

-- +migrate Down
DROP TABLE transfer;