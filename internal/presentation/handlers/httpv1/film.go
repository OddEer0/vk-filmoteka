package httpv1

import (
	filmUseCase "github.com/OddEer0/vk-filmoteka/internal/app/usecases/film_usecase"
	"net/http"
)

type (
	FilmHandler interface {
		Create(res http.ResponseWriter, req *http.Request) error
		Delete(res http.ResponseWriter, req *http.Request) error
		GetByQuery(res http.ResponseWriter, req *http.Request) error
		Update(res http.ResponseWriter, req *http.Request) error
		SearchByNameAndActorName(res http.ResponseWriter, req *http.Request) error
	}

	filmHandler struct {
		filmUseCase.FilmUseCase
	}
)

func NewFilmHandler(useCase filmUseCase.FilmUseCase) FilmHandler {
	return filmHandler{
		FilmUseCase: useCase,
	}
}

func (a filmHandler) Create(res http.ResponseWriter, req *http.Request) error {
	//TODO implement me
	panic("implement me")
}

func (a filmHandler) Delete(res http.ResponseWriter, req *http.Request) error {
	//TODO implement me
	panic("implement me")
}

func (a filmHandler) GetByQuery(res http.ResponseWriter, req *http.Request) error {
	//TODO implement me
	panic("implement me")
}

func (a filmHandler) Update(res http.ResponseWriter, req *http.Request) error {
	//TODO implement me
	panic("implement me")
}

func (a filmHandler) SearchByNameAndActorName(res http.ResponseWriter, req *http.Request) error {
	//TODO implement me
	panic("implement me")
}
