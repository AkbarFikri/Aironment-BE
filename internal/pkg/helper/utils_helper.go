package helper

import (
	"github.com/AkbarFikri/hackfestuc2024_backend/internal/pkg/model"
	"github.com/gin-gonic/gin"
)

func GetUserLoginData(c *gin.Context) model.UserTokenData {
	getUser, _ := c.Get("user")
	user := getUser.(model.UserTokenData)

	return user
}
