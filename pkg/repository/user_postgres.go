package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/lonelyRei/course-work/entities"
	"strings"
)

// UserPostgres - структура репозитория для взаимодействия с данными пользователя
type UserPostgres struct {
	db *sqlx.DB
}

// NewUserPostgres - конструктор репозитория для взаимодействия с данными пользователя
func NewUserPostgres(db *sqlx.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

// GetUserData - получение информации о пользователе
func (r *UserPostgres) GetUserData(id int) (entities.FullUser, error) {
	// Создаем переменную, куда сохраним результат запроса
	var user entities.FullUser

	// Создаем запрос к базе данных
	query := fmt.Sprintf("SELECT id, name, username, about, image FROM %s WHERE id=$1",
		usersTableName)

	// Выполняем запрос
	err := r.db.Get(&user, query, id)

	// Возвращаем пользователя и ошибку
	return user, err
}

// SetUserData - обновление информации о пользователе
func (r *UserPostgres) SetUserData(id int, user entities.UpdateUserInput) error {
	// Переменная для сохранения строки запроса к бд
	setValues := make([]string, 0)

	// аргументы запроса
	args := make([]interface{}, 0)

	// идентификатор текущего аргумента
	argId := 1

	// Если имя не nil, то добавляем в запрос на изменение информации о пользователе к бд
	if user.Name != nil {
		setValues = append(setValues, fmt.Sprintf(`name=$%d`, argId))
		args = append(args, *user.Name)
		argId++
	}

	// Если изображение не nil, то добавляем в запрос на изменение информации о пользователе к бд
	if user.Image != nil {
		setValues = append(setValues, fmt.Sprintf(`image=$%d`, argId))
		args = append(args, *user.Image)
		argId++
	}

	// Если информация не nil, то добавляем в запрос на изменение информации о пользователе к бд
	if user.About != nil {
		setValues = append(setValues, fmt.Sprintf(`about=$%d`, argId))
		args = append(args, *user.About)
		argId++
	}

	// Формируем строку запроса (поля для изменения)
	setQuery := strings.Join(setValues, ", ")

	// Формируем строку запроса (создаем сам запрос)
	query := fmt.Sprintf("UPDATE %s SET %s WHERE id=%d", usersTableName, setQuery, id)

	// Выполняем запрос
	_, err := r.db.Exec(query, args...)

	// Возвращаем ошибку
	return err
}
