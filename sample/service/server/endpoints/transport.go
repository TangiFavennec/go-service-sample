package endpoints

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	httptransport "github.com/go-kit/kit/transport/http"

	clientModel "github.com/TangiFavennec/go-service-sample/sample/service/client/model"
	clientRequest "github.com/TangiFavennec/go-service-sample/sample/service/client/request"
	clientResponse "github.com/TangiFavennec/go-service-sample/sample/service/client/response"
	data "github.com/TangiFavennec/go-service-sample/sample/service/data"
	"github.com/TangiFavennec/go-service-sample/sample/service/server"
)

var (
	// ErrBadRouting is returned when an expected path variable is missing.
	// It always indicates programmer error.
	ErrBadRouting = errors.New("inconsistent mapping between route and handler (programmer error)")
)

// MakeHTTPHandler mounts all of the service endpoints into an http.Handler.
// Useful in a decksvc server.
func MakeHTTPHandler(s server.SampleService, logger log.Logger) http.Handler {
	r := mux.NewRouter()
	e := MakeServerEndpoints(s)
	options := []httptransport.ServerOption{
		httptransport.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		httptransport.ServerErrorEncoder(encodeError),
	}

	// POST    /decks/                          adds another Deck
	// GET     /decks/:id                       retrieves the given Deck by id
	// PUT     /decks/:id                       post updated Deck information about the Deck
	// PATCH   /decks/:id                       partial updated Deck information
	// DELETE  /decks/:id                       remove the given Deck
	// GET     /decks/:id/cards                retrieve Cards associated with the Deck
	// GET     /decks/:id/cards/:cardID         retrieve a particular Deck Card
	// POST    /decks/:id/cards                add a new Card
	// DELETE  /decks/:id/cards/:cardID         remove an Card

	r.Methods("POST").Path("/decks").Handler(httptransport.NewServer(
		e.PostDeckEndpoint,
		decodePostDeckRequest,
		encodeResponse,
		options...,
	))
	r.Methods("GET").Path("/decks/{id}").Handler(httptransport.NewServer(
		e.GetDeckEndpoint,
		decodeGetDeckRequest,
		encodeResponse,
		options...,
	))
	r.Methods("PUT").Path("/decks/{id}").Handler(httptransport.NewServer(
		e.PutDeckEndpoint,
		decodePutDeckRequest,
		encodeResponse,
		options...,
	))
	r.Methods("GET").Path("/decks").Handler(httptransport.NewServer(
		e.GetDecksEndpoint,
		decodeGetDecksRequest,
		encodeResponse,
		options...,
	))
	r.Methods("DELETE").Path("/decks/{id}").Handler(httptransport.NewServer(
		e.DeleteDeckEndpoint,
		decodeDeleteDeckRequest,
		encodeResponse,
		options...,
	))
	r.Methods("GET").Path("/decks/{id}/cards").Handler(httptransport.NewServer(
		e.GetCardsEndpoint,
		decodeGetCardsRequest,
		encodeResponse,
		options...,
	))
	r.Methods("GET").Path("/decks/{id}/cards/{cardID}").Handler(httptransport.NewServer(
		e.GetCardEndpoint,
		decodeGetCardRequest,
		encodeResponse,
		options...,
	))
	r.Methods("POST").Path("/decks/{id}/cards").Handler(httptransport.NewServer(
		e.PostCardEndpoint,
		decodePostCardRequest,
		encodeResponse,
		options...,
	))
	r.Methods("DELETE").Path("/decks/{id}/cards/{cardID}").Handler(httptransport.NewServer(
		e.DeleteCardEndpoint,
		decodeDeleteCardRequest,
		encodeResponse,
		options...,
	))
	return r
}

func decodePostDeckRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req clientRequest.PostDeck
	if e := json.NewDecoder(r.Body).Decode(&req.Deck); e != nil {
		return nil, e
	}
	return req, nil
}

func decodeGetDeckRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, ErrBadRouting
	}
	return clientRequest.GetDeck{ID: id}, nil
}

func decodePutDeckRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, ErrBadRouting
	}
	var Deck clientModel.Deck
	if err := json.NewDecoder(r.Body).Decode(&Deck); err != nil {
		return nil, err
	}
	return clientRequest.PutDeck{
		ID:   id,
		Deck: Deck,
	}, nil
}

func decodeGetDecksRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	return nil, nil
}

func decodeDeleteDeckRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, ErrBadRouting
	}
	return clientRequest.DeleteDeck{ID: id}, nil
}

func decodeGetCardsRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, ErrBadRouting
	}
	return clientRequest.GetCards{DeckID: id}, nil
}

func decodeGetCardRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, ErrBadRouting
	}
	cardID, ok := vars["cardID"]
	if !ok {
		return nil, ErrBadRouting
	}
	return clientRequest.GetCard{
		DeckID: id,
		CardID: cardID,
	}, nil
}

func decodePostCardRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, ErrBadRouting
	}
	var Card clientModel.Card
	if err := json.NewDecoder(r.Body).Decode(&Card); err != nil {
		return nil, err
	}
	return clientRequest.PostCard{
		DeckID: id,
		Card:   Card,
	}, nil
}

func decodeDeleteCardRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, ErrBadRouting
	}
	cardID, ok := vars["cardID"]
	if !ok {
		return nil, ErrBadRouting
	}
	return clientRequest.DeleteCard{
		DeckID: id,
		CardID: cardID,
	}, nil
}

func encodePostDeckRequest(ctx context.Context, req *http.Request, request interface{}) error {
	// r.Methods("POST").Path("/decks")
	req.URL.Path = "/decks"
	return encodeRequest(ctx, req, request)
}

func encodeGetDeckRequest(ctx context.Context, req *http.Request, request interface{}) error {
	// r.Methods("GET").Path("/decks/{id}")
	r := request.(clientRequest.GetDeck)
	deckID := url.QueryEscape(r.ID)
	req.URL.Path = "/decks/" + deckID
	return encodeRequest(ctx, req, request)
}

func encodePutDeckRequest(ctx context.Context, req *http.Request, request interface{}) error {
	// r.Methods("PUT").Path("/decks/{id}")
	r := request.(clientRequest.PutDeck)
	deckID := url.QueryEscape(r.ID)
	req.URL.Path = "/decks/" + deckID
	return encodeRequest(ctx, req, request)
}

func encodeGetDecksRequest(ctx context.Context, req *http.Request, request interface{}) error {
	// r.Methods("GET").Path("/decks")
	req.URL.Path = "/decks"
	return encodeRequest(ctx, req, request)
}

func encodeDeleteDeckRequest(ctx context.Context, req *http.Request, request interface{}) error {
	// r.Methods("DELETE").Path("/decks/{id}")
	r := request.(clientRequest.DeleteDeck)
	deckID := url.QueryEscape(r.ID)
	req.URL.Path = "/decks/" + deckID
	return encodeRequest(ctx, req, request)
}

func encodeGetCardsRequest(ctx context.Context, req *http.Request, request interface{}) error {
	// r.Methods("GET").Path("/decks/{id}/cards")
	r := request.(clientRequest.GetCards)
	deckID := url.QueryEscape(r.DeckID)
	req.URL.Path = "/decks/" + deckID + "/cards"
	return encodeRequest(ctx, req, request)
}

func encodeGetCardRequest(ctx context.Context, req *http.Request, request interface{}) error {
	// r.Methods("GET").Path("/decks/{id}/cards/{cardID}")
	r := request.(clientRequest.GetCard)
	deckID := url.QueryEscape(r.DeckID)
	cardID := url.QueryEscape(r.CardID)
	req.URL.Path = "/decks/" + deckID + "/cards/" + cardID
	return encodeRequest(ctx, req, request)
}

