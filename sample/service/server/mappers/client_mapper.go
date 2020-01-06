package mapper

import (
	client "github.com/TangiFavennec/go-service-sample/sample/service/client/model"
	model "github.com/TangiFavennec/go-service-sample/sample/service/model"
)

// ToClientCard : Card model object to Card client object
func ToClientCard(input model.Card) client.Card {
	return client.Card{
		ID:     input.ID,
		First:  input.First,
		Second: input.Second,
	}
}

// ToClientCards : Card model object to Card client object (list version)
func ToClientCards(inputList []model.Card) []client.Card {
	var res []client.Card
	for _, val := range inputList {
		res = append(res, ToClientCard(val))
	}
	return res
}

// ToClientDeck : Deck model object to Deck client object
func ToClientDeck(input model.Deck) client.Deck {
	return client.Deck{
		ID:    input.ID,
		Name:  input.Name,
		Cards: ToClientCards(input.Cards),
	}
}

// ToClientDecks : Deck model object to Deck client object (list version)
func ToClientDecks(inputList []model.Deck) []client.Deck {
	var res []client.Deck
	for _, val := range inputList {
		res = append(res, ToClientDeck(val))
	}
	return res
}

// FromClientCard : Card client object to Card model object
func FromClientCard(input client.Card) model.Card {
	return model.Card{
		ID:     input.ID,
		First:  input.First,
		Second: input.Second,
	}
}

// FromClientCards : Card model object to Card client object (list version)
func FromClientCards(inputList []client.Card) []model.Card {
	var res []model.Card
	for _, val := range inputList {
		res = append(res, FromClientCard(val))
	}
	return res
}

// FromClientDeck : Deck client object to Deck model object
func FromClientDeck(input client.Deck) model.Deck {
	return model.Deck{
		ID:    input.ID,
		Name:  input.Name,
		Cards: FromClientCards(input.Cards),
	}
}

// FromClientDecks : Deck model object to Deck client object (list version)
func FromClientDecks(inputList []client.Deck) []model.Deck {
	var res []model.Deck
	for _, val := range inputList {
		res = append(res, FromClientDeck(val))
	}
	return res
}
