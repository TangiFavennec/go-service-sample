package response

// PutDeck /decks PUT response
type PutDeck struct {
	Err error `json:"err,omitempty"`
}

func (r PutDeck) error() error { return nil }
