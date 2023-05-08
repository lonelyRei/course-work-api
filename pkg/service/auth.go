package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/lonelyRei/course-work/entities"
	"github.com/lonelyRei/course-work/pkg/repository"
	"time"
)

const (
	// Соль для хэширования пароля
	salt = "cskdcmksckldsmkcd3991&&&!/"

	// Время жизни токена
	tokenTTL = 12 * time.Hour

	// Ключ сигнатуры
	signedKey = "cmsdmcscsd"
)

// tokenClaims - claims из токена
type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

// AuthService - структура сервиса авторизации
type AuthService struct {
	repo repository.Authorization
}

// NewAuthService - конструктор сервиса авторизации
func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

// CreateUser - метод создания нового пользователя
func (s *AuthService) CreateUser(user entities.User) (int, error) {
	// Хэшируем пароль пользователя
	user.Password = s.generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

// GetUser - метод получения пользователя
func (s *AuthService) GetUser(username, password string) (entities.User, error) {
	// Получаем пользователя из базы
	user, err := s.repo.GetUser(username, s.generatePasswordHash(password))

	if err != nil {
		return entities.User{}, err
	}

	return user, nil
}

// GenerateToken - метод для генерации нового токена
func (s *AuthService) GenerateToken(username, password string) (string, error) {
	// Получаем пользователя из базы
	user, err := s.repo.GetUser(username, s.generatePasswordHash(password))

	if err != nil {
		return "", err
	}

	// генерируем токен
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			// Токен будет жить 12 часов
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})

	return token.SignedString([]byte(signedKey))
}

// ParseToken - парсит токен и возвращает идентификатор пользователя
func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Проверяем подпись токена
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signedKey), nil
	})

	if err != nil {
		return 0, err
	}

	// Приводим к кастомным claims
	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("unable to get correct token claims type")
	}

	return claims.UserId, nil

}

// generatePasswordHash - метод для генерации хэшированного пароля
func (s *AuthService) generatePasswordHash(password string) string {
	// генерируем новый пароль с помощью sha1
	hash := sha1.New()
	hash.Write([]byte(password))

	// Возвращаем хэш версию пароля с добавлением случайных символов
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
