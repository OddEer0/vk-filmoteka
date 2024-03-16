package tests

import (
	tokenService "github.com/OddEer0/vk-filmoteka/internal/app/services/token_service"
	userService "github.com/OddEer0/vk-filmoteka/internal/app/services/user_service"
	authUseCase "github.com/OddEer0/vk-filmoteka/internal/app/usecases/auth_usecase"
	mockRepository "github.com/OddEer0/vk-filmoteka/internal/infrastructure/storage/mock_repository"
)

type MockUseCases struct {
	AuthUseCase authUseCase.AuthUseCase
}

func NewUseCases() *MockUseCases {
	userRepo := mockRepository.NewUserRepository()
	tokenRepo := mockRepository.NewTokenRepository()

	return &MockUseCases{
		AuthUseCase: authUseCase.New(userService.New(userRepo), tokenService.New(tokenRepo), userRepo),
	}
}
