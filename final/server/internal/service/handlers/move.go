package handlers

import (
	"net/http"
	"pvms-final/internal/service/requests"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

var HttpQueue = make(chan requests.MoveRequest, 50)

func Move(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewMoveRequest(r)
	if err != nil {
		Log(r).WithError(err).Error("failed to parse Move request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	Log(r).Infof("Putting request to HTTP queue {%s: %d %d}", request.Tick, request.X, request.Y)
	Log(r).Infof("Chan len: %d", len(HttpQueue))

	HttpQueue <- request

	Log(r).Info("Request has been put")

	w.WriteHeader(http.StatusNoContent)
}
