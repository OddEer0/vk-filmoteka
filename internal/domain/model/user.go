package model

import "github.com/OddEer0/ck-filmoteka/internal/domain/valuesobject"

type User struct {
	Name     string                `json:"name" validate:"required,"`
	Password valuesobject.Password `json:"password" validate:"required"`
	Role     string                `json:"role" validate:"required,userRole"`
}
