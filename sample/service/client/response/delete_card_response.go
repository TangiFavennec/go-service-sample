package response

// DeleteCard /decks/{Deck_id}/cards/{Card_id} DELETE response
type DeleteCard struct {
	Err error `json:"err,omitempty"`
}

func (r DeleteCard) error() error { return r.Err }
