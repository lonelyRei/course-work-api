package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/lonelyRei/course-work/entities"
	"net/http"
	"strconv"
)

// GetUserData - обработчик получения данных о пользователе
func (h *Handler) GetUserData(ctx *gin.Context) {
	// Получаем идентификатор пользователя из контекста
	userId, err := getUserId(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Message)
		return
	}

	// Вызываем метод сервиса для получения данных о пользователе
	user, er := h.services.GetUserData(userId)
	if er != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, er.Error())
		return
	}

	// Отдаем ответ со статусом 200
	ctx.JSON(http.StatusOK, map[string]interface{}{"user": user})
}

// SetUserData - обработчик изменения пользовательских данных
func (h *Handler) SetUserData(ctx *gin.Context) {
	// Получаем идентификатор пользователя из контекста
	userId, err := getUserId(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Message)
		return
	}

	// Создаем переменную для парсинга тела запроса
	var input entities.UpdateUserInput

	// читаем тело запроса
	if err := ctx.BindJSON(&input); err != nil {
		// Если не удалось прочитать тело запроса, то генерируем ошибку
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	// Вызываем метод сервиса для получения данных о пользователе
	er := h.services.SetUserData(userId, input)

	if er != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, er.Error())
		return
	}

	// Отдаем ответ со статусом 200
	ctx.JSON(http.StatusOK, map[string]interface{}{})
}

// GetUserById - метод для получения подробной информации о пользователе по идентификатору
func (h *Handler) GetUserById(ctx *gin.Context) {
	// Получаем query параметры
	userId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "invalid id param")
		return
	}

	// Вызываем метод сервиса для получения данных о пользователе
	user, er := h.services.GetUserData(userId)
	if er != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, er.Error())
		return
	}

	// Отдаем ответ со статусом 200
	ctx.JSON(http.StatusOK, map[string]interface{}{"user": user})
}
