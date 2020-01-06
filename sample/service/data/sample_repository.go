package data

import (
	"errors"

	model "github.com/TangiFavennec/go-service-sample/sample/service/model"
)

// SampleRepository is a simple CRUD interface for user Decks.
type SampleRepository interface {
	PostDeck(p model.Deck) error
	GetDeck(id string) (model.Deck, error)
	PutDeck(id string, p model.Deck) error
	GetDecks() ([]model.Deck, error)
	DeleteDeck(id string) error
	GetCards(DeckID string) ([]model.Card, error)
	GetCard(DeckID string, CardID string) (model.Card, error)
	PostCard(DeckID string, a model.Card) error
	DeleteCard(DeckID string, CardID string) error
}

var (
	// ErrInconsistentIDs : Inconsistent ids error
	ErrInconsistentIDs = errors.New("inconsistent IDs")
	// ErrAlreadyExists : Item exists already
	ErrAlreadyExists = errors.New("already exists")
	// ErrNotFound : Item not found
	ErrNotFound = errors.New("not found")
)
