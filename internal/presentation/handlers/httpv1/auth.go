package httpv1

import (
	_ "github.com/OddEer0/ck-filmoteka/internal/presentation/dto"
	"net/http"
)

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

type ErrorHandler struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

// @Summary Регистрация пользователя
// @Description Ответом при успешном регистрация получаем свои данные
// @Tags auth
// @Accept json
// @Produce json
// @Success 200 {object} dto.RegistrationDto "Данные созданного пользователя"
// @Failure 404 {object} ErrorHandler "Ошибка 404"
// @Router /http/v1/auth/registration [post]
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
