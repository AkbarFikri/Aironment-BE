package helper

import (
	"github.com/gin-gonic/gin"

	"github.com/AkbarFikri/Aironment-BE/internal/pkg/model"
)

func SuccessResponse(c *gin.Context, data model.ServiceResponse) {
	res := &model.Response{
		Error:   data.Error,
		Message: data.Message,
		Payload: data.Payload,
	}

	c.JSON(data.Code, res)
	return
}

func ErrorResponse(c *gin.Context, data model.ServiceResponse) {
	res := &model.Response{
		Error:   data.Error,
		Message: data.Message,
		Payload: data.Payload,
	}

	c.AbortWithStatusJSON(data.Code, res)
	return
}
