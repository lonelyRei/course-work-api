package configs

import "github.com/spf13/viper"

// Константы для инициализации конфигурационного файла
const (
	configPath     = "configs"
	configFilename = "config"
)

const (
	ConfigPort = "port"
)

// InitConfig - метод инициализации конфига
func InitConfig() error {
	// Добавляем путь к папке конфигурации
	viper.AddConfigPath(configPath)

	// Добавляем путь к конфигурационному файлу
	viper.SetConfigName(configFilename)

	// Возвращаем метод для чтения из конфига
	return viper.ReadInConfig()
}
