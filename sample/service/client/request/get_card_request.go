package request

// GetCard /decks/{Deck_id}/cards/{Card_id} GET request
type GetCard struct {
	DeckID string
	CardID string
}
