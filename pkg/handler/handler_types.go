package handler

// authUserInput - тип сущности, описывающей тело запроса для авторизации
type authUserInput struct {
	// Username пользователя
	Username string `json:"username" binding:"required"`

	// Пароль пользователя
	Password string `json:"password" binding:"required"`
}

// authUserOutput - тип сущности, описывающей тело ответа на запрос логина
type authUserOutput struct {
	// Идентификатор пользователя
	Id int `json:"id"`

	// Имя пользователя
	Name string `json:"name"`

	// Username пользователя
	Username string `json:"username"`
}
