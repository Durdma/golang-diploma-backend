package config

import (
	"fmt"
	"github.com/spf13/viper"
	"strings"
	"time"
)

// Дефолтные значения параметров конфигурации
const (
	defaultHttpPort               = "8080"
	defaultHttpRWTimeout          = 10 * time.Second
	defaultHttpMaxHeaderMegabytes = 1
	defaultLoggerLevel            = 5
	defaultAccessTokenTTL         = 15 * time.Minute
	defaultRefreshTokenTTL        = 24 * time.Hour * 30
)

// Config - Сущность для конфигураций
type Config struct {
	Mongo       MongoConfig
	HTTP        HTTPConfig
	Auth        AuthConfig
	LoggerLevel int
}

// MongoConfig - Конфиг MongoDB
type MongoConfig struct {
	URI      string
	User     string
	Password string
	Name     string
}

// AuthConfig - Конфиг аутентификации
type AuthConfig struct {
	JWT          JWTConfig
	PasswordSalt string
}

// JWTConfig - Конфиг системы токенов
type JWTConfig struct {
	AccessTokenTTL  time.Duration
	RefreshTokenTTL time.Duration
	SigningKey      string
}

// HTTPConfig - Конфиг http подключения
type HTTPConfig struct {
	Port         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

// TODO Переписать функции, где извлечение идет из ОС. Добавить эти переменные либо в .env, либо в main.yml
// TODO Посмотреть подключение к монге, можно ли в локальной бд использовать для входа логин+пароль
// Init - функция, создает конфиг из переменных окружения
// На вход подается файл .env в котором описываются значения конфигурации
// Если указаны не все параметры конфигурации, то используются дефолтные константы
func Init(path string) (*Config, error) {
	setDefaults()

	if err := parseConfigFile(path); err != nil {
		return nil, err
	}

	if err := parseEnv(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := unmarshal(&cfg); err != nil {
		return nil, err
	}

	setFromEnv(&cfg)

	return &cfg, nil
}

// setDefaults - Устанавливает заданным полям конфигурации дефолтные значения
func setDefaults() {
	viper.SetDefault("http.port", defaultHttpPort)
	viper.SetDefault("http.timeout.read", defaultHttpRWTimeout)
	viper.SetDefault("http.timeout.write", defaultHttpRWTimeout)
	viper.SetDefault("logger.level", defaultLoggerLevel)
	viper.SetDefault("auth.accessTokenTTL", defaultAccessTokenTTL)
	viper.SetDefault("auth.refreshTokenTTL", defaultRefreshTokenTTL)
}

// unmarshal - Парсит из viper объекта необходимые значения конфигурации в заданные структуры
func unmarshal(cfg *Config) error {
	if err := viper.UnmarshalKey("logger.level", &cfg.LoggerLevel); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("mongo", &cfg.Mongo); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("http", &cfg.HTTP); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("auth", &cfg.Auth); err != nil {
		return err
	}

	return nil
}

// setFromEnv - Устанавливает в конфиг переменные из ОС.
func setFromEnv(cfg *Config) {
	cfg.Mongo.URI = viper.GetString("uri")
	cfg.Mongo.User = viper.GetString("user")
	cfg.Mongo.Password = viper.GetString("password")
	cfg.Auth.PasswordSalt = viper.GetString("salt")
	cfg.Auth.JWT.SigningKey = viper.GetString("signing_key")
}

// TODO Посмотреть подробнее, почему файл конфига находится только по абсолютному пути.
// TODO сделать подключение по относительному
// parseConfigFile - Считывает конфиг из файла конфигурации
func parseConfigFile(filepath string) error {
	path := strings.Split(filepath, "\\")
	fmt.Println(path[:len(path)-1])
	fmt.Println(path[len(path)-1])

	ap := ""

	for i := 0; i < len(path)-1; i++ {
		ap += path[i] + "\\"
	}

	fmt.Println(ap)

	viper.AddConfigPath(ap)
	viper.SetConfigName(strings.Split(path[len(path)-1], ".")[0])

	return viper.ReadInConfig()
}

// parseEnv - Устанавливает связь между ключами viper и именами переменных в ОС
func parseEnv() error {
	if err := parseMongoEnvVariables(); err != nil {
		return err
	}

	if err := parsePasswordEnvVariables(); err != nil {
		return err
	}

	return parseJWTEnvVariables()
}

// parseMongoEnvVariables - Устанавливает в viper соответствия ключам и переменным ОС
func parseMongoEnvVariables() error {
	// Устанавливает префикс для поиска переменных в ОС (в данном случае будут искаться переменные: MONGO_VARIABLE,
	// MONGO_URI и т.д.)
	viper.SetEnvPrefix("mongo")

	// Устанавливает связь между переменной из конфига с переменными viper.
	// Если передается 1 аргумент, то viper будет искать переменную PREFIX_VARIABLE в данном случае MONGO_URI
	// Если бы было несколько аргументов в функцию, то по факту для uri были псевдонимы
	if err := viper.BindEnv("uri"); err != nil {
		return err
	}

	if err := viper.BindEnv("user"); err != nil {
		return err
	}

	return viper.BindEnv("password")
}

// parsePasswordEnvVariables - Аналогично функции parseMongoEnvVariables
func parsePasswordEnvVariables() error {
	viper.SetEnvPrefix("password")
	return viper.BindEnv("salt")
}

// parseJWTEnvVariables - Аналогично функции parseMongoEnvVariables
func parseJWTEnvVariables() error {
	viper.SetEnvPrefix("jwt")
	return viper.BindEnv("signing_key")
}
