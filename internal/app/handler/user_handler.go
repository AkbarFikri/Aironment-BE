package handler

import (
	"github.com/gin-gonic/gin"

	"github.com/AkbarFikri/hackfestuc2024_backend/internal/app/service"
	"github.com/AkbarFikri/hackfestuc2024_backend/internal/pkg/helper"

)

type UserHandler struct {
	UserService service.UserService
}

func NewUser(us service.UserService) UserHandler {
	return UserHandler{
		UserService: us,
	}
}

func (h *UserHandler) CurrentUser(ctx *gin.Context) {
	user := helper.GetUserLoginData(ctx)

	data, err := h.UserService.Current(user)
	if err != nil {
		helper.ErrorResponse(ctx, data)
		return
	}

	helper.SuccessResponse(ctx, data)
}
