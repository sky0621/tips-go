package main

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

const schemaFile = "schema/schema.sql"

type mysqlConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	DryRun   bool
}

func main() {
	cfg, err := loadMySQLConfig()
	if err != nil {
		log.Fatalf("load config: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := applySchema(ctx, cfg, schemaFile); err != nil {
		log.Fatalf("apply schema: %v", err)
	}

	log.Printf("schema %s applied successfully", schemaFile)
}

func loadMySQLConfig() (mysqlConfig, error) {
	cfg := mysqlConfig{
		Host:     getEnv("SQLDEF_MYSQL_HOST", "127.0.0.1"),
		Port:     getEnv("SQLDEF_MYSQL_PORT", "3306"),
		User:     os.Getenv("SQLDEF_MYSQL_USER"),
		Password: os.Getenv("SQLDEF_MYSQL_PASSWORD"),
		Database: os.Getenv("SQLDEF_MYSQL_DATABASE"),
		DryRun:   isTruthy(os.Getenv("SQLDEF_DRY_RUN")),
	}

	if cfg.User == "" {
		return mysqlConfig{}, errors.New("SQLDEF_MYSQL_USER must be set")
	}
	if cfg.Database == "" {
		return mysqlConfig{}, errors.New("SQLDEF_MYSQL_DATABASE must be set")
	}

	return cfg, nil
}

func applySchema(ctx context.Context, cfg mysqlConfig, schemaPath string) error {
	if _, err := exec.LookPath("mysqldef"); err != nil {
		return fmt.Errorf("mysqldef not found in PATH: %w", err)
	}

	absSchemaPath, err := filepath.Abs(schemaPath)
	if err != nil {
		return fmt.Errorf("resolve schema path: %w", err)
	}

	payload, err := os.ReadFile(absSchemaPath)
	if err != nil {
		return fmt.Errorf("read schema: %w", err)
	}

	args := []string{
		"--host", cfg.Host,
		"--port", cfg.Port,
		"--user", cfg.User,
		"--password", cfg.Password,
	}

	if cfg.DryRun {
		args = append(args, "--dry-run")
	}

	args = append(args, cfg.Database)

	cmd := exec.CommandContext(ctx, "mysqldef", args...)
	cmd.Stdin = bytes.NewReader(payload)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func isTruthy(v string) bool {
	switch v {
	case "1", "true", "TRUE", "True", "yes", "YES", "on", "ON":
		return true
	default:
		return false
	}
}
