package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/lonelyRei/course-work/entities"
)

// AuthPostgres - структура репозитория для авторизации
type AuthPostgres struct {
	db *sqlx.DB
}

// NewAuthPostgres - конструктор репозитория авторизации
func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

// CreateUser - метод создания нового пользователя
func (r *AuthPostgres) CreateUser(user entities.User) (int, error) {
	// Создаем переменную, в которую запишем результат запроса
	var id int

	// Создаем запрос к базе данных
	query := fmt.Sprintf(
		"INSERT INTO %s (name, username, password_hash) values ($1, $2, $3) RETURNING id",
		usersTableName)

	// Выполяем запрос
	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)

	// Записываем результат запроса в переменную id
	if err := row.Scan(&id); err != nil {
		// Если при выполнении запроса возникла ошибка, то возвращаем ее на слой выше
		return 0, err
	}

	// Возвращаем id нового пользователя
	return id, nil
}

// GetUser - метод для получения пользователя из базы данных
func (r *AuthPostgres) GetUser(username, password string) (entities.User, error) {
	// Создаем переменную, в которую запишем результат запроса
	var user entities.User

	// Создаем запрос к базе данных
	query := fmt.Sprintf("SELECT id, name, username FROM %s WHERE username=$1 AND password_hash=$2",
		usersTableName)

	// Выполняем запрос
	err := r.db.Get(&user, query, username, password)

	// Возвращем пользователя и ошибку
	return user, err
}
