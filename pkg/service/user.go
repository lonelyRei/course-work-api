package service

import (
	"github.com/lonelyRei/course-work/entities"
	"github.com/lonelyRei/course-work/pkg/repository"
)

// UserService - структура сервиса взаимодействия с данными пользователей
type UserService struct {
	repo repository.User
}

// NewUserService - конструктор сервиса авторизации
func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

// GetUserData - получение информации о пользователе
func (s *UserService) GetUserData(id int) (entities.FullUser, error) {
	return s.repo.GetUserData(id)
}

// SetUserData - обновление информации о пользователе
func (s *UserService) SetUserData(id int, user entities.UpdateUserInput) error {
	return s.repo.SetUserData(id, user)
}
