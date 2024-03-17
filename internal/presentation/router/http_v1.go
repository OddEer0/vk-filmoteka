package router

import (
	"net/http"
	"strings"

	"github.com/OddEer0/vk-filmoteka/internal/presentation/handlers/httpv1"
)

const (
	HttpV1Prefix = "/http/v1"
)

func HttpV1Router(res http.ResponseWriter, req *http.Request) error {
	path := strings.TrimPrefix(req.URL.Path, HttpV1Prefix)

	switch {
	case strings.HasPrefix(path, "/auth"):
		return HttpV1RouterAuth(res, req)
	case strings.HasPrefix(path, "/actor"):
		return HttpV1RouterActor(res, req)
	case strings.HasPrefix(path, "/film"):
		return HttpV1RouterFilm(res, req)
	case strings.HasPrefix(path, "/search/film"):
	default:
		http.NotFound(res, req)
	}
	return nil
}

func HttpV1RouterAuth(res http.ResponseWriter, req *http.Request) error {
	path := strings.TrimPrefix(req.URL.Path, HttpV1Prefix+"/auth")
	appHandler := httpv1.NewAppHandler()

	switch {
	case req.Method == http.MethodPost && path == "/registration":
		return appHandler.AuthHandler.Registration(res, req)
	case req.Method == http.MethodPost && path == "/login":
		return appHandler.AuthHandler.Login(res, req)
	case req.Method == http.MethodPost && path == "/refresh":
		return appHandler.AuthHandler.Refresh(res, req)
	case req.Method == http.MethodGet && path == "/logout":
		return appHandler.AuthHandler.Logout(res, req)
	default:
		http.NotFound(res, req)
	}
	return nil
}

func HttpV1RouterActor(res http.ResponseWriter, req *http.Request) error {
	appHandler := httpv1.NewAppHandler()

	switch req.Method {
	case http.MethodGet:
		return appHandler.ActorHandler.GetByQuery(res, req)
	case http.MethodPost:
		return appHandler.ActorHandler.Create(res, req)
	case http.MethodPut:
		return appHandler.ActorHandler.Update(res, req)
	case http.MethodDelete:
		return appHandler.ActorHandler.Delete(res, req)
	default:
		http.NotFound(res, req)
	}
	return nil
}

func HttpV1RouterFilm(res http.ResponseWriter, req *http.Request) error {
	appHandler := httpv1.NewAppHandler()

	switch req.Method {
	case http.MethodGet:
		return appHandler.FilmHandler.GetByQuery(res, req)
	case http.MethodPost:
		return appHandler.FilmHandler.Create(res, req)
	case http.MethodPut:
		return appHandler.FilmHandler.Update(res, req)
	case http.MethodDelete:
		return appHandler.FilmHandler.Delete(res, req)
	default:
		http.NotFound(res, req)
	}
	return nil
}
