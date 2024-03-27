package service

import (
	"pvms-final/internal/service/handlers"

	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape"
)

func (s *service) router() chi.Router {
	r := chi.NewRouter()

	r.Use(
		ape.RecoverMiddleware(s.log),
		ape.LoganMiddleware(s.log),
		ape.CtxMiddleware(
			handlers.CtxLog(s.log),
			handlers.CtxBoard(&[100][100]string{}),
			handlers.CtxQueue(s.queue),
		),
	)

	r.Route("/integrations/pvms-final", func(r chi.Router) {
		r.Post("/move", handlers.Move)
	})

	return r
}
