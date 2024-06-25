package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/AkbarFikri/Aironment-BE/internal/app/service"
	"github.com/AkbarFikri/Aironment-BE/internal/pkg/helper"
	"github.com/AkbarFikri/Aironment-BE/internal/pkg/model"
)

type CommunityHandler struct {
	CommunityService service.CommunityService
	PaymentService   service.PaymentService
}

func NewCommunity(cs service.CommunityService, ps service.PaymentService) CommunityHandler {
	return CommunityHandler{
		CommunityService: cs,
		PaymentService:   ps,
	}
}

func (h *CommunityHandler) CreateCommunity(ctx *gin.Context) {
	user := helper.GetUserLoginData(ctx)

	var req model.CommunityRequest

	if err := ctx.ShouldBind(&req); err != nil {
		helper.ErrorResponse(ctx, model.ServiceResponse{
			Code:    http.StatusBadRequest,
			Error:   true,
			Message: "invalid request payload",
		})
		return
	}

	data, err := h.PaymentService.GenerateUrlToken(user, req)
	if err != nil {
		helper.ErrorResponse(ctx, data)
		return
	}

	helper.SuccessResponse(ctx, data)
}

func (h *CommunityHandler) GetCommunities(ctx *gin.Context) {
	data, err := h.CommunityService.FetchCommunity()
	if err != nil {
		helper.ErrorResponse(ctx, data)
		return
	}

	helper.SuccessResponse(ctx, data)
}

func (h *CommunityHandler) GetCommunityDetails(ctx *gin.Context) {
	CommID := ctx.Param("id")
	if CommID == "" {
		helper.ErrorResponse(ctx, model.ServiceResponse{
			Code:    http.StatusBadRequest,
			Error:   true,
			Message: "invalid request payload",
		})
		return
	}

	data, err := h.CommunityService.FetchCommunityDetails(CommID)
	if err != nil {
		helper.ErrorResponse(ctx, data)
		return
	}

	helper.SuccessResponse(ctx, data)
}

func (h *CommunityHandler) JoinCommunity(ctx *gin.Context) {
	user := helper.GetUserLoginData(ctx)
	var req model.MemberRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		helper.ErrorResponse(ctx, model.ServiceResponse{
			Code:    http.StatusBadRequest,
			Error:   true,
			Message: "invalid request payload",
		})
		return
	}

	data, err := h.CommunityService.JoinCommunity(user, req.CommunityID)
	if err != nil {
		helper.ErrorResponse(ctx, data)
		return
	}

	helper.SuccessResponse(ctx, data)
}
