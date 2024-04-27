package service

import (
	"net/http"

	"github.com/AkbarFikri/hackfestuc2024_backend/internal/app/repository"
	"github.com/AkbarFikri/hackfestuc2024_backend/internal/pkg/model"
)

type UserService struct {
	UserRepository repository.UserRepository
}

func NewUser(ur repository.UserRepository) UserService {
	return UserService{
		UserRepository: ur,
	}
}

func (s *UserService) Current(req model.UserTokenData) (model.ServiceResponse, error) {
	user, err := s.UserRepository.FindById(req.ID)
	if err != nil {
		return model.ServiceResponse{
			Code:    http.StatusBadRequest,
			Error:   true,
			Message: "invalid id provided, user not found",
		}, err
	}

	res := model.UserProfileResponse{
		ID:       user.ID,
		Email:    user.Email,
		FullName: user.FullName,
	}

	return model.ServiceResponse{
		Code:    http.StatusOK,
		Error:   false,
		Message: "succesfully find user",
		Payload: res,
	}, nil
}
