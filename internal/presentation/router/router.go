package router

import (
	"net/http"
	"strings"
)

func NewAppRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		switch {
		case strings.HasPrefix(req.URL.Path, HttpV1Prefix):
			HttpV1Router(res, req)
		default:
			http.NotFound(res, req)
		}
	})

	return mux
}