func encodePostCardRequest(ctx context.Context, req *http.Request, request interface{}) error {
	// r.Methods("POST").Path("/decks/{id}/cards")
	r := request.(clientRequest.PostCard)
	deckID := url.QueryEscape(r.DeckID)
	req.URL.Path = "/decks/" + deckID + "/cards"
	return encodeRequest(ctx, req, request)
}

func encodeDeleteCardRequest(ctx context.Context, req *http.Request, request interface{}) error {
	// r.Methods("DELETE").Path("/decks/{id}/cards/{cardID}")
	r := request.(clientRequest.DeleteCard)
	deckID := url.QueryEscape(r.DeckID)
	cardID := url.QueryEscape(r.CardID)
	req.URL.Path = "/decks/" + deckID + "/cards/" + cardID
	return encodeRequest(ctx, req, request)
}

func decodePostDeckResponse(_ context.Context, resp *http.Response) (interface{}, error) {
	var response clientResponse.PostDeck
	err := json.NewDecoder(resp.Body).Decode(&response)
	return response, err
}

func decodeGetDeckResponse(_ context.Context, resp *http.Response) (interface{}, error) {
	var response clientResponse.GetDeck
	err := json.NewDecoder(resp.Body).Decode(&response)
	return response, err
}

func decodePutDeckResponse(_ context.Context, resp *http.Response) (interface{}, error) {
	var response clientResponse.PutDeck
	err := json.NewDecoder(resp.Body).Decode(&response)
	return response, err
}

func decodeGetDecksResponse(_ context.Context, resp *http.Response) (interface{}, error) {
	var response clientResponse.GetDecks
	err := json.NewDecoder(resp.Body).Decode(&response)
	return response, err
}

func decodeDeleteDeckResponse(_ context.Context, resp *http.Response) (interface{}, error) {
	var response clientResponse.DeleteDeck
	err := json.NewDecoder(resp.Body).Decode(&response)
	return response, err
}

func decodeGetCardResponse(_ context.Context, resp *http.Response) (interface{}, error) {
	var response clientResponse.GetCard
	err := json.NewDecoder(resp.Body).Decode(&response)
	return response, err
}

func decodeGetCardsResponse(_ context.Context, resp *http.Response) (interface{}, error) {
	var response clientResponse.GetCards
	err := json.NewDecoder(resp.Body).Decode(&response)
	return response, err
}

func decodePostCardResponse(_ context.Context, resp *http.Response) (interface{}, error) {
	var response clientResponse.PostCard
	err := json.NewDecoder(resp.Body).Decode(&response)
	return response, err
}

func decodeDeleteCardResponse(_ context.Context, resp *http.Response) (interface{}, error) {
	var response clientResponse.DeleteCard
	err := json.NewDecoder(resp.Body).Decode(&response)
	return response, err
}

// errorer is implemented by all concrete response types that may contain
// errors. It allows us to change the HTTP response code without needing to
// trigger an endpoint (transport-level) error. For more information, read the
// big comment in endpoints.go.
type errorer interface {
	error() error
}

// encodeResponse is the common method to encode all response types to the
// clientRequest. I chose to do it this way because, since we're using JSON, there's no
// reason to provide anything more specific. It's certainly possible to
// specialize on a per-response (per-method) basis.
func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		// Not a Go kit transport error, but a business-logic error.
		// Provide those as HTTP errors.
		encodeError(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

// encodeRequest likewise JSON-encodes the request to the HTTP request body.
// Don't use it directly as a transport/http.Client EncodeRequestFunc:
// decksvc endpoints require mutating the HTTP method and request path.
func encodeRequest(_ context.Context, req *http.Request, request interface{}) error {
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(request)
	if err != nil {
		return err
	}
	req.Body = ioutil.NopCloser(&buf)
	return nil
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil error")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(codeFrom(err))
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func codeFrom(err error) int {
	switch err {
	case data.ErrNotFound:
		return http.StatusNotFound
	case data.ErrAlreadyExists, data.ErrInconsistentIDs:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
