package httpv1

import (
	actorUseCase "github.com/OddEer0/vk-filmoteka/internal/app/usecases/actor_usecase"
	"net/http"
)

type (
	ActorHandler interface {
		Create(res http.ResponseWriter, req *http.Request) error
		Delete(res http.ResponseWriter, req *http.Request) error
		GetByQuery(res http.ResponseWriter, req *http.Request) error
		Update(res http.ResponseWriter, req *http.Request) error
	}

	actorHandler struct {
		actorUseCase.ActorUseCase
	}
)

func NewActorHandler(useCase actorUseCase.ActorUseCase) ActorHandler {
	return actorHandler{
		ActorUseCase: useCase,
	}
}

func (a actorHandler) Create(res http.ResponseWriter, req *http.Request) error {
	//TODO implement me
	panic("implement me")
}

func (a actorHandler) Delete(res http.ResponseWriter, req *http.Request) error {
	//TODO implement me
	panic("implement me")
}

func (a actorHandler) GetByQuery(res http.ResponseWriter, req *http.Request) error {
	//TODO implement me
	panic("implement me")
}

func (a actorHandler) Update(res http.ResponseWriter, req *http.Request) error {
	//TODO implement me
	panic("implement me")
}
