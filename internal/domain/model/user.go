package model

import "github.com/OddEer0/ck-filmoteka/internal/domain/valuesobject"

type User struct {
	Id       string                `json:"id" validate:"required,uuidv4"`
	Name     string                `json:"name" validate:"required,min=3,max=100"`
	Password valuesobject.Password `json:"password" validate:"required"`
	Role     string                `json:"role" validate:"required,userRole"`
}
