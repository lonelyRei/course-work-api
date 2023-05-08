package handler

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/lonelyRei/course-work/pkg/service"
	"time"
)

// Handler - структура обработчика запросов
// services - указатель на структуру сервиса
type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

// InitRoutes - метод для инициализации обработчиков
// настраиваем конечные точки
func (h *Handler) InitRoutes() *gin.Engine {
	// Инициализация роутера
	router := gin.New()

	// Устанавливаем cors
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "http://localhost:5173/", "*"},
		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET", "DELETE"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type", "Accept-Encoding"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Credentials", "Access-Control-Allow-Headers", "Access-Control-Allow-Methods"},
		AllowCredentials: true,

		AllowOriginFunc: func(origin string) bool {
			fmt.Println(origin == "http://localhost:5173" || origin == "http://localhost:5173/")
			return origin == "http://localhost:5173" || origin == "http://localhost:5173/"
		},
		MaxAge: 12 * time.Hour,
	}))

	// Группируем обработчики по маршрутам
	// Группируем обработчики для авторизации и регистрации
	auth := router.Group("/auth")
	{
		// Обработчик регистрации нового пользователя
		auth.POST("/sign-up", h.signUp)

		// Обработчик авторизации
		auth.POST("sign-in", h.signIn)
	}

	// Группируем обработчики для api
	api := router.Group("/api")
	{
		// Группа обработчиков для работы с пользователями
		user := api.Group("/user")
		{
			// Обработчик получения всех данных о пользователе
			user.GET("", h.userIdentity, h.GetUserData)

			// Обработчик изменения информации о пользователе
			user.PUT("", h.userIdentity, h.SetUserData)

			// Обработчик получения данных о пользователе по идентификатору
			user.GET("/:id", h.GetUserById)
		}
		// Группа обработчиков для работы с записями
		posts := api.Group("/posts")
		{
			// Обработчик создания записи
			posts.POST("", h.userIdentity, h.CreatePost)

			// Обработчик получения записи по идентификатору
			posts.GET("/:id", h.GetPostById)

			// Обработчик получения всех записей
			posts.GET("", h.GetAllPosts)

			// Обработчик изменения записи
			posts.PUT("/:id", h.userIdentity, h.UpdatePost)

			// Обработчик удаления записи
			posts.DELETE("/:id", h.userIdentity, h.DeletePost)

			// Обработчик изменения рейтинга записи
			// параметры: id - идентификатор записи, action: increase, decrease - увеличение / уменьшение рейтинга
			posts.PUT("/rating/:id/:action", h.userIdentity, h.SetRating)

			// Обработчик получения всех постов пользователя по его идентификатору
			posts.GET("/user/:id", h.GetUserPosts)
		}
	}
	// Возвращеаем сконфигурированный роутер
	return router
}
