package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/lonelyRei/course-work/entities"
	"net/http"
	"strconv"
)

// CreatePost - метод создания записи
func (h *Handler) CreatePost(ctx *gin.Context) {
	// Получаем id пользователя (проверяем авторизованность)
	userId, err := getUserId(ctx)
	if err != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, err.Message)
		return
	}

	// Читаем тело запроса
	var input entities.Post
	if err := ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	// Обращаемся к уровню сервиса для создания записи
	postId, er := h.services.Post.CreatePost(userId, input)
	if er != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, er.Error())
		return
	}

	// Возвращаем 200 код и json с id списка
	ctx.JSON(http.StatusOK, map[string]interface{}{"id": postId})
}

// GetPostById - метод получения записи по идентификатору
func (h *Handler) GetPostById(ctx *gin.Context) {
	// Получаем query параметры
	postId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "invalid id param")
		return
	}

	// Передаем запрос на уровень сервиса
	post, err := h.services.Post.GetPostById(postId)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	// Возвращаем ответ
	ctx.JSON(http.StatusOK, post)
}

// getAllPostsResponse - структура ответа на запрос получения всех пользователей
type getAllPostsResponse struct {
	Data []entities.Post `json:"data"`
}

// GetAllPosts - метод получения всех записей
func (h *Handler) GetAllPosts(ctx *gin.Context) {
	// Передаем запрос на уровень сервиса
	posts, err := h.services.Post.GetAllPosts()
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	// Возвращаем ответ
	ctx.JSON(http.StatusOK, getAllPostsResponse{Data: posts})
}

// UpdatePost - метод изменения записи
func (h *Handler) UpdatePost(ctx *gin.Context) {
	// Получаем идентификатор пользователя
	userId, er := getUserId(ctx)
	if er != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, er.Message)
		return
	}

	// Получаем query параметры
	postId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "invalid id param")
		return
	}

	// Получаем тело запроса
	var input entities.UpdatePostInput
	if err := ctx.BindJSON(&input); err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	// Вызываем метод сервиса
	if err := h.services.Post.UpdatePost(userId, postId, input); err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{})
}

// DeletePost - метод удаления записи
func (h *Handler) DeletePost(ctx *gin.Context) {
	// Получаем идентификатор пользователя
	userId, er := getUserId(ctx)
	if er != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, er.Message)
		return
	}

	// Получаем query параметры
	postId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "invalid id param")
		return
	}

	// Передаем запрос на уровень сервиса
	err = h.services.Post.DeletePost(userId, postId)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{})
}

// SetRating - метод для изменения рейтинга записи
func (h *Handler) SetRating(ctx *gin.Context) {
	// Проверяем пользователя на авторизованность
	_, er := getUserId(ctx)
	if er != nil {
		newErrorResponse(ctx, http.StatusUnauthorized, er.Message)
		return
	}

	// Получаем query параметры:
	// Получаем идентификатор записи
	postId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "invalid id param")
		return
	}

	// Получаем тип действия (increase или decrease)
	action := ctx.Param("action")
	if action != entities.IncreaseRatingAction && action != entities.DecreaseRatingAction {
		newErrorResponse(ctx, http.StatusBadRequest, "invalid action param")
		return
	}

	// Передаем запрос на уровень сервиса
	rating, err := h.services.Post.SetRating(postId, action)

	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{"rating": rating})
}

// GetUserPosts - метод получения всех постов конкретного пользователя
func (h *Handler) GetUserPosts(ctx *gin.Context) {
	// Получаем query параметры:
	// Получаем идентификатор пользователя
	userId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		newErrorResponse(ctx, http.StatusBadRequest, "invalid id param")
		return
	}

	posts, err := h.services.Post.GetUserPosts(userId)
	if err != nil {
		newErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	// Возвращаем ответ
	ctx.JSON(http.StatusOK, getAllPostsResponse{Data: posts})
}
