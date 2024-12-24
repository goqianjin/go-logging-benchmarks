package bench

import (
	"os"
	"testing"
)

func TestEvent(t *testing.T) {
	t.Logf("Log a simple message without any contexual fields")

	for _, v := range loggers {
		l := v.new(os.Stdout)
		l.logEvent(logMsg)
	}
}

func TestEventFmt(t *testing.T) {
	t.Logf("Log a simple message using string formatting verbs")
	for _, v := range loggers {
		l := v.new(os.Stdout)
		l.logEventFmt(logMsgFmt, logMsgArgs...)
	}
}

func TestEventCtx(t *testing.T) {
	t.Logf("Log an event with several contextual fields")

	for _, v := range loggers {
		l := v.new(os.Stdout)
		l.logEventCtx(logMsg)
	}
}

func TestEventCtxWeak(t *testing.T) {
	t.Logf("Log an event with weakly typed contextual fields")

	for _, v := range loggers {
		l := v.newWithCtx(os.Stdout)
		l.logEventCtxWeak(logMsg)
	}
}

func TestEventAccumulatedCtx(t *testing.T) {
	t.Logf("Log an event with some accumulated context")

	for _, v := range loggers {
		l := v.newWithCtx(os.Stdout)
		l.logEvent(logMsg)
	}
}
