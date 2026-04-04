package config

import (
	"log/slog"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	From     string
	Password string
	Host     string
	Port     int
}

var logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))

func Load() *Config {
	if err := godotenv.Load("config.env"); err != nil {
		logger.Error("failed to load config.env",
			"file", "config.env",
			"error", err.Error(),
		)
		os.Exit(1)
	}

	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		logger.Error("invalid PORT in env",
			"value", os.Getenv("PORT"),
			"error", err.Error(),
		)
		os.Exit(1)
	}

	config := Config{
		Host:     os.Getenv("HOST"),
		Password: os.Getenv("PASSWORD"),
		From:     os.Getenv("FROM"),
		Port:     port,
	}

	return &config
}
