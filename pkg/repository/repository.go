package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/lonelyRei/course-work/entities"
)

// Authorization - интерфейс авторизации
type Authorization interface {
	// CreateUser - метод создания нового пользователя
	CreateUser(user entities.User) (int, error)

	// GetUser - метод получения пользователя из базы данных
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

// Repository - структура, имплементирующая все интерфейсы пакета
type Repository struct {
	Authorization
	User
	Post
}

// NewRepository - конструктор репозитория
func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		User:          NewUserPostgres(db),
		Post:          NewPostPostgres(db),
	}
}
