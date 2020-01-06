package server

import (
	"context"

	client "github.com/TangiFavennec/go-service-sample/sample/service/client/model"
	model "github.com/TangiFavennec/go-service-sample/sample/service/model"
)

// SampleService is a simple CRUD interface for user Decks.
type SampleService interface {
	PostDeck(ctx context.Context, p model.Deck) error
	GetDeck(ctx context.Context, id string) (client.Deck, error)
	PutDeck(ctx context.Context, id string, p model.Deck) error
	GetDecks(ctx context.Context) ([]client.Deck, error)
	DeleteDeck(ctx context.Context, id string) error
	GetCards(ctx context.Context, DeckID string) ([]client.Card, error)
	GetCard(ctx context.Context, DeckID string, CardID string) (client.Card, error)
	PostCard(ctx context.Context, DeckID string, a model.Card) error
	DeleteCard(ctx context.Context, DeckID string, CardID string) error
}
