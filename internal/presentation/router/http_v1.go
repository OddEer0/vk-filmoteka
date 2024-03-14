package router

import (
	"github.com/OddEer0/ck-filmoteka/internal/presentation/handlers/httpv1"
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
	appHandler := httpv1.NewAppHandler()

	switch {
	case req.Method == http.MethodPost && path == "/registration":
		appHandler.AuthHandler.Registration(res, req)
	case req.Method == http.MethodPost && path == "/login":
		appHandler.AuthHandler.Login(res, req)
	case req.Method == http.MethodPost && path == "/refresh":
		appHandler.AuthHandler.Refresh(res, req)
	case req.Method == http.MethodGet && path == "/logout":
		appHandler.AuthHandler.Logout(res, req)
	default:
		http.NotFound(res, req)
	}
}

func HttpV1RouterActor(res http.ResponseWriter, req *http.Request) {
	appHandler := httpv1.NewAppHandler()

	switch req.Method {
	case http.MethodGet:
		appHandler.ActorHandler.GetByQuery(res, req)
	case http.MethodPost:
		appHandler.ActorHandler.Create(res, req)
	case http.MethodPut:
		appHandler.ActorHandler.Update(res, req)
	case http.MethodDelete:
		appHandler.ActorHandler.Delete(res, req)
	default:
		http.NotFound(res, req)
	}
}

func HttpV1RouterFilm(res http.ResponseWriter, req *http.Request) {
	appHandler := httpv1.NewAppHandler()

	switch req.Method {
	case http.MethodGet:
		appHandler.FilmHandler.GetByQuery(res, req)
	case http.MethodPost:
		appHandler.FilmHandler.Create(res, req)
	case http.MethodPut:
		appHandler.FilmHandler.Update(res, req)
	case http.MethodDelete:
		appHandler.FilmHandler.Delete(res, req)
	default:
		http.NotFound(res, req)
	}
}
