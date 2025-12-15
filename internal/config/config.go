package config

import (
	"os"
)

type Config struct {
	DBURL string
}

func New() *Config {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		dbURL = "postgres://user:pass@localhost:5432/notes?sslmode=disable"
	}
	return &Config{DBURL: dbURL}
}
