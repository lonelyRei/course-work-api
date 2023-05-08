package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authHeader = "Authorization"
	tokenType  = "Bearer"
	userCtx    = "userId"
)

// userIdentity - middleware для проверки авторизованного пользователя
func (h *Handler) userIdentity(ctx *gin.Context) {
	// Получаем заголовок Authorization
	header := ctx.GetHeader(authHeader)
	if header == "" {
		newErrorResponse(ctx, http.StatusUnauthorized, "empty auth header")
		return
	}

	// Разбиваем заголовок на две части
	headerParts := strings.Split(header, " ")

	if len(headerParts) != 2 {
		newErrorResponse(ctx, http.StatusUnauthorized, "invalid auth header")
		return
	}

	if headerParts[0] != tokenType {
		newErrorResponse(ctx, http.StatusUnauthorized, "wrong token type")
		return
	}

	// Парсим токен из заголовка
	userId, err := h.services.ParseToken(headerParts[1])

	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	// Записываем в контекст идентификатор пользователя
	ctx.Set(userCtx, userId)
}

// getUserId - метод для получения идентификатора пользователя из контекста
func getUserId(c *gin.Context) (int, *error) {
	// Получаем идентификатор пользователя
	id, ok := c.Get(userCtx)

	if !ok {
		return 0, &error{"user id not found"}
	}

	// Преобразуем к целому числу
	idInt, ok := id.(int)
	if !ok {
		return 0, &error{"user id is of invalid type"}
	}

	return idInt, nil
}
