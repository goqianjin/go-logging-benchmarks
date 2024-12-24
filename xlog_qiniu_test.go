package bench

import (
	"encoding/json"
	"fmt"
	"io"
	"qiniu.com/kodo/libs/xlog.v1"
)

func qxLogAttrs() []any {
	return []any{
		fmt.Sprintf("bytes=%v", ctxBodyBytes),
		fmt.Sprintf("request=%v", ctxRequest),
		fmt.Sprintf("elapsed_time_ms=%v", ctxTimeElapsedMs),
		fmt.Sprintf("user=%v", ctxUser),
		fmt.Sprintf("now=%v", ctxTime),
		fmt.Sprintf("months=%v", ctxMonths),
		fmt.Sprintf("primes=%v", ctxFirst10Primes),
		fmt.Sprintf("users=%v", ctxUsers),
		fmt.Sprintf("error=%v", ctxErr),
	}
}

func newQXLog(w io.Writer) *xlog.Logger {
	xlog.SetOutput(w)
	return xlog.NewDummy()
}

func newQXLogCtx(w io.Writer) *xlog.Logger {
	xlog.SetOutput(w)
	ctxGloabalInfo, _ := json.Marshal(qxLogAttrs())
	return xlog.NewWith(string(ctxGloabalInfo))
}

type qxLogBench struct {
	l *xlog.Logger
}

func (b *qxLogBench) new(w io.Writer) logBenchmark {
	return &qxLogBench{
		l: newQXLog(w),
	}
}

func (b *qxLogBench) newWithCtx(w io.Writer) logBenchmark {
	return &qxLogBench{
		l: newQXLogCtx(w),
	}
}

func (b *qxLogBench) name() string {
	return "QXLog"
}

func (b *qxLogBench) logEvent(msg string) {
	b.l.Info(msg)
}

func (b *qxLogBench) logEventFmt(msg string, args ...any) {
	b.l.Infof(msg, args...)
}

func (b *qxLogBench) logEventCtx(msg string) {
	b.l.Infof(msg, qxLogAttrs()...)
}

func (b *qxLogBench) logEventCtxWeak(msg string) {
	b.l.Infof(msg, alternatingKeyValuePairs()...)
}

func (b *qxLogBench) logDisabled(msg string) {
	b.l.Debug(msg)
}

func (b *qxLogBench) logDisabledFmt(msg string, args ...any) {
	b.l.Debugf(msg, args...)
}

func (b *qxLogBench) logDisabledCtx(msg string) {
	b.l.Debugf(msg, qxLogAttrs()...)
}

func (b *qxLogBench) logDisabledCtxWeak(msg string) {
	b.l.Debugf(msg, alternatingKeyValuePairs()...)
}
