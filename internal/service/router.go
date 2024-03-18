package service

import (
	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape"

	"github.com/KKitsun/usdc-tracker-svc/internal/config"
	"github.com/KKitsun/usdc-tracker-svc/internal/service/handlers"
	"github.com/KKitsun/usdc-tracker-svc/internal/storage/postgres"
)

func (s *service) router(cfg config.Config) chi.Router {
	r := chi.NewRouter()

	r.Use(
		ape.RecoverMiddleware(s.log),
		ape.LoganMiddleware(s.log),
		ape.CtxMiddleware(
			handlers.CtxLog(s.log),
			handlers.CtxDB(postgres.NewTransferStorage(cfg.DB())),
		),
	)
	r.Route("/integrations/usdc-tracker-svc", func(r chi.Router) {
		r.Get("/transfers", handlers.GetTransfer)
	})

	return r
}
