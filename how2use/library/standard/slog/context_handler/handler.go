package main

import (
	"context"
	"log/slog"
)

const (
	LogKeyTraceID = "trace_id"
	LogKeyUserID  = "user_id"
)

var keys = []string{LogKeyTraceID, LogKeyUserID}

type AppLogHandler struct {
	slog.Handler
}

var _ slog.Handler = &AppLogHandler{}

func (h *AppLogHandler) Handle(ctx context.Context, r slog.Record) error {
	for _, key := range keys {
		if v := ctx.Value(key); v != nil {
			r.AddAttrs(slog.Attr{Key: key, Value: slog.AnyValue(v)})
		}
	}
	return h.Handler.Handle(ctx, r)
}
