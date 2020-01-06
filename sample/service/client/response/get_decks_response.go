package response

import (
	clientModel "github.com/TangiFavennec/go-service-sample/sample/service/client/model"
)

// GetDecks /decks GET response
type GetDecks struct {
	Decks []clientModel.Deck `json:"decks,omitempty"`
	Err   error              `json:"err,omitempty"`
}

func (r GetDecks) error() error { return r.Err }
