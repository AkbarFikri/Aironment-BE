package service

import (
	"net/http"
	"time"

	"github.com/google/uuid"

	"github.com/AkbarFikri/Aironment-BE/internal/app/entity"
	"github.com/AkbarFikri/Aironment-BE/internal/app/repository"
	"github.com/AkbarFikri/Aironment-BE/internal/pkg/helper"
	"github.com/AkbarFikri/Aironment-BE/internal/pkg/model"
)

type AuthService struct {
	UserRepository repository.UserRepository
}

func NewAuth(ur repository.UserRepository) AuthService {
	return AuthService{
		UserRepository: ur,
	}
}

func (s *AuthService) Register(req model.UserCreateRequest) (model.ServiceResponse, error) {
	_, err := s.UserRepository.FindByEmail(req.Email)
	if err == nil {
		return model.ServiceResponse{
			Code:    http.StatusBadRequest,
			Error:   true,
			Message: "email already used",
		}, err
	}

	hashPass, err := helper.HashPassword(req.Password)
	if err != nil {
		return model.ServiceResponse{
			Code:    http.StatusInternalServerError,
			Error:   true,
			Message: "Something went wrong",
			Payload: err.Error(),
		}, err
	}

	user := entity.User{
		ID:       uuid.New().String(),
		Email:    req.Email,
		Password: hashPass,
		FullName: req.FullName,
	}

	if err := s.UserRepository.Insert(user); err != nil {
		return model.ServiceResponse{
			Code:    http.StatusInternalServerError,
			Error:   true,
			Message: "Something went wrong, failed to create user",
			Payload: err.Error(),
		}, err
	}

	res := model.UserCreateResponse{
		ID: user.ID,
	}

	return model.ServiceResponse{
		Code:    http.StatusOK,
		Error:   false,
		Message: "successfully create user",
		Payload: res,
	}, nil
}

func (s *AuthService) Login(req model.LoginRequest) (model.ServiceResponse, error) {
	user, err := s.UserRepository.FindByEmail(req.Email)
	if err != nil {
		return model.ServiceResponse{
			Code:    http.StatusBadRequest,
			Error:   true,
			Message: "Invalid Email or Password",
			Payload: err.Error(),
		}, err
	}

	if err := helper.ComparePassword(user.Password, req.Password); err != nil {
		return model.ServiceResponse{
			Code:    http.StatusBadRequest,
			Error:   true,
			Message: "Invalid Email or Password",
			Payload: nil,
		}, err
	}

	accessData := map[string]interface{}{
		"id":    user.ID,
		"email": user.Email,
	}

	accessToken, err := helper.SignJWT(accessData, 3)
	if err != nil {
		return model.ServiceResponse{
			Code:    http.StatusInternalServerError,
			Error:   true,
			Message: "Something went wrong",
			Payload: nil,
		}, err
	}

	res := model.LoginResponse{
		Token:    accessToken,
		ExpireAt: time.Now().Add(3 * time.Hour),
	}

	return model.ServiceResponse{
		Code:    http.StatusOK,
		Error:   false,
		Message: "Successfully login",
		Payload: res,
	}, nil
}
