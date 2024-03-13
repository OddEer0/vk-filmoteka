package model

import "time"

type Film struct {
	Id          string    `json:"id" validate:"required,uuidv4"`
	Name        string    `json:"name" validate:"required,min=1,max=150"`
	Description *string   `json:"omitempty,description" validate:"omitempty,max=1000"`
	ReleaseDate time.Time `json:"createdAt" validate:"required"`
	Rate        float32   `json:"rate" validate:"min=0,max=10"`
}
