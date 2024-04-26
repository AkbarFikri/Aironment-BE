package service

import (
	"github.com/AkbarFikri/hackfestuc2024_backend/internal/app/repository"
	"github.com/AkbarFikri/hackfestuc2024_backend/internal/pkg/model"
)

type AuthService struct {
	UserRepository repository.UserRepository
}

func NewAuth(ur repository.UserRepository) AuthService {
	return AuthService{
		UserRepository: ur,
	}
}

func (s *AuthService) Register(req model.UserCreateRequest) (model.UserCreateResponse, error) {
	return model.UserCreateResponse{}, nil
}

func (s *AuthService) Login(req model.LoginRequest) (model.LoginResponse, error) {
	return model.LoginResponse{}, nil
}
