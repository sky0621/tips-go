package main

import (
	"fmt"
	"github.com/pkg/errors"
	"log/slog"
	"os"
	"runtime"
)

func main() {
	handler := slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
		ReplaceAttr: replaceAttr,
	})
	slog.SetDefault(slog.New(handler))

	if err := f(); err != nil {
		slog.Info("エラー発生", slog.Any("err", err))
	}
}

func f() error { return g() }
func g() error { return errors.New("テストエラー") }

func replaceAttr(groups []string, a slog.Attr) slog.Attr {
	if a.Value.Kind() == slog.KindAny {
		if err, ok := a.Value.Any().(error); ok {
			a.Value = fmtErr(err)
		}
	}
	return a
}

func fmtErr(err error) slog.Value {
	// エラーメッセージ
	attrs := []slog.Attr{slog.String("msg", err.Error())}

	// pkg/errors の StackTrace を取り出す
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
		// runtime.Frame から文字列化
		var lines []string
		for i := len(frames) - 1; i >= 0; i-- {
			pc := uintptr(frames[i]) - 1
			fn := runtime.FuncForPC(pc)
			if fn == nil {
				continue
			}
			file, line := fn.FileLine(pc)
			lines = append(lines, fmt.Sprintf("%s %s:%d", fn.Name(), file, line))
		}
		attrs = append(attrs, slog.Any("trace", lines))
	}
	return slog.GroupValue(attrs...)
}
