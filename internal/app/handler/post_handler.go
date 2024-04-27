package handler

import (
	"github.com/gin-gonic/gin"

	"github.com/AkbarFikri/hackfestuc2024_backend/internal/app/service"

)

type PostHandler struct {
	PostService service.PostService
}

func NewPost(ps service.PostService) PostHandler {
	return PostHandler{
		PostService: ps,
	}
}

func (h *PostHandler) GetData(ctx *gin.Context) {
	
}