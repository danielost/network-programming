package service

import (
	"net"
	"net/http"

	"pvms-final/internal/config"
	"pvms-final/internal/service/requests"

	"gitlab.com/distributed_lab/kit/copus/types"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type service struct {
	log      *logan.Entry
	copus    types.Copus
	listener net.Listener

	queue chan requests.MoveRequest
}

func (s *service) run() error {
	s.log.Info("Service started")
	r := s.router()

	// Запускаємо горутину
	TicTacToeRunner(s.log, s.queue)

	if err := s.copus.RegisterChi(r); err != nil {
		return errors.Wrap(err, "cop failed")
	}

	// go func() {
	// 	for {
	// 		time.Sleep(time.Second * 5)
	// 		for i := 0; i < 5; i++ {
	// 			<-handlers.HttpQueue
	// 		}
	// 	}
	// }()

	return http.Serve(s.listener, r)
}

func newService(cfg config.Config) *service {
	return &service{
		log:      cfg.Log(),
		copus:    cfg.Copus(),
		listener: cfg.Listener(),

		queue: make(chan requests.MoveRequest, 50),
	}
}

func Run(cfg config.Config) {
	if err := newService(cfg).run(); err != nil {
		panic(err)
	}
}
