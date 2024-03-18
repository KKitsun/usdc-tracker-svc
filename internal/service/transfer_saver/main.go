package transfer_saver

import (
	"context"
	"encoding/hex"
	"log"
	"math/big"

	"github.com/KKitsun/usdc-tracker-svc/internal/storage"
	"github.com/pkg/errors"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	usdcContractAddr   = "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"
	logTransferSigHash = "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"
)

func EventsListener(client *ethclient.Client, tStorage storage.TransferStorage) {
	// logTransferSig := []byte("Transfer(address,address,uint256)")
	// logTransferSigHash := crypto.Keccak256Hash(logTransferSig)

	contractAddress := common.HexToAddress(usdcContractAddr)
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			if vLog.Topics[0].Hex() == logTransferSigHash {
				if len(vLog.Topics) != 3 {
					continue
				}
				fromAddress := "0x" + string(vLog.Topics[1].Hex())[26:]
				toAddress := "0x" + string(vLog.Topics[2].Hex())[26:]
				hexInt := new(big.Int)
				hexInt.SetString(hex.EncodeToString(vLog.Data), 16)
				decimals := big.NewInt(1000000)
				decimalValue := new(big.Float).Quo(new(big.Float).SetInt(hexInt), new(big.Float).SetInt(decimals))

				_, err := tStorage.Transfer().InsertTransfer(storage.Transfer{
					TxHash: vLog.TxHash.String(),
					From:   fromAddress,
					To:     toAddress,
					Value:  decimalValue.String(),
				})
				if err != nil {
					panic(errors.Wrap(err, "failed to insert transfer"))
				}
			}
		}
	}

}
