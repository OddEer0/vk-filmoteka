package model

import "time"

type Film struct {
	Name        string    `json:"name" validate:"required,min=1,max=150"`
	Description *string   `json:"omitempty,description" validate:"max=1000"`
	CreatedAt   time.Time `json:"createdAt" validate:"required"`
	Rate        float32   `json:"rate" validate:"required,min=0,max=10"`
}
