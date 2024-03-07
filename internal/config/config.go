package config

import (
	"github.com/spf13/viper"
	filepath2 "path/filepath"
	"time"
)

// Дефолтные значения параметров конфигурации
const (
	defaultHttpHost      = "localhost"
	defaultHttpPort      = "8080"
	defaultHttpRWTimeout = 10 * time.Second
	//defaultHttpMaxHeaderMegabytes = 1
	defaultLoggerLevel     = 5
	defaultAccessTokenTTL  = 15 * time.Minute
	defaultRefreshTokenTTL = 24 * time.Hour * 30
)

// Config - Сущность для конфигураций
type Config struct {
	Mongo       MongoConfig
	HTTP        HTTPConfig
	Auth        AuthConfig
	Email       EmailConfig
	LoggerLevel int
	CacheTTL    time.Duration
}

// MongoConfig - Конфиг MongoDB
type MongoConfig struct {
	URI          string
	User         string
	Password     string
	DatabaseName string
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
	Host         string
	Port         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type EmailConfig struct {
	ListID       string
	ClientID     string
	ClientSecret string
	Provider     string
	Port         string
	Email        string
	Password     string
}

// Init - функция, создает конфиг из переменных окружения
// На вход подается путь до файлов конфигурации всего приложения и переменных окружения с чувствительной информацией
// Если указаны не все параметры конфигурации, то используются дефолтные константы
func Init(path string, envPath string) (*Config, error) {
	setDefaults()

	var cfg Config

	if err := parseConfigFile(path); err != nil {
		return nil, err
	}

	if err := unmarshal(&cfg); err != nil {
		return nil, err
	}

	if err := parseConfigFile(envPath); err != nil {
		return nil, err
	}

	if err := unmarshalEnv(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

// setDefaults - Устанавливает заданным полям конфигурации дефолтные значения
func setDefaults() {
	viper.SetDefault("http.host", defaultHttpHost)
	viper.SetDefault("http.port", defaultHttpPort)
	viper.SetDefault("http.timeout.read", defaultHttpRWTimeout)
	viper.SetDefault("http.timeout.write", defaultHttpRWTimeout)
	viper.SetDefault("logger.level", defaultLoggerLevel)
	viper.SetDefault("auth.accessTokenTTL", defaultAccessTokenTTL)
	viper.SetDefault("auth.refreshTokenTTL", defaultRefreshTokenTTL)
}

// unmarshal - Парсит из viper объекта необходимые значения конфигурации в заданные структуры, полученные из yml
func unmarshal(cfg *Config) error {
	if err := viper.UnmarshalKey("logger.level", &cfg.LoggerLevel); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("http", &cfg.HTTP); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("auth", &cfg.Auth.JWT); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("cache.ttl", &cfg.CacheTTL); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("email.provider", &cfg.Email.Provider); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("email.port", &cfg.Email.Port); err != nil {
		return err
	}

	return nil
}

// unmarshalEnv - Парсит из viper объекта необходимые значения конфигурации в заданные структуры, полученные из .env
func unmarshalEnv(cfg *Config) error {
	if err := viper.UnmarshalKey("mongo_uri", &cfg.Mongo.URI); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("mongo_user", &cfg.Mongo.User); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("mongo_pass", &cfg.Mongo.Password); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("mongo_databasename", &cfg.Mongo.DatabaseName); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("hash_salt", &cfg.Auth.PasswordSalt); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("secret", &cfg.Email.ClientSecret); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("id", &cfg.Email.ClientID); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("list_id", &cfg.Email.ListID); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("email_sender", &cfg.Email.Email); err != nil {
		return err
	}

	if err := viper.UnmarshalKey("email_password", &cfg.Email.Password); err != nil {
		return err
	}

	return viper.UnmarshalKey("jwt_signing_key", &cfg.Auth.JWT.SigningKey)
}

// parseConfigFile - Считывает конфиг из файла конфигурации
func parseConfigFile(filepath string) error {
	path := filepath2.Dir(filepath)
	name := filepath2.Base(filepath)

	viper.AddConfigPath(path)
	viper.SetConfigName(name)

	return viper.ReadInConfig()
}
