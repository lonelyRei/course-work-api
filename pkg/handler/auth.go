package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/lonelyRei/course-work/entities"
	"net/http"
)

/*
	Пакет, содерщащий обработчики для авторизации и регистрации пользователей
*/

// signUp - обработчик регистрации нового пользователя
func (h *Handler) signUp(ctx *gin.Context) {
	// Объявляем переменную, в которую будем парсить Json
	var input entities.User

	// читаем тело запроса
	if err := ctx.BindJSON(&input); err != nil {
		// Если не удалось прочитать тело запроса, то генерируем ошибку
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	// Передаем данные на сервис
	id, err := h.services.Authorization.CreateUser(input)

	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{"id": id})
}

// signIn - обработчик авторизации пользователя
func (h *Handler) signIn(ctx *gin.Context) {
	// Объявляем переменную, в которую будем парсить Json
	var input authUserInput

	// читаем тело запроса
	if err := ctx.BindJSON(&input); err != nil {
		// Если не удалось прочитать тело запроса, то генерируем ошибку
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	// генерируем токен
	token, err := h.services.GenerateToken(input.Username, input.Password)

	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	// Получаем пользователя
	user, err := h.services.GetUser(input.Username, input.Password)

	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{"token": token,
		"user": authUserOutput{Username: user.Username, Name: user.Name, Id: user.Id}})
}
