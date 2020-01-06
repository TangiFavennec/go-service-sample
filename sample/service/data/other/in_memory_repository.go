package inmemory

import (
	data "github.com/TangiFavennec/go-service-sample/sample/service/data"
	model "github.com/TangiFavennec/go-service-sample/sample/service/model"
)

type repository struct {
	m map[string]model.Deck
}

// NewInmemRepository In Memory Service Constructor
func NewInmemRepository() data.SampleRepository {
	return &repository{
		m: map[string]model.Deck{},
	}
}

func (s *repository) PostDeck(p model.Deck) error {
	if _, ok := s.m[p.ID]; ok {
		return data.ErrAlreadyExists // POST = create, don't overwrite
	}
	s.m[p.ID] = p
	return nil
}

func (s *repository) GetDeck(id string) (model.Deck, error) {
	p, ok := s.m[id]
	if !ok {
		return model.Deck{}, data.ErrNotFound
	}
	return p, nil
}

func (s *repository) PutDeck(id string, p model.Deck) error {
	if id != p.ID {
		return data.ErrInconsistentIDs
	}
	s.m[id] = p // PUT = create or update
	return nil
}

func (s *repository) GetDecks() ([]model.Deck, error) {
	decks := []model.Deck{}
    for _, val := range s.m {
        decks = append(decks, val)
    }
	return decks, nil
}

func (s *repository) DeleteDeck(id string) error {
	if _, ok := s.m[id]; !ok {
		return data.ErrNotFound
	}
	delete(s.m, id)
	return nil
}

func (s *repository) GetCards(DeckID string) ([]model.Card, error) {
	p, ok := s.m[DeckID]
	if !ok {
		return []model.Card{}, data.ErrNotFound
	}
	return p.Cards, nil
}

func (s *repository) GetCard(DeckID string, CardID string) (model.Card, error) {
	p, ok := s.m[DeckID]
	if !ok {
		return model.Card{}, data.ErrNotFound
	}
	for _, Card := range p.Cards {
		if Card.ID == CardID {
			return Card, nil
		}
	}
	return model.Card{}, data.ErrNotFound
}

func (s *repository) PostCard(DeckID string, a model.Card) error {
	p, ok := s.m[DeckID]
	if !ok {
		return data.ErrNotFound
	}
	for _, Card := range p.Cards {
		if Card.ID == a.ID {
			return data.ErrAlreadyExists
		}
	}
	p.Cards = append(p.Cards, a)
	s.m[DeckID] = p
	return nil
}

func (s *repository) DeleteCard(DeckID string, CardID string) error {
	p, ok := s.m[DeckID]
	if !ok {
		return data.ErrNotFound
	}
	newCards := make([]model.Card, 0, len(p.Cards))
	for _, Card := range p.Cards {
		if Card.ID == CardID {
			continue // delete
		}
		newCards = append(newCards, Card)
	}
	if len(newCards) == len(p.Cards) {
		return data.ErrNotFound
	}
	p.Cards = newCards
	s.m[DeckID] = p
	return nil
}
