package request

// DeleteCard /decks/{Deck_id}/cards/{Card_id} DELETE request
type DeleteCard struct {
	DeckID string
	CardID string
}
