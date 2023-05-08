package entities

// Post - тип структуры, описывающей запись, созданную пользователем
type Post struct {
	// Идентификатор записи
	Id int `json:"id" db:"id"`

	// Идентификатор пользователя
	UserId int `json:"user_id" db:"user_id"`

	// Заголовок записи
	Title string `json:"title" binding:"required"`

	// Содержимое записи
	Content string `json:"content" binding:"required"`

	// Изображение в записи
	Image *string `json:"image"`

	// Рейтинг записи
	Rating int `json:"rating"`
}

// UpdatePostInput - входные параметры для изменения данных о записи
type UpdatePostInput struct {
	// Заголовок записи
	Title *string `json:"title"`

	// Содержимое записи
	Content *string `json:"content"`

	// Изображение записи
	Image *string `json:"image"`
}

// Констатнты для взаимодействия с рейтингом записи
const (
	// Увеличение рейтинга записи
	IncreaseRatingAction = "increase"

	// Уменьшение рейтинга записи
	DecreaseRatingAction = "decrease"
)
