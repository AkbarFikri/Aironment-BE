package service

import (
	"net/http"

	"github.com/AkbarFikri/Aironment-BE/internal/app/repository"
	"github.com/AkbarFikri/Aironment-BE/internal/pkg/model"
)

type PostService struct {
	PostRepository repository.PostRepository
}

func NewPost(pr repository.PostRepository) PostService {
	return PostService{
		PostRepository: pr,
	}
}

func (s *PostService) FetchPostByCategory(id int) (model.ServiceResponse, error) {
	posts, err := s.PostRepository.FindByCategory(id)
	if err != nil {
		return model.ServiceResponse{
			Code:    http.StatusInternalServerError,
			Error:   true,
			Message: "failed to found post",
		}, err
	}

	var res []model.PostResponse

	for _, v := range posts {
		dump := model.PostResponse{
			ID:          v.ID,
			Title:       v.Title,
			Description: v.Description,
			ImageUrl:    v.ImageUrl,
			Link:        v.Link,
			Category:    v.Category,
		}

		res = append(res, dump)
	}

	return model.ServiceResponse{
		Code:    http.StatusOK,
		Error:   false,
		Message: "success",
		Payload: res,
	}, nil
}
