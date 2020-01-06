package endpoints

import (
	"context"
	"net/url"
	"strings"

	clientModel "github.com/TangiFavennec/go-service-sample/sample/service/client/model"
	clientRequest "github.com/TangiFavennec/go-service-sample/sample/service/client/request"
	clientResponse "github.com/TangiFavennec/go-service-sample/sample/service/client/response"
	model "github.com/TangiFavennec/go-service-sample/sample/service/model"
	server "github.com/TangiFavennec/go-service-sample/sample/service/server"
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"

	mapper "github.com/TangiFavennec/go-service-sample/sample/service/server/mappers"
)

// Endpoints for current microservice
type Endpoints struct {
	PostDeckEndpoint   endpoint.Endpoint
	GetDeckEndpoint    endpoint.Endpoint
	PutDeckEndpoint    endpoint.Endpoint
	GetDecksEndpoint   endpoint.Endpoint
	DeleteDeckEndpoint endpoint.Endpoint
	GetCardsEndpoint   endpoint.Endpoint
	GetCardEndpoint    endpoint.Endpoint
	PostCardEndpoint   endpoint.Endpoint
	DeleteCardEndpoint endpoint.Endpoint
}

// MakeServerEndpoints returns an Endpoints struct where each endpoint invokes
// the corresponding method on the provided service. Useful in a decksvc
// server.
func MakeServerEndpoints(s server.SampleService) Endpoints {
	return Endpoints{
		PostDeckEndpoint:   MakePostDeckEndpoint(s),
		GetDeckEndpoint:    MakeGetDeckEndpoint(s),
		PutDeckEndpoint:    MakePutDeckEndpoint(s),
		GetDecksEndpoint:   MakeGetDecksEndpoint(s),
		DeleteDeckEndpoint: MakeDeleteDeckEndpoint(s),
		GetCardsEndpoint:   MakeGetCardsEndpoint(s),
		GetCardEndpoint:    MakeGetCardEndpoint(s),
		PostCardEndpoint:   MakePostCardEndpoint(s),
		DeleteCardEndpoint: MakeDeleteCardEndpoint(s),
	}
}

// MakeClientEndpoints returns an Endpoints struct where each endpoint invokes
// the corresponding method on the remote instance, via a transport/http.Client.
// Useful in a decksvc client.
func MakeClientEndpoints(instance string) (Endpoints, error) {
	if !strings.HasPrefix(instance, "http") {
		instance = "http://" + instance
	}
	tgt, err := url.Parse(instance)
	if err != nil {
		return Endpoints{}, err
	}
	tgt.Path = ""

	options := []httptransport.ClientOption{}

	// Note that the request encoders need to modify the request URL, changing
	// the path. That's fine: we simply need to provide specific encoders for
	// each endpoint.

	return Endpoints{
		PostDeckEndpoint:   httptransport.NewClient("POST", tgt, encodePostDeckRequest, decodePostDeckResponse, options...).Endpoint(),
		GetDeckEndpoint:    httptransport.NewClient("GET", tgt, encodeGetDeckRequest, decodeGetDeckResponse, options...).Endpoint(),
		PutDeckEndpoint:    httptransport.NewClient("PUT", tgt, encodePutDeckRequest, decodePutDeckResponse, options...).Endpoint(),
		GetDecksEndpoint:   httptransport.NewClient("GET", tgt, encodeGetDecksRequest, decodeGetDecksResponse, options...).Endpoint(),
		DeleteDeckEndpoint: httptransport.NewClient("DELETE", tgt, encodeDeleteDeckRequest, decodeDeleteDeckResponse, options...).Endpoint(),
		GetCardsEndpoint:   httptransport.NewClient("GET", tgt, encodeGetCardsRequest, decodeGetCardsResponse, options...).Endpoint(),
		GetCardEndpoint:    httptransport.NewClient("GET", tgt, encodeGetCardRequest, decodeGetCardResponse, options...).Endpoint(),
		PostCardEndpoint:   httptransport.NewClient("POST", tgt, encodePostCardRequest, decodePostCardResponse, options...).Endpoint(),
		DeleteCardEndpoint: httptransport.NewClient("DELETE", tgt, encodeDeleteCardRequest, decodeDeleteCardResponse, options...).Endpoint(),
	}, nil
}

// PostDeck implements Service. Primarily useful in a client.
func (e Endpoints) PostDeck(ctx context.Context, p model.Deck) error {
	request := clientRequest.PostDeck{Deck: mapper.ToClientDeck(p)}
	response, err := e.PostDeckEndpoint(ctx, request)
	if err != nil {
		return err
	}
	resp := response.(clientResponse.PostDeck)
	return resp.Err
}

// GetDeck implements Service. Primarily useful in a client.
func (e Endpoints) GetDeck(ctx context.Context, id string) (clientModel.Deck, error) {
	request := clientRequest.GetDeck{ID: id}
	response, err := e.GetDeckEndpoint(ctx, request)
	if err != nil {
		return mapper.ToClientDeck(model.Deck{}), err
	}
	resp := response.(clientResponse.GetDeck)
	return resp.Deck, resp.Err
}

// PutDeck implements Service. Primarily useful in a client.
func (e Endpoints) PutDeck(ctx context.Context, id string, p model.Deck) error {
	request := clientRequest.PutDeck{ID: id, Deck: mapper.ToClientDeck(p)}
	response, err := e.PutDeckEndpoint(ctx, request)
	if err != nil {
		return err
	}
	resp := response.(clientResponse.PutDeck)
	return resp.Err
}

