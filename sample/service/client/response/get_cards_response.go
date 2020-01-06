package response

import (
	clientModel "github.com/TangiFavennec/go-service-sample/sample/service/client/model"
)

// GetCards /decks GET response
type GetCards struct {
	Cards []clientModel.Card `json:"cards,omitempty"`
	Err   error              `json:"err,omitempty"`
}

func (r GetCards) error() error { return r.Err }
