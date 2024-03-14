package router

import (
	"net/http"
	"strings"
)

type AppRouter struct{}

func (r *AppRouter) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	switch {
	case strings.HasPrefix(req.URL.Path, HttpV1Prefix):
		HttpV1Router(res, req)
	default:
		http.NotFound(res, req)
	}
}
