package caller

import (
	"runtime"

	"go.opentelemetry.io/otel/attribute"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
)

type CallerInfo struct {
	File     string
	Function string
	Line     int
}

// https://github.com/open-telemetry/opentelemetry-go-contrib/blob/main/bridges/otelslog/handler.go#L192
func GetCallerInfo(skip int) (CallerInfo, error) {
	pc, _, _, _ := runtime.Caller(skip)

	fs := runtime.CallersFrames([]uintptr{pc})
	f, _ := fs.Next()

	info := CallerInfo{
		File:     f.File,
		Function: f.Function,
		Line:     f.Line,
	}

	return info, nil
}

func (c *CallerInfo) OtelKV() []attribute.KeyValue {
	return []attribute.KeyValue{
		semconv.CodeFilepath(c.File),
		semconv.CodeFunction(c.Function),
		semconv.CodeLineNumber(c.Line),
	}
}
