package config

import (
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
)

type Config struct {
	Env         string `yaml:"env" env:"ENV" env-required:"true"`
	StoragePath string `yaml:"storage_path" env-requiered:"true"`
	HTTPServer  `yaml:"http_server"`
	Dsn         `yaml:"dsn"`
	// Путь, куда будут загружаться аватарки
	FileUploadPath string `yaml:"FILE_UPLOAD_PATH" env-default:"../../../frontend/public/avatars/"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env-default:"localhost:8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
	User        string        `yaml:"user" env-requiered:"true"`
	Password    string        `yaml:"password" env-requiered:"true" env:"HTTP_SERVER_PASSWORD"`
}

type Dsn struct {
	Host     string `yaml:"DB_HOST"`
	Port     int    `yaml:"DB_PORT"`
	User     string `yaml:"DB_USER"`
	Password string `yaml:"DB_PASSWORD"`
	DBname   string `yaml:"DB_NAME"`
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, errors.Wrap(err, "error loading .env file")
	}

	configPath := os.Getenv("CONFIG_PATH")

	if configPath == "" {
		return nil, errors.New("CONFIG_PATH is not set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return nil, errors.Errorf("config file %s is not exist", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		return nil, errors.Wrap(err, "cannot read config")
	}

	return &cfg, nil
}
