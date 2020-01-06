package response

// PostCard /decks/{id}/cards POST response
type PostCard struct {
	Err error `json:"err,omitempty"`
}

func (r PostCard) error() error { return r.Err }
