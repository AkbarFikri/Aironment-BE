package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/AkbarFikri/Aironment-BE/internal/app/service"
	"github.com/AkbarFikri/Aironment-BE/internal/pkg/helper"
	"github.com/AkbarFikri/Aironment-BE/internal/pkg/model"
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
	id := ctx.Param("id")
	if id == "" {
		helper.ErrorResponse(ctx, model.ServiceResponse{
			Code:    http.StatusBadRequest,
			Error:   true,
			Message: "invalid request payload",
		})
		return
	}

	idInt, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		helper.ErrorResponse(ctx, model.ServiceResponse{
			Code:    http.StatusBadRequest,
			Error:   true,
			Message: "invalid request payload",
		})
		return
	}

	data, err := h.PostService.FetchPostByCategory(int(idInt))
	if err != nil {
		helper.ErrorResponse(ctx, data)
		return
	}

	helper.SuccessResponse(ctx, data)

}
