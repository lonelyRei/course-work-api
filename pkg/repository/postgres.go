package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

// DBConfig - конфигурация для подключения к базе данных
type DBConfig struct {
	// Хост для подключения
	Host string

	// Порт для подключения
	Port string

	// Имя пользователя для подключения
	Username string

	// Пароль для подключения
	Password string

	// Имя базы данных
	DBName string

	// Протокол подключения
	SSLMode string
}

const (
	driverName = "postgres"
)

// NewPostgresDB - синглтон базы данных
func NewPostgresDB(config DBConfig) (*sqlx.DB, error) {
	// Пробуем установить соединение с базой данных и сразу же пингуем ее
	dbInstance, err := sqlx.Connect(driverName, getFormattedConfigData(config))

	if err != nil {
		return nil, err
	}

	return dbInstance, nil
}

// getFormattedConfigData - метод для получения строки для подключения к базе данных postgres
func getFormattedConfigData(config DBConfig) string {
	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		config.Host,
		config.Port,
		config.Username,
		config.DBName,
		config.Password,
		config.SSLMode)
}
