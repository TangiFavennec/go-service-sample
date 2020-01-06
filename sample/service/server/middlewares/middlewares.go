package server

import (
	"context"
	"time"

	clientModel "github.com/TangiFavennec/go-service-sample/sample/service/client/model"
	"github.com/TangiFavennec/go-service-sample/sample/service/model"
	"github.com/TangiFavennec/go-service-sample/sample/service/server"
	"github.com/go-kit/kit/log"
)

// Middleware describes a service (as opposed to endpoint) middleware.
type Middleware func(server.SampleService) server.SampleService

// LoggingMiddleware : Plug MiddleWare to input SampleService
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next server.SampleService) server.SampleService {
		return &loggingMiddleware{
			next:   next,
			logger: logger,
		}
	}
}

type loggingMiddleware struct {
	next   server.SampleService
	logger log.Logger
}

func (mw loggingMiddleware) PostDeck(ctx context.Context, p model.Deck) (err error) {
	defer func(begin time.Time) {
		mw.logger.Log("method", "PostDeck", "id", p.ID, "took", time.Since(begin), "err", err)
	}(time.Now())
	return mw.next.PostDeck(ctx, p)
}

func (mw loggingMiddleware) GetDeck(ctx context.Context, id string) (p clientModel.Deck, err error) {
	defer func(begin time.Time) {
		mw.logger.Log("method", "GetDeck", "id", id, "took", time.Since(begin), "err", err)
	}(time.Now())
	return mw.next.GetDeck(ctx, id)
}

func (mw loggingMiddleware) PutDeck(ctx context.Context, id string, p model.Deck) (err error) {
	defer func(begin time.Time) {
		mw.logger.Log("method", "PutDeck", "id", id, "took", time.Since(begin), "err", err)
	}(time.Now())
	return mw.next.PutDeck(ctx, id, p)
}

func (mw loggingMiddleware) GetDecks(ctx context.Context) (decks []clientModel.Deck, err error) {
	defer func(begin time.Time) {
		mw.logger.Log("method", "GetDecks", "took", time.Since(begin), "err", err)
	}(time.Now())
	return mw.next.GetDecks(ctx)
}

func (mw loggingMiddleware) DeleteDeck(ctx context.Context, id string) (err error) {
	defer func(begin time.Time) {
		mw.logger.Log("method", "DeleteDeck", "id", id, "took", time.Since(begin), "err", err)
	}(time.Now())
	return mw.next.DeleteDeck(ctx, id)
}

func (mw loggingMiddleware) GetCards(ctx context.Context, DeckID string) (cards []clientModel.Card, err error) {
	defer func(begin time.Time) {
		mw.logger.Log("method", "GetCards", "DeckID", DeckID, "took", time.Since(begin), "err", err)
	}(time.Now())
	return mw.next.GetCards(ctx, DeckID)
}

func (mw loggingMiddleware) GetCard(ctx context.Context, DeckID string, CardID string) (a clientModel.Card, err error) {
	defer func(begin time.Time) {
		mw.logger.Log("method", "GetCard", "DeckID", DeckID, "CardID", CardID, "took", time.Since(begin), "err", err)
	}(time.Now())
	return mw.next.GetCard(ctx, DeckID, CardID)
}

func (mw loggingMiddleware) PostCard(ctx context.Context, DeckID string, a model.Card) (err error) {
	defer func(begin time.Time) {
		mw.logger.Log("method", "PostCard", "DeckID", DeckID, "took", time.Since(begin), "err", err)
	}(time.Now())
	return mw.next.PostCard(ctx, DeckID, a)
}

func (mw loggingMiddleware) DeleteCard(ctx context.Context, DeckID string, CardID string) (err error) {
	defer func(begin time.Time) {
		mw.logger.Log("method", "DeleteCard", "DeckID", DeckID, "CardID", CardID, "took", time.Since(begin), "err", err)
	}(time.Now())
	return mw.next.DeleteCard(ctx, DeckID, CardID)
}
