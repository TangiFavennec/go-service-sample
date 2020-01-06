package response

// DeleteDeck /decks/{Deck_id} DELETE response
type DeleteDeck struct {
	Err error `json:"err,omitempty"`
}

func (r DeleteDeck) error() error { return r.Err }
