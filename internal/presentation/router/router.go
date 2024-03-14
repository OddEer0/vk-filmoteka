package router

import (
	"net/http"
	"strings"
)

func AppRouter(res http.ResponseWriter, req *http.Request) {
	switch {
	case strings.HasPrefix(req.URL.Path, HttpV1Prefix):
		HttpV1Router(res, req)
	default:
		http.NotFound(res, req)
	}
}
