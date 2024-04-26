package handler

import (
	"github.com/gin-gonic/gin"

	"github.com/AkbarFikri/hackfestuc2024_backend/internal/app/service"

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
	ctx.JSON(200, gin.H{"status" : "yep"})
}

func (h *AuthHandler) Login(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"status" : "yep"})
}
