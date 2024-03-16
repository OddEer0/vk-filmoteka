package tests

import (
	userService "github.com/OddEer0/vk-filmoteka/internal/app/services/user_service"
	mockRepository "github.com/OddEer0/vk-filmoteka/internal/infrastructure/storage/mock_repository"
)

type MockServices struct {
	UserService userService.Service
}

func NewServices() *MockServices {
	userRepo := mockRepository.NewUserRepository()

	return &MockServices{
		UserService: userService.New(userRepo),
	}
}
