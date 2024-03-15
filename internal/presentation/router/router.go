package router

import (
	appErrors "github.com/OddEer0/ck-filmoteka/internal/common/lib/app_errors"
	"log/slog"
	"net/http"
	"strings"
)

func appRouterHandler(res http.ResponseWriter, req *http.Request) error {
	switch {
	case strings.HasPrefix(req.URL.Path, HttpV1Prefix):
		return HttpV1Router(res, req)
	default:
		http.NotFound(res, req)
	}
	return nil
}

func NewAppRouter(log *slog.Logger) *http.ServeMux {
	mux := http.NewServeMux()
	middleware := appErrors.LoggingMiddleware(log)

	mux.HandleFunc("/", middleware(appRouterHandler))

	return mux
}
