package main

import (
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/lonelyRei/course-work"
	"github.com/lonelyRei/course-work/configs"
	"github.com/lonelyRei/course-work/pkg/handler"
	"github.com/lonelyRei/course-work/pkg/repository"
	"github.com/lonelyRei/course-work/pkg/service"
	"github.com/spf13/viper"
	"log"
	"os"
)

func main() {
	// Читаем конфигурационный файл
	if err := configs.InitConfig(); err != nil {
		log.Fatalf("error happened while reading config file: %s", err.Error())
	}

	// Загружаем переменные окружения
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error happened while reading env configuration: %s", err.Error())
	}

	// Инициализация базы данных
	db, err := repository.NewPostgresDB(repository.DBConfig{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.ssl_mode"),
	})

	if err != nil {
		log.Fatalf("error happedned while connecting database: %s", err.Error())
	}

	// Выполняем внедрение зависимостей
	// Создаем экземпляр структуры репозитория
	repos := repository.NewRepository(db)
	// Создаем экземпляр структуры сервиса
	services := service.NewService(repos)
	// Создаем экземпляр структуры маршрутизатора
	handler := handler.NewHandler(services)

	// Инициализация экземпляра сервера
	srv := new(course_work.Server)

	// Запускаем сервер, предварительно читая порт из файлов конфигурации и устанавливая маршрутизатор
	if err := srv.Run(viper.GetString(configs.ConfigPort), handler.InitRoutes()); err != nil {
		log.Fatalf("error happened while running server: %s", err.Error())
	}
}
