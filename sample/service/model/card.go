package model

// Card is a field of a user Deck.
// ID should be unique within the Deck (at a minimum).
type Card struct {
	ID     string
	First  string
	Second string
}
