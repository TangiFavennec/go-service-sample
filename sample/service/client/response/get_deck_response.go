package response

import (
	clientModel "github.com/TangiFavennec/go-service-sample/sample/service/client/model"
)

// GetDeck /decks GET response
type GetDeck struct {
	Deck clientModel.Deck `json:"deck,omitempty"`
	Err  error            `json:"err,omitempty"`
}

func (r GetDeck) error() error { return r.Err }
