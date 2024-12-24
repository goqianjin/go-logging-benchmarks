package bench

import (
	"context"
	"fmt"
	"io"
	"log/slog"
)

func slogAttrs() []slog.Attr {
	return []slog.Attr{
		slog.Int("bytes", ctxBodyBytes),
		slog.String("request", ctxRequest),
		slog.Float64("elapsed_time_ms", ctxTimeElapsedMs),
		slog.Any("user", ctxUser),
		slog.Time("now", ctxTime),
		slog.Any("months", ctxMonths),
		slog.Any("primes", ctxFirst10Primes),
		slog.Any("users", ctxUsers),
		slog.Any("error", ctxErr),
	}
}

func newSlog(w io.Writer, format string) *slog.Logger {
	if format == "text" {
		return slog.New(slog.NewTextHandler(w, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		}))
	}
	return slog.New(slog.NewJSONHandler(w, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))

}

func newSlogWithCtx(w io.Writer, attr []slog.Attr, format string) *slog.Logger {
	if format == "text" {
		return slog.New(slog.NewTextHandler(w, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		}).WithAttrs(attr))
	}
	return slog.New(slog.NewJSONHandler(w, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}).WithAttrs(attr))
}

type slogBench struct {
	l *slog.Logger

	format string
}

func (b *slogBench) new(w io.Writer) logBenchmark {
	return &slogBench{
		l: newSlog(w, b.format),
	}
}

func (b *slogBench) newWithCtx(w io.Writer) logBenchmark {
	return &slogBench{
		l: newSlogWithCtx(w, slogAttrs(), b.format),
	}
}

func (b *slogBench) name() string {
	if b.format == "text" {
		return "SlogText"
	}
	return "SlogJSON"
}

func (b *slogBench) logEvent(msg string) {
	b.l.Info(msg)
}

func (b *slogBench) logEventFmt(msg string, args ...any) {
	b.l.Info(fmt.Sprintf(msg, args...))
}

func (b *slogBench) logEventCtx(msg string) {
	b.l.LogAttrs(
		context.Background(),
		slog.LevelInfo,
		msg,
		slogAttrs()...,
	)
}

func (b *slogBench) logEventCtxWeak(msg string) {
	b.l.Info(msg, alternatingKeyValuePairs()...)
}

func (b *slogBench) logDisabled(msg string) {
	b.l.Debug(msg)
}

func (b *slogBench) logDisabledFmt(msg string, args ...any) {
	b.l.Debug(fmt.Sprintf(msg, args...))
}

func (b *slogBench) logDisabledCtx(msg string) {
	b.l.LogAttrs(
		context.Background(),
		slog.LevelDebug,
		msg,
		slogAttrs()...,
	)
}

func (b *slogBench) logDisabledCtxWeak(msg string) {
	b.l.Debug(msg, alternatingKeyValuePairs()...)
}
