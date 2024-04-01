package service

import (
	"net"
	"net/http"
	"sync"

	"github.com/KKitsun/usdc-tracker-svc/internal/config"
	"github.com/KKitsun/usdc-tracker-svc/internal/service/transfer_saver"
	"github.com/KKitsun/usdc-tracker-svc/internal/storage/postgres"

	"gitlab.com/distributed_lab/kit/copus/types"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type service struct {
	log      *logan.Entry
	copus    types.Copus
	listener net.Listener
}

func (s *service) run(cfg config.Config) error {
	s.log.Info("Service started")
	r := s.router(cfg)

	if err := s.copus.RegisterChi(r); err != nil {
		return errors.Wrap(err, "cop failed")
	}

	return http.Serve(s.listener, r)
}

func runTransferSaver(cfg config.Config, wg *sync.WaitGroup) {
	client, contract := cfg.EthereumConnection()
	wg.Add(1)
	go func() {
		defer func() {
			wg.Done()
		}()
		transfer_saver.EventsListener(client, contract, postgres.NewTransferStorage(cfg.DB()))
	}()
}

func newService(cfg config.Config) *service {
	return &service{
		log:      cfg.Log(),
		copus:    cfg.Copus(),
		listener: cfg.Listener(),
	}
}

func Run(cfg config.Config) {
	wg := &sync.WaitGroup{}

	runTransferSaver(cfg, wg)

	if err := newService(cfg).run(cfg); err != nil {
		panic(err)
	}
}
