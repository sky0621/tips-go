package log

import (
	"fmt"
	"github.com/pkg/errors"
	"log/slog"
	"os"
	"runtime"
	"strings"
)

func SetDefaultLogger() {
	slog.SetDefault(NewLogger())
}

func NewLogger() *slog.Logger {
	return slog.New(newLogHandler())
}

func newLogHandler() *slog.JSONHandler {
	return slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level:       slog.LevelInfo,
		AddSource:   true,
		ReplaceAttr: replaceSlogAttr,
	})
}

func replaceSlogAttr(_ []string, a slog.Attr) slog.Attr {
	if a.Value.Kind() == slog.KindAny {
		if err, ok := a.Value.Any().(error); ok {
			a.Value = fmtErr(err)
		}
	}
	return a
}

func fmtErr(err error) slog.Value {
	attrs := []slog.Attr{slog.String("msg", err.Error())}

	type stackTracer interface {
		StackTrace() errors.StackTrace
	}
	var st stackTracer
	for e := err; e != nil; e = errors.Unwrap(e) {
		if tracer, ok := e.(stackTracer); ok {
			st = tracer
		}
	}
	if st != nil {
		frames := st.StackTrace()
		var lines []string
		for i := len(frames) - 1; i >= 0; i-- {
			pc := uintptr(frames[i]) - 1
			fn := runtime.FuncForPC(pc)
			if fn == nil {
				continue
			}
			file, line := fn.FileLine(pc)
			if strings.Contains(file, "runtime") {
				continue
			}
			lines = append(lines, fmt.Sprintf("%s %s:%d", fn.Name(), file, line))
		}
		attrs = append(attrs, slog.Any("trace", lines))
	}
	return slog.GroupValue(attrs...)
}
