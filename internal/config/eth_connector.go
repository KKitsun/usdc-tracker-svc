package config

import (
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/pkg/errors"
	"gitlab.com/distributed_lab/figure/v3"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
)

type EthereumConnecter interface {
	EthereumConnection() *ethclient.Client
}

type ethereumConnecter struct {
	getter kv.Getter
	once   comfig.Once
}

func NewEthereumConnecter(getter kv.Getter) EthereumConnecter {
	return &ethereumConnecter{
		getter: getter,
	}
}

type ethereumConnecterCfg struct {
	URL string `fig:"url,required"`
}

func (ec *ethereumConnecter) readConfig() ethereumConnecterCfg {
	config := ethereumConnecterCfg{}
	err := figure.Out(&config).
		From(kv.MustGetStringMap(ec.getter, "ethereum_connection")).
		Please()
	if err != nil {
		panic(errors.Wrap(err, "failed to figure out"))
	}

	return config
}

func (ec *ethereumConnecter) EthereumConnection() *ethclient.Client {
	return ec.once.Do(func() interface{} {
		config := ec.readConfig()

		ethClient, err := ethclient.Dial(config.URL)
		if err != nil {
			panic(errors.Wrap(err, "failed to connect to the ethereum mainnet"))
		}

		return ethClient
	}).(*ethclient.Client)
}
