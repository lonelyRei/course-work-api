package service

import (
	"github.com/lonelyRei/course-work/entities"
	"github.com/lonelyRei/course-work/pkg/repository"
)

// PostService - структура сервиса взаимодействия с записями
type PostService struct {
	repo repository.Post
}

func NewPostService(repo repository.Post) *PostService {
	return &PostService{repo: repo}
}

// CreatePost - создание записи
func (s *PostService) CreatePost(userid int, post entities.Post) (int, error) {
	return s.repo.CreatePost(userid, post)
}

// GetPostById - получение записи по идентификатору
func (s *PostService) GetPostById(postId int) (entities.Post, error) {
	return s.repo.GetPostById(postId)
}

// GetAllPosts - получение всех записей
func (s *PostService) GetAllPosts() ([]entities.Post, error) {
	return s.repo.GetAllPosts()
}

// DeletePost - удаление записи
func (s *PostService) DeletePost(userId int, postId int) error {
	return s.repo.DeletePost(userId, postId)
}

// UpdatePost - изменение записи
func (s *PostService) UpdatePost(userId, postId int, updatedPost entities.UpdatePostInput) error {
	return s.repo.UpdatePost(userId, postId, updatedPost)
}

// SetRating - метод для изменения рейтинга записи
func (s *PostService) SetRating(postId int, action string) (int, error) {
	return s.repo.SetRating(postId, action)
}

// GetUserPosts - метод получения всех постов конкретного пользователя
func (s *PostService) GetUserPosts(userId int) ([]entities.Post, error) {
	return s.repo.GetUserPosts(userId)
}
