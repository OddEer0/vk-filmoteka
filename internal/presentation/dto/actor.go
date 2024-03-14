package dto

import "time"

type (
	CreateActorDto struct {
		Name     string    `json:"name" validate:"required,min=1,max=100"`
		Gender   string    `json:"gender" validate:"required,gender"`
		Birthday time.Time `json:"birhday" validate:"required,dateIsLessNow"`
	}

	UpdateActorDto struct {
		Id       string    `json:"id" validate:"required,uuidv4"`
		Name     string    `json:"name" validate:"required,min=1,max=100"`
		Gender   string    `json:"gender" validate:"required,gender"`
		Birthday time.Time `json:"birhday" validate:"required,dateIsLessNow"`
	}
)
