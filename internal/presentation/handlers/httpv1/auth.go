package httpv1

import "net/http"

type (
	AuthHandler interface {
		Registration(res http.ResponseWriter, req *http.Request) error
		Login(res http.ResponseWriter, req *http.Request) error
		Logout(res http.ResponseWriter, req *http.Request) error
		Refresh(res http.ResponseWriter, req *http.Request) error
	}

	authHandler struct{}
)

func NewAuthHandler() AuthHandler {
	return authHandler{}
}

func (a authHandler) Registration(res http.ResponseWriter, req *http.Request) error {
	//TODO implement me
	panic("implement me")
}

func (a authHandler) Login(res http.ResponseWriter, req *http.Request) error {
	//TODO implement me
	panic("implement me")
}

func (a authHandler) Logout(res http.ResponseWriter, req *http.Request) error {
	//TODO implement me
	panic("implement me")
}

func (a authHandler) Refresh(res http.ResponseWriter, req *http.Request) error {
	//TODO implement me
	panic("implement me")
}
