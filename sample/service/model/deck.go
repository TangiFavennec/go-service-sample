package model

// Deck represents a single user Deck.
// ID should be globally unique.
type Deck struct {
	ID        string    
	Name      string    
	Cards []Card 
}