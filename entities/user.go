package entities

// User - тип сущности, описывающей пользователя
type User struct {
	// Идентификатор пользователя
	Id int `json:"-" db:"id"`

	// Имя пользователя
	Name string `json:"name" binding:"required"`

	// Username пользователя
	Username string `json:"username" binding:"required"`

	// Пароль пользователя
	Password string `json:"password" binding:"required"`
}

// FullUser - тип сущности, описывающей полного пользователя
type FullUser struct {
	// Идентификатор пользователя
	Id int `json:"id" db:"id"`

	// Имя пользователя
	Name string `json:"name" binding:"required"`

	// Username пользователя
	Username string `json:"username" binding:"required"`

	// О пользователе
	About string `json:"about"`

	// Изображение пользователя
	Image string `json:"image"`
}

// UpdateUserInput - входные параметры для изменения данных о пользователе
type UpdateUserInput struct {
	// Имя пользователя
	Name *string `json:"name"`

	// О пользователе
	About *string `json:"about"`

	// Изображение пользователя
	Image *string `json:"image"`
}
