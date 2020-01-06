package request

import (
	clientModel "github.com/TangiFavennec/go-service-sample/sample/service/client/model"
)

// PostCard /decks/{id}/cards POST request
type PostCard struct {
	DeckID string
	Card   clientModel.Card
}
