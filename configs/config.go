package configs

import (
	"github.com/kelseyhightower/envconfig"
	"os"
)

type Config struct {
	AppPort   string `default:"8080"`
	DBHost    string `default:"0.0.0.0"`
	DBPort    string `default:"5432"`
	DBUser    string `default:"root"`
	DBPass    string `default:"root"`
	DBName    string `default:"film_library_db"`
	DBSSLMode string `default:"disable"`
}

func NewConfig() Config {
	var s Config
	err := envconfig.Process("", &s)
	if err != nil {
		return Config{}
	}

	if os.Getenv("APP_PORT") != "" {
		s.AppPort = os.Getenv("APP_PORT")
	}
	if os.Getenv("DB_HOST") != "" {
		s.DBHost = os.Getenv("DB_HOST")
	}
	if os.Getenv("DB_PORT") != "" {
		s.DBPort = os.Getenv("DB_PORT")
	}
	if os.Getenv("DB_USER") != "" {
		s.DBUser = os.Getenv("DB_USER")
	}
	if os.Getenv("DB_PASS") != "" {
		s.DBPass = os.Getenv("DB_PASS")
	}
	if os.Getenv("DB_NAME") != "" {
		s.DBName = os.Getenv("DB_NAME")
	}
	if os.Getenv("DB_SSL_MODE") != "" {
		s.DBSSLMode = os.Getenv("DB_SSL_MODE")
	}

	return s
}
