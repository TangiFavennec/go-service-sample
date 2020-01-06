package response

// PostDeck /decks POST response
type PostDeck struct {
	Err error `json:"err,omitempty"`
}

func (r PostDeck) error() error { return r.Err }
