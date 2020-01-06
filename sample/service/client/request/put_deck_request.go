package request

import (
	clientModel "github.com/TangiFavennec/go-service-sample/sample/service/client/model"
)

// PutDeck /decks PUT request
type PutDeck struct {
	ID      string
	Deck clientModel.Deck
}
