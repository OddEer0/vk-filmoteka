package logger

import (
	"bytes"
	"log/slog"
	"net/http"
	"time"
)

type responseWriter struct {
	http.ResponseWriter
	status      int
	wroteHeader bool
	buf         *bytes.Buffer
}

func (rw *responseWriter) Write(b []byte) (int, error) {
	rw.buf.Write(b)
	return rw.ResponseWriter.Write(b)
}

func wrapResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{ResponseWriter: w, buf: bytes.NewBuffer(nil)}
}

func (rw *responseWriter) Status() int {
	return rw.status
}

func (rw *responseWriter) WriteHeader(code int) {
	if rw.wroteHeader {
		return
	}

	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
	rw.wroteHeader = true

	return
}

func LoggingMiddleware(logger *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(res http.ResponseWriter, req *http.Request) {
			start := time.Now()
			wrapped := wrapResponseWriter(res)
			next.ServeHTTP(wrapped, req)

			if wrapped.status < 400 {
				logger.Debug(
					"DEBUG",
					"status", wrapped.status,
					"method", req.Method,
					"path", req.URL.EscapedPath(),
					"duration", time.Since(start),
					"response", wrapped.buf.String(),
				)
			}
		}

		return http.HandlerFunc(fn)
	}
}
