package handlers

import (
	"context"
	"net/http"
	"pvms-final/internal/service/requests"

	"gitlab.com/distributed_lab/logan/v3"
)

type ctxKey int

const (
	logCtxKey ctxKey = iota
	boardCtxKey
	queueKey
	statusChanKey
)

func CtxLog(entry *logan.Entry) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, logCtxKey, entry)
	}
}

func Log(r *http.Request) *logan.Entry {
	return r.Context().Value(logCtxKey).(*logan.Entry)
}

func CtxBoard(entry *[100][100]string) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, boardCtxKey, entry)
	}
}

func Board(r *http.Request) *[100][100]string {
	return r.Context().Value(boardCtxKey).(*[100][100]string)
}

func CtxQueue(entry chan requests.MoveRequest) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, queueKey, entry)
	}
}

func Queue(r *http.Request) chan requests.MoveRequest {
	return r.Context().Value(queueKey).(chan requests.MoveRequest)
}
