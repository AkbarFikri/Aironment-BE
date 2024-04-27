package handler

import (
	"github.com/gin-gonic/gin"

	"github.com/AkbarFikri/hackfestuc2024_backend/internal/app/service"
)

type PaymentHandler struct {
	CommunityService service.CommunityService
	PaymentService   service.PaymentService
}

func NewPayment(cs service.CommunityService,
	ps service.PaymentService) PaymentHandler {
	return PaymentHandler{
		CommunityService: cs,
		PaymentService:   ps,
	}
}

func (h *PaymentHandler) Verify(ctx *gin.Context) {
	var notificationPayload map[string]interface{}

	err := ctx.ShouldBind(&notificationPayload)
	if err != nil {
		return
	}

	orderId, exists := notificationPayload["order_id"].(string)
	if !exists {
		return
	}

	success := h.PaymentService.VerifyPayment(orderId)
	if !success {
		return
	}
	h.CommunityService.VerifiedCommunity(orderId)
}
