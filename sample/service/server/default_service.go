package server

import (
	"context"
	"sync"

	client "github.com/TangiFavennec/go-service-sample/sample/service/client/model"
	data "github.com/TangiFavennec/go-service-sample/sample/service/data"
	inmem "github.com/TangiFavennec/go-service-sample/sample/service/data/other"
	model "github.com/TangiFavennec/go-service-sample/sample/service/model"
	mapper "github.com/TangiFavennec/go-service-sample/sample/service/server/mappers"
)

type defaultService struct {
	mtx  sync.RWMutex
	repo data.SampleRepository
}

// NewdefaultService In Memory Service Constructor
func NewdefaultService() SampleService {
	return &defaultService{
		repo: inmem.NewInmemRepository(),
	}
}

func (s *defaultService) PostDeck(ctx context.Context, p model.Deck) error {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	return s.repo.PostDeck(p)
}

func (s *defaultService) GetDeck(ctx context.Context, id string) (client.Deck, error) {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	p, ok := s.repo.GetDeck(id)
	return mapper.ToClientDeck(p), ok
}

func (s *defaultService) PutDeck(ctx context.Context, id string, p model.Deck) error {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	return s.repo.PutDeck(id, p)
}

func (s *defaultService) GetDecks(ctx context.Context) ([]client.Deck, error) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	d, ok := s.repo.GetDecks()
	return mapper.ToClientDecks(d), ok
}

func (s *defaultService) DeleteDeck(ctx context.Context, id string) error {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	return s.repo.DeleteDeck(id)
}

func (s *defaultService) GetCards(ctx context.Context, DeckID string) ([]client.Card, error) {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	p, ok := s.repo.GetCards(DeckID)
	return mapper.ToClientCards(p), ok
}

func (s *defaultService) GetCard(ctx context.Context, DeckID string, CardID string) (client.Card, error) {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	a, ok := s.repo.GetCard(DeckID, CardID)
	return mapper.ToClientCard(a), ok
}

func (s *defaultService) PostCard(ctx context.Context, DeckID string, a model.Card) error {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	return s.repo.PostCard(DeckID, a)
}

func (s *defaultService) DeleteCard(ctx context.Context, DeckID string, CardID string) error {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	return s.repo.DeleteCard(DeckID, CardID)
}
