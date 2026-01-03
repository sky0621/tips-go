package rdb

import (
	"fmt"
	"os"
	"strconv"
)

const (
	defaultHost     = "127.0.0.1"
	defaultPort     = 34507
	defaultUser     = "app"
	defaultPassword = "app"
	defaultDatabase = "ddd_aggregate"
)

type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

func LoadConfig() (Config, error) {
	port, err := parsePort(getEnv("DDD_AGG_DB_PORT", strconv.Itoa(defaultPort)))
	if err != nil {
		return Config{}, err
	}

	cfg := Config{
		Host:     getEnv("DDD_AGG_DB_HOST", defaultHost),
		Port:     port,
		User:     getEnv("DDD_AGG_DB_USER", defaultUser),
		Password: getEnv("DDD_AGG_DB_PASSWORD", defaultPassword),
		Database: getEnv("DDD_AGG_DB_NAME", defaultDatabase),
	}

	if cfg.User == "" {
		return Config{}, fmt.Errorf("DDD_AGG_DB_USER must be set")
	}
	if cfg.Database == "" {
		return Config{}, fmt.Errorf("DDD_AGG_DB_NAME must be set")
	}

	return cfg, nil
}

func (c Config) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&charset=utf8mb4&loc=Local", c.User, c.Password, c.Host, c.Port, c.Database)
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func parsePort(raw string) (int, error) {
	port, err := strconv.Atoi(raw)
	if err != nil || port <= 0 {
		return 0, fmt.Errorf("invalid DDD_AGG_DB_PORT: %q", raw)
	}
	return port, nil
}