// GetDecks implements Service. Primarily useful in a client.
func (e Endpoints) GetDecks(ctx context.Context) ([]clientModel.Deck, error) {
	response, err := e.GetDecksEndpoint(ctx, nil)
	if err != nil {
		return nil, err
	}
	resp := response.(clientResponse.GetDecks)
	return resp.Decks, resp.Err
}

// DeleteDeck implements Service. Primarily useful in a client.
func (e Endpoints) DeleteDeck(ctx context.Context, id string) error {
	request := clientRequest.DeleteDeck{ID: id}
	response, err := e.DeleteDeckEndpoint(ctx, request)
	if err != nil {
		return err
	}
	resp := response.(clientResponse.DeleteDeck)
	return resp.Err
}

// GetCards implements Service. Primarily useful in a client.
func (e Endpoints) GetCards(ctx context.Context, deckID string) ([]clientModel.Card, error) {
	request := clientRequest.GetCards{DeckID: deckID}
	response, err := e.GetCardsEndpoint(ctx, request)
	if err != nil {
		return nil, err
	}
	resp := response.(clientResponse.GetCards)
	return resp.Cards, resp.Err
}

// GetCard implements Service. Primarily useful in a client.
func (e Endpoints) GetCard(ctx context.Context, deckID string, cardID string) (clientModel.Card, error) {
	request := clientRequest.GetCard{DeckID: deckID, CardID: cardID}
	response, err := e.GetCardEndpoint(ctx, request)
	if err != nil {
		return mapper.ToClientCard(model.Card{}), err
	}
	resp := response.(clientResponse.GetCard)
	return resp.Card, resp.Err
}

// PostCard implements Service. Primarily useful in a client.
func (e Endpoints) PostCard(ctx context.Context, deckID string, a model.Card) error {
	request := clientRequest.PostCard{DeckID: deckID, Card: mapper.ToClientCard(a)}
	response, err := e.PostCardEndpoint(ctx, request)
	if err != nil {
		return err
	}
	resp := response.(clientResponse.PostCard)
	return resp.Err
}

// DeleteCard implements Service. Primarily useful in a client.
func (e Endpoints) DeleteCard(ctx context.Context, deckID string, cardID string) error {
	request := clientRequest.DeleteCard{DeckID: deckID, CardID: cardID}
	response, err := e.DeleteCardEndpoint(ctx, request)
	if err != nil {
		return err
	}
	resp := response.(clientResponse.DeleteCard)
	return resp.Err
}

// MakePostDeckEndpoint returns an endpoint via the passed service.
// Primarily useful in a server.
func MakePostDeckEndpoint(s server.SampleService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(clientRequest.PostDeck)
		e := s.PostDeck(ctx, mapper.FromClientDeck(req.Deck))
		return clientResponse.PostDeck{Err: e}, e
	}
}

// MakeGetDeckEndpoint returns an endpoint via the passed service.
// Primarily useful in a server.
func MakeGetDeckEndpoint(s server.SampleService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(clientRequest.GetDeck)
		p, e := s.GetDeck(ctx, req.ID)
		return clientResponse.GetDeck{Deck: p, Err: e}, e
	}
}

// MakePutDeckEndpoint returns an endpoint via the passed service.
// Primarily useful in a server.
func MakePutDeckEndpoint(s server.SampleService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(clientRequest.PutDeck)
		e := s.PutDeck(ctx, req.ID, mapper.FromClientDeck(req.Deck))
		return clientResponse.PutDeck{Err: e}, e
	}
}

// MakeGetDecksEndpoint returns an endpoint via the passed service.
// Primarily useful in a server.
func MakeGetDecksEndpoint(s server.SampleService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		decks, e := s.GetDecks(ctx)
		return clientResponse.GetDecks{Decks: decks, Err: e}, e
	}
}

// MakeDeleteDeckEndpoint returns an endpoint via the passed service.
// Primarily useful in a server.
func MakeDeleteDeckEndpoint(s server.SampleService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(clientRequest.DeleteDeck)
		e := s.DeleteDeck(ctx, req.ID)
		return clientResponse.DeleteDeck{Err: e}, e
	}
}

// MakeGetCardsEndpoint returns an endpoint via the passed service.
// Primarily useful in a server.
func MakeGetCardsEndpoint(s server.SampleService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(clientRequest.GetCards)
		a, e := s.GetCards(ctx, req.DeckID)
		return clientResponse.GetCards{Cards: a, Err: e}, e
	}
}

// MakeGetCardEndpoint returns an endpoint via the passed service.
// Primarily useful in a server.
func MakeGetCardEndpoint(s server.SampleService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(clientRequest.GetCard)
		a, e := s.GetCard(ctx, req.DeckID, req.CardID)
		return clientResponse.GetCard{Card: a, Err: e}, e
	}
}

// MakePostCardEndpoint returns an endpoint via the passed service.
// Primarily useful in a server.
func MakePostCardEndpoint(s server.SampleService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(clientRequest.PostCard)
		e := s.PostCard(ctx, req.DeckID, mapper.FromClientCard(req.Card))
		return clientResponse.PostCard{Err: e}, e
	}
}

// MakeDeleteCardEndpoint returns an endpoint via the passed service.
// Primarily useful in a server.
func MakeDeleteCardEndpoint(s server.SampleService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(clientRequest.DeleteCard)
		e := s.DeleteCard(ctx, req.DeckID, req.CardID)
		return clientResponse.DeleteCard{Err: e}, e
	}
}
