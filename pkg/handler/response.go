package handler

import (
	"github.com/gin-gonic/gin"
)

type error struct {
	Message string `json:"message"`
}

// newErrorResponse - метод отправки ошибки в ответ на запрос
func newErrorResponse(ctx *gin.Context, statusCode int, message string) {
	ctx.AbortWithStatusJSON(statusCode, error{message})
}
