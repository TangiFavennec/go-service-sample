package response

import (
	clientModel "github.com/TangiFavennec/go-service-sample/sample/service/client/model"
)

// GetCard /decks/{Deck_id}/cards/{Card_id} GET response
type GetCard struct {
	Card clientModel.Card `json:"card,omitempty"`
	Err  error            `json:"err,omitempty"`
}

func (r GetCard) error() error { return r.Err }
