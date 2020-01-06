package request

import (
	clientModel "github.com/TangiFavennec/go-service-sample/sample/service/client/model"
)

// PostDeck /decks POST request
type PostDeck struct {
	Deck clientModel.Deck
}
