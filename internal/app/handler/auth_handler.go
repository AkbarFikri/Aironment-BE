package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/AkbarFikri/Aironment-BE/internal/app/service"
	"github.com/AkbarFikri/Aironment-BE/internal/pkg/helper"
	"github.com/AkbarFikri/Aironment-BE/internal/pkg/model"
)

type AuthHandler struct {
	AuthService service.AuthService
}

func NewAuth(as service.AuthService) AuthHandler {
	return AuthHandler{
		AuthService: as,
	}
}

func (h *AuthHandler) Register(ctx *gin.Context) {
	var req model.UserCreateRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		helper.ErrorResponse(ctx, model.ServiceResponse{
			Code:    http.StatusBadRequest,
			Error:   true,
			Message: "invalid request payload",
		})
		return
	}

	data, err := h.AuthService.Register(req)
	if err != nil {
		helper.ErrorResponse(ctx, data)
		return
	}

	helper.SuccessResponse(ctx, data)
}

func (h *AuthHandler) Login(ctx *gin.Context) {
	var req model.LoginRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		helper.ErrorResponse(ctx, model.ServiceResponse{
			Code:    http.StatusBadRequest,
			Error:   true,
			Message: "invalid request payload",
		})
		return
	}

	data, err := h.AuthService.Login(req)
	if err != nil {
		helper.ErrorResponse(ctx, data)
		return
	}

	helper.SuccessResponse(ctx, data)
}
