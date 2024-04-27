package service

import (
	"net/http"
	"strconv"

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

func (s *UserService) FetchAirQualitysPoints(req model.AqiParam) (model.ServiceResponse, error) {
	latitude, err := strconv.ParseFloat(req.Latitude, 64)
	if err != nil {
		return model.ServiceResponse{
			Code:    http.StatusBadRequest,
			Error:   true,
			Message: "invalid latitude format",
			Payload: latitude,
		}, err
	}

	longitude, err := strconv.ParseFloat(req.Longitude, 64)
	if err != nil {
		return model.ServiceResponse{
			Code:    http.StatusBadRequest,
			Error:   true,
			Message: "invalid longitude format",
			Payload: longitude,
		}, err
	}

	return model.ServiceResponse{}, nil
}
