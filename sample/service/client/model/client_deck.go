package model

// Deck represents a single user Deck.
// ID should be globally unique.
type Deck struct {
	ID    string `json:"id"`
	Name  string `json:"name,omitempty"`
	Cards []Card `json:"cards,omitempty"`
}
