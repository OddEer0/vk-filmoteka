package httpv1

import "net/http"

type (
	ActorHandler interface {
		Create(res http.ResponseWriter, req *http.Request) error
		Delete(res http.ResponseWriter, req *http.Request) error
		GetByQuery(res http.ResponseWriter, req *http.Request) error
		Update(res http.ResponseWriter, req *http.Request) error
	}

	actorHandler struct{}
)

func NewActorHandler() ActorHandler {
	return actorHandler{}
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
