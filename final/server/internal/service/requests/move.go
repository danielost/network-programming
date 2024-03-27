package requests

import (
	"encoding/json"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/pkg/errors"
)

type MoveRequest struct {
	Tick string `json:"tick"`
	X    int    `json:"x"`
	Y    int    `json:"y"`
}

func NewMoveRequest(r *http.Request) (MoveRequest, error) {
	var request MoveRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return request, errors.Wrap(err, "failed to unmarshal")
	}

	return request, request.validate()
}

func (r *MoveRequest) validate() error {
	return validation.Errors{
		"tick": validation.Validate(&r.Tick, validation.Required, validation.In("X", "0")),
		"x":    validation.Validate(&r.X, validation.Required, validation.Min(1), validation.Max(100)),
		"y":    validation.Validate(&r.Y, validation.Required, validation.Min(1), validation.Max(100)),
	}.Filter()
}
