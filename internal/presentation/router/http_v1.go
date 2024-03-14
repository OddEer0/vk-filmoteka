package router

import (
	"net/http"
	"strings"
)

const (
	HttpV1Prefix = "/http/v1"
)

func HttpV1Router(res http.ResponseWriter, req *http.Request) {
	path := strings.TrimPrefix(req.URL.Path, HttpV1Prefix)

	switch path {
	case "/auth":
		HttpV1RouterAuth(res, req)
	case "/actor":
		HttpV1RouterActor(res, req)
	case "/film":
		HttpV1RouterFilm(res, req)
	case "/search/film":
	default:
		http.NotFound(res, req)
	}
}

func HttpV1RouterAuth(res http.ResponseWriter, req *http.Request) {
	path := strings.TrimPrefix(req.URL.Path, HttpV1Prefix+"/auth")
	switch {
	case req.Method == http.MethodPost && path == "/registration":
	case req.Method == http.MethodPost && path == "/login":
	case req.Method == http.MethodPost && path == "/refresh":
	case req.Method == http.MethodGet && path == "/logout":
	default:
		http.NotFound(res, req)
	}
}

func HttpV1RouterActor(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
	case http.MethodPost:
	case http.MethodPut:
	case http.MethodDelete:
	default:
		http.NotFound(res, req)
	}
}

func HttpV1RouterFilm(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
	case http.MethodPost:
	case http.MethodPut:
	case http.MethodDelete:
	default:
		http.NotFound(res, req)
	}
}
