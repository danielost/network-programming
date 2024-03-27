package service

import (
	"bytes"
	"fmt"
	"math/rand"
	"net/http"
	"pvms-client/internal/config"
	"time"

	"gitlab.com/distributed_lab/logan/v3"
)

var (
	ticks = []string{"X", "0"}
)

const (
	url = "http://localhost:8000/integrations/pvms-final/move"
)

type service struct {
	log *logan.Entry
}

func (s *service) run() error {
	s.log.Info("Client started")
	client := &http.Client{}

	i := 1
	for {
		tick := ticks[rand.Intn(len(ticks))]
		x := rand.Intn(100) + 1
		y := rand.Intn(100) + 1

		var jsonStr = []byte(fmt.Sprintf("{ \"x\": %d, \"y\": %d, \"tick\": \"%s\" }", x, y, tick))
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
		if err != nil {
			s.log.Error("Failed to construct an http request")
			panic(err)
		}
		req.Header.Set("X-Custom-Header", "myvalue")
		req.Header.Set("Content-Type", "application/json")

		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}

		s.log.Infof("%d: response Status: %s", i, resp.Status)

		resp.Body.Close()
		i++
		time.Sleep(time.Millisecond * 100)
	}
}

func newService(cfg config.Config) *service {
	return &service{
		log: cfg.Log(),
	}
}

func Run(cfg config.Config) {
	if err := newService(cfg).run(); err != nil {
		panic(err)
	}
}
