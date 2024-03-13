package aggregate_test

import (
	"github.com/OddEer0/ck-filmoteka/internal/domain/model"
	"github.com/OddEer0/ck-filmoteka/internal/domain/valuesobject"
	"github.com/google/uuid"
)

type InMemTestData struct {
	correctUser model.User
}

func getTestData() *InMemTestData {
	correctUserPassword, _ := valuesobject.NewPassword("correct12")

	return &InMemTestData{
		correctUser: model.User{
			Id:       uuid.New().String(),
			Name:     "Marlen",
			Password: correctUserPassword,
			Role:     "ADMIN",
		},
	}
}
