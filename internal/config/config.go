package config

import (
	"log/slog"
	"os"
	"strconv"
	"strings"

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
	loadEnvFiles()

	cfg := &Config{
		From:     mustEnv("FROM"),
		Password: mustEnv("PASSWORD"),
		Host:     mustEnv("HOST"),
		Port:     envInt("PORT", 465),
	}

	validate(cfg)

	logger.Info("config loaded",
		"host", cfg.Host,
		"port", cfg.Port,
		"from", cfg.From,
	)

	return cfg
}

func loadEnvFiles() {
	candidates := make([]string, 0, 4)

	if custom := strings.TrimSpace(os.Getenv("ENV_FILE")); custom != "" {
		candidates = append(candidates, custom)
	}

	candidates = append(candidates,
		"config.env",
		".env",
	)

	for _, file := range candidates {
		if _, err := os.Stat(file); err == nil {
			if err := godotenv.Load(file); err == nil {
				logger.Info("env file loaded", "file", file)
				return
			}
			logger.Warn("env file exists but failed to load", "file", file, "error", err.Error())
		}
	}

	logger.Info("using process environment only (no env file loaded)")
}

func mustEnv(key string) string {
	v, ok := os.LookupEnv(key)
	if !ok || strings.TrimSpace(v) == "" {
		logger.Error("missing required env", "key", key)
		os.Exit(1)
	}
	return strings.TrimSpace(v)
}

func envInt(key string, fallback int) int {
	v, ok := os.LookupEnv(key)
	if !ok || strings.TrimSpace(v) == "" {
		return fallback
	}

	n, err := strconv.Atoi(strings.TrimSpace(v))
	if err != nil {
		logger.Error("invalid integer env", "key", key, "value", v, "error", err.Error())
		os.Exit(1)
	}
	return n
}

func validate(cfg *Config) {
	if cfg.Port <= 0 {
		logger.Error("invalid PORT", "value", cfg.Port)
		os.Exit(1)
	}
}
