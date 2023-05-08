package service

import (
	"github.com/lonelyRei/course-work/entities"
	"github.com/lonelyRei/course-work/pkg/repository"
)

// Authorization - интерфейс авторизации
type Authorization interface {
	// CreateUser - Метод создания нового пользователя
	CreateUser(user entities.User) (int, error)

	// GenerateToken - метод для генерации токена
	GenerateToken(username, password string) (string, error)

	// ParseToken - парсит токен и возвращает id пользователя
	ParseToken(token string) (int, error)

	// GetUser - возвращает всю информацию о пользователе, необходимую после прохождения авторизации
	GetUser(username, password string) (entities.User, error)
}

// User - интерфейс взаимодействия с пользователем
type User interface {
	// GetUserData - получение информации о пользователе
	GetUserData(id int) (entities.FullUser, error)

	// SetUserData - обновление информации о пользователе
	SetUserData(id int, user entities.UpdateUserInput) error
}

// Post - интерфейс взаимодействия с записями
type Post interface {
	// CreatePost - создание записи
	CreatePost(userid int, post entities.Post) (int, error)

	// GetPostById - получение записи по идентификатору
	GetPostById(postId int) (entities.Post, error)

	// GetAllPosts - получение всех записей
	GetAllPosts() ([]entities.Post, error)

	// DeletePost - удаление записи
	DeletePost(userId int, postId int) error

	// UpdatePost - изменение записи
	UpdatePost(userId, postId int, updatedPost entities.UpdatePostInput) error

	// SetRating - метод для изменения рейтинга записи
	SetRating(postId int, action string) (int, error)

	// GetUserPosts - метод получения всех постов конкретного пользователя
	GetUserPosts(userId int) ([]entities.Post, error)
}

// Service - структура, имплементирующая все интерфейсы пакета
type Service struct {
	Authorization
	User
	Post
}

// NewService - конструктор сервиса.
// Параметром принимает указатель на репозиторий
func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		User:          NewUserService(repos.User),
		Post:          NewPostService(repos.Post),
	}
}
