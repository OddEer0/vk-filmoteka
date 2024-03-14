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

// @Summary Регистрация пользователя
// @Description Ответом при успешном регистрация получаем свои данные
// @Tag auth
// @Accept json
// @Produce json
// @Success 200 {object} PureUser
// @Failure 404 {object} ErrorResponse
// @Router /http/v1/auth/registration
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
