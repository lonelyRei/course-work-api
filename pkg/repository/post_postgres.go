package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/lonelyRei/course-work/entities"
	"strings"
)

// PostPostgres - структура репозитория для взаимодействия с записями
type PostPostgres struct {
	db *sqlx.DB
}

// NewPostPostgres - конструктор репозитория для взаимодействия с записями
func NewPostPostgres(db *sqlx.DB) *PostPostgres {
	return &PostPostgres{db: db}
}

// CreatePost - создание записи
func (r *PostPostgres) CreatePost(userid int, post entities.Post) (int, error) {
	// Создаем переменную, в которую запишем результат запроса
	var id int

	var query string

	var row *sql.Row

	// Добавляем идентификатор пользователя
	post.UserId = userid

	if post.Image != nil {
		query = fmt.Sprintf("INSERT INTO %s (user_id, title, content, image) values ($1, $2, $3, $4) RETURNING id",
			postsTableName)

		// Выполяем запрос
		row = r.db.QueryRow(query, post.UserId, post.Title, post.Content, post.Image)
	} else {
		query = fmt.Sprintf("INSERT INTO %s (user_id, title, content) values ($1, $2, $3) RETURNING id",
			postsTableName)

		// Выполяем запрос
		row = r.db.QueryRow(query, post.UserId, post.Title, post.Content)
	}

	// Записываем результат запроса в переменную id
	if err := row.Scan(&id); err != nil {
		// Если при выполнении запроса возникла ошибка, то возвращаем ее на слой выше
		return 0, err
	}

	return id, nil
}

// GetPostById - получение записи по идентификатору
func (r *PostPostgres) GetPostById(postId int) (entities.Post, error) {
	// Создаем переменную для результата запроса
	var post entities.Post

	// Создаем строку запроса
	query := fmt.Sprintf("SELECT id, user_id, title, content, image, rating FROM %s where id=$1", postsTableName)

	// Выполняем запрос
	err := r.db.Get(&post, query, postId)

	// Возвращаем пользователя и ошибку
	return post, err
}

// GetAllPosts - получение всех записей
func (r *PostPostgres) GetAllPosts() ([]entities.Post, error) {
	// Создаем переменную, куда сохраним результат запроса
	var posts []entities.Post

	// Создаем строку запроса
	query := fmt.Sprintf("SELECT * FROM %s", postsTableName)

	// Выполняем запрос
	err := r.db.Select(&posts, query)
	return posts, err
}

// DeletePost - удаление записи
func (r *PostPostgres) DeletePost(userId int, postId int) error {
	// Создаем переменную для сохранения результата первого запроса
	var dbUserId int

	// Создаем запрос для получения идентификатора пользователя
	userIdQuery := fmt.Sprintf("SELECT user_id FROM %s WHERE id=$1", postsTableName)

	// Выполняем запрос
	err := r.db.Get(&dbUserId, userIdQuery, postId)
	if err != nil {
		return err
	}

	// Если идентификаторы пользователей (создавшего запись и удаляющего) не совпадают,
	// то создаем ошибку
	if dbUserId != userId {
		return errors.New("have no access to do this")
	}

	// Создаем строку запроса
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", postsTableName)

	// Выполняем запрос
	_, err = r.db.Exec(query, postId)

	return err
}

// UpdatePost - изменение записи
func (r *PostPostgres) UpdatePost(userId, postId int, updatedPost entities.UpdatePostInput) error {
	// Создаем переменную для сохранения результата первого запроса
	var dbUserId int

	// Создаем запрос для получения идентификатора пользователя
	userIdQuery := fmt.Sprintf("SELECT user_id FROM %s WHERE id=$1", postsTableName)

	// Выполняем запрос
	err := r.db.Get(&dbUserId, userIdQuery, postId)
	if err != nil {
		return err
	}

	// Если идентификаторы пользователей (создавшего запись и удаляющего) не совпадают,
	// то создаем ошибку
	if dbUserId != userId {
		return errors.New("have no access to do this")
	}

	// Переменная для сохранения строки запроса к бд
	setValues := make([]string, 0)

	// аргументы запроса
	args := make([]interface{}, 0)

	// идентификатор текущего аргумента
	argId := 1

	// Если заголовк не nil, то добавляем в запрос на изменение
	if updatedPost.Title != nil {
		setValues = append(setValues, fmt.Sprintf(`title=$%d`, argId))
		args = append(args, *updatedPost.Title)
		argId++
	}

	// Если содержание не nil, то добавляем в запрос на изменение
	if updatedPost.Content != nil {
		setValues = append(setValues, fmt.Sprintf(`content=$%d`, argId))
		args = append(args, *updatedPost.Content)
		argId++
	}

	// Если изобржаение не nil, то добавляем в запрос на изменение
	if updatedPost.Image != nil {
		setValues = append(setValues, fmt.Sprintf(`image=$%d`, argId))
		args = append(args, *updatedPost.Image)
		argId++
	}

	// Формируем строку запроса (поля для изменения)
	setQuery := strings.Join(setValues, ", ")

	// Формируем строку запроса (создаем сам запрос)
	query := fmt.Sprintf("UPDATE %s SET %s WHERE id=%d", postsTableName, setQuery, postId)

	// Выполняем запрос
	_, err = r.db.Exec(query, args...)

	// Возвращаем ошибку
	return err
}

// SetRating - метод для изменения рейтинга записи
func (r *PostPostgres) SetRating(postId int, action string) (int, error) {
	// Получаем рейтинг
	var currentRating int

	// Создаем строку запроса для получения текущего рейтинга
	query := fmt.Sprintf("SELECT rating from %s where id=$1", postsTableName)

	err := r.db.Get(&currentRating, query, postId)

	if err != nil {
		return 0, err
	}

	// Формируем строку запроса на изменение рейтинга
	setQuery := fmt.Sprintf("UPDATE %s set rating=$1 where id=%d", postsTableName, postId)

	var ratingAfterChange int

	// Если необходимо увеличить рейтинг записи
	if action == entities.IncreaseRatingAction {
		ratingAfterChange = currentRating + 1
		_, err = r.db.Exec(setQuery, ratingAfterChange)
	} else {
		// Если необходимо уменьшить рейтинг записи
		ratingAfterChange = currentRating - 1
		_, err = r.db.Exec(setQuery, ratingAfterChange)
	}

	if err != nil {
		return 0, err
	}

	return ratingAfterChange, nil
}

// GetUserPosts - метод получения всех постов конкретного пользователя
func (r *PostPostgres) GetUserPosts(userId int) ([]entities.Post, error) {
	// Создаем переменную, куда сохраним результат запроса
	var posts []entities.Post

	// Создаем строку запроса
	query := fmt.Sprintf("SELECT * FROM %s where user_id=$1", postsTableName)

	// Выполняем запрос
	err := r.db.Select(&posts, query, userId)
	return posts, err
}
