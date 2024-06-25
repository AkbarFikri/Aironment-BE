package service

import (
	"errors"
	"net/http"

	"github.com/google/uuid"

	"github.com/AkbarFikri/Aironment-BE/internal/app/entity"
	"github.com/AkbarFikri/Aironment-BE/internal/app/repository"
	"github.com/AkbarFikri/Aironment-BE/internal/pkg/model"
)

type CommunityService struct {
	CommunityRepository       repository.CommunityRepository
	CommunityMemberRepository repository.CommunityMemberRepository
	UserRepository            repository.UserRepository
	PostRepository            repository.PostRepository
}

func NewCommunity(cr repository.CommunityRepository, cmr repository.CommunityMemberRepository,
	ur repository.UserRepository, pr repository.PostRepository) CommunityService {
	return CommunityService{
		CommunityMemberRepository: cmr,
		UserRepository:            ur,
		PostRepository:            pr,
		CommunityRepository:       cr,
	}
}

func (s *CommunityService) VerifiedCommunity(invoiceId string) (model.ServiceResponse, error) {
	community, err := s.CommunityRepository.FindByInvoiceId(invoiceId)
	if err != nil {
		return model.ServiceResponse{
			Code:    http.StatusBadRequest,
			Error:   true,
			Message: "community with id provided not found",
		}, err
	}

	community.Status = "verified"

	if err := s.CommunityRepository.Update(community); err != nil {
		return model.ServiceResponse{
			Code:    http.StatusInternalServerError,
			Error:   true,
			Message: "failed to update community information",
		}, err
	}

	return model.ServiceResponse{
		Code:    http.StatusOK,
		Error:   false,
		Message: "successfully verified community",
	}, nil
}

func (s *CommunityService) FetchCommunity() (model.ServiceResponse, error) {
	communitys, err := s.CommunityRepository.FindAll()
	if err != nil {
		return model.ServiceResponse{
			Code:    http.StatusBadRequest,
			Error:   true,
			Message: "community not found",
		}, err
	}

	if len(communitys) == 0 {
		return model.ServiceResponse{
			Code:    http.StatusNotFound,
			Error:   true,
			Message: "community not found",
			Payload: communitys,
		}, errors.New("no data found")
	}

	var res []model.CommunityDataResponse

	for _, c := range communitys {
		dump := model.CommunityDataResponse{
			ID:             c.ID,
			Name:           c.Name,
			Description:    c.Description,
			ProfilePicture: c.ProfilePicture,
			CoverPicture:   c.CoverPicture,
		}

		res = append(res, dump)
	}

	return model.ServiceResponse{
		Code:    http.StatusOK,
		Error:   false,
		Message: "successfully find all comunitys",
		Payload: res,
	}, nil
}

func (s *CommunityService) FetchCommunityDetails(id string) (model.ServiceResponse, error) {
	community, err := s.CommunityRepository.FindById(id)
	if err != nil {
		return model.ServiceResponse{
			Code:    http.StatusBadRequest,
			Error:   true,
			Message: "something went wrong, community not found",
		}, err
	}

	posts, err := s.PostRepository.FindAllImageByCommunity(community.ID)
	if err != nil {
		return model.ServiceResponse{
			Code:    http.StatusBadRequest,
			Error:   true,
			Message: "something went wrong, community galery not found",
		}, err
	}

	var postImageUrl []string

	for _, c := range posts {
		postImageUrl = append(postImageUrl, c.ImageUrl)
	}

	members, err := s.CommunityMemberRepository.FindByCommunityId(community.ID)
	if err != nil {
		return model.ServiceResponse{
			Code:    http.StatusInternalServerError,
			Error:   true,
			Message: "something went wrong, members not found",
		}, err
	}

	ids := []string{}

	for _, c := range members {
		ids = append(ids, c.UserID)
	}

	users, err := s.UserRepository.FindManyById(ids)
	if err != nil {
		return model.ServiceResponse{
			Code:    http.StatusInternalServerError,
			Error:   true,
			Message: "something went wrong, members not found",
		}, err
	}

	var membersComm []model.UserProfileResponse

	for _, c := range users {
		dump := model.UserProfileResponse{
			ID:       c.ID,
			Email:    c.Email,
			FullName: c.FullName,
		}
		membersComm = append(membersComm, dump)
	}

	res := model.CommunityDetails{
		Detail: model.Details{
			Description: community.Description,
			Galery:      postImageUrl,
		},
		Member: membersComm,
	}

	return model.ServiceResponse{
		Code:    http.StatusOK,
		Error:   false,
		Message: "successfully find all comunitys",
		Payload: res,
	}, nil
}

func (s *CommunityService) JoinCommunity(user model.UserTokenData, CommID string) (model.ServiceResponse, error) {
	Community, err := s.CommunityRepository.FindById(CommID)
	if err != nil {
		return model.ServiceResponse{
			Code:    http.StatusBadRequest,
			Error:   true,
			Message: "community not foud",
		}, err
	}

	member := entity.Member{
		ID:          uuid.NewString(),
		UserID:      user.ID,
		CommunityID: Community.ID,
	}

	if err := s.CommunityMemberRepository.Insert(member); err != nil {
		return model.ServiceResponse{
			Code:    http.StatusInternalServerError,
			Error:   true,
			Message: "failed to save member to database",
		}, err
	}

	res := model.Member{
		ID:          member.ID,
		UserID:      member.UserID,
		CommunityID: member.CommunityID,
	}

	return model.ServiceResponse{
		Code:    http.StatusOK,
		Error:   false,
		Message: "success",
		Payload: res,
	}, nil
}
