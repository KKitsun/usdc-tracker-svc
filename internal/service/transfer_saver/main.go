package transfer_saver

import (
	"log"
	"math/big"

	"github.com/KKitsun/usdc-tracker-svc/internal/service/USDC_contract"
	"github.com/KKitsun/usdc-tracker-svc/internal/storage"
	"github.com/pkg/errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	UsdcDecimalPoint = 1000000
)

func EventsListener(client *ethclient.Client, usdcContractAddr string, tStorage storage.TransferStorage) {
	contractAddress := common.HexToAddress(usdcContractAddr)

	eventFilterer, err := USDC_contract.NewUSDCFilterer(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	eventSink := make(chan *USDC_contract.USDCTransfer)

	sub, err := eventFilterer.WatchTransfer(nil, eventSink, nil, nil)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case event := <-eventSink:

			decimals := big.NewInt(UsdcDecimalPoint)
			decimalValue := new(big.Float).Quo(new(big.Float).SetInt(event.Value), new(big.Float).SetInt(decimals))

			_, err := tStorage.Transfer().InsertTransfer(storage.Transfer{
				TxHash: event.Raw.TxHash.String(),
				From:   event.From.String(),
				To:     event.To.String(),
				Value:  decimalValue.String(),
			})
			if err != nil {
				panic(errors.Wrap(err, "failed to insert transfer"))
			}

		}
	}

}
