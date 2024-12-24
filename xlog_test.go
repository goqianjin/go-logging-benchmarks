package bench

import (
	"fmt"
	"github.com/goqianjin/common-libs/xlog"
	"io"
	"log/slog"
)

func xlogAttrs() []any {
	return []any{
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

func newXLog(w io.Writer, addReqID, addSource bool) xlog.Logger {
	l, _ := xlog.NewSLog(w, xlog.SlogOption{
		AutoReqID: &addReqID,
		AddSource: &addSource,
		//Format: xlog.FormatJSON,
	})
	return l
}

func newXLogCtx(w io.Writer, addReqID, addSource bool) xlog.Logger {
	l, _ := xlog.NewSLog(w, xlog.SlogOption{
		AutoReqID: &addReqID,
		AddSource: &addSource,
		//Format: xlog.FormatJSON,
		Args: xlogAttrs(),
	})
	return l
}

type xLogBench struct {
	l xlog.Logger

	addReqID  bool
	addSource bool
}

func (b *xLogBench) new(w io.Writer) logBenchmark {
	return &xLogBench{
		l: newXLog(w, b.addReqID, b.addSource),
	}
}

func (b *xLogBench) newWithCtx(w io.Writer) logBenchmark {
	return &xLogBench{
		l: newXLogCtx(w, b.addReqID, b.addSource),
	}
}

func (b *xLogBench) name() string {
	if b.addReqID || b.addSource {
		return "RichXLog"
	}
	return "XLog"
}

func (b *xLogBench) logEvent(msg string) {
	b.l.Info(msg)
}

func (b *xLogBench) logEventFmt(msg string, args ...any) {
	b.l.Info(fmt.Sprintf(msg, args...))
}

func (b *xLogBench) logEventCtx(msg string) {
	// (1) slogBench 通过 slog.(*Logger).LogAttr JSON 实现，SLOG 性能会好一些
	// (2) slogBench 改为 slog.(*Logger).LogAttr Text 实现时，xLog性能好一些
	// (3) slogBench 使用 slog.(*Logger).LogAttr 比 slog.(*Logger).Log 性能好
	b.l.Info(msg, xlogAttrs()...)
}

func (b *xLogBench) logEventCtxWeak(msg string) {
	// ?
	b.l.Info(msg, alternatingKeyValuePairs()...)
}

func (b *xLogBench) logDisabled(msg string) {
	b.l.Debug(msg)
}

func (b *xLogBench) logDisabledFmt(msg string, args ...any) {
	b.l.Debug(fmt.Sprintf(msg, args...))
}

func (b *xLogBench) logDisabledCtx(msg string) {
	b.l.Debug(msg, xlogAttrs()...)
}

func (b *xLogBench) logDisabledCtxWeak(msg string) {
	b.l.Debug(msg, alternatingKeyValuePairs()...)
}
