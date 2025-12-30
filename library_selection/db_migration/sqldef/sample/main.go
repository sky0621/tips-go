package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/sqldef/sqldef"
	"github.com/sqldef/sqldef/database"
	"github.com/sqldef/sqldef/database/mysql"
	"github.com/sqldef/sqldef/parser"
	"github.com/sqldef/sqldef/schema"
)

const schemaFile = "schema/schema.sql"

type mysqlConfig struct {
	Host     string
	Port     int
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
	port, err := parsePort(getEnv("SQLDEF_MYSQL_PORT", "6306"))
	if err != nil {
		return mysqlConfig{}, err
	}

	cfg := mysqlConfig{
		Host:     getEnv("SQLDEF_MYSQL_HOST", "127.0.0.1"),
		Port:     port,
		User:     getEnv("SQLDEF_MYSQL_USER", "app"),
		Password: getEnv("SQLDEF_MYSQL_PASSWORD", "app"),
		Database: getEnv("SQLDEF_MYSQL_DATABASE", "sample"),
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
	absSchemaPath, err := filepath.Abs(schemaPath)
	if err != nil {
		return fmt.Errorf("resolve schema path: %w", err)
	}

	desiredDDLs, err := sqldef.ReadFiles([]string{absSchemaPath})
	if err != nil {
		return fmt.Errorf("read schema: %w", err)
	}

	db, err := mysql.NewDatabase(database.Config{
		DbName:   cfg.Database,
		User:     cfg.User,
		Password: cfg.Password,
		Host:     cfg.Host,
		Port:     cfg.Port,
		SslMode:  "preferred",
	})
	if err != nil {
		return fmt.Errorf("open database: %w", err)
	}
	defer db.Close()

	if err := db.DB().PingContext(ctx); err != nil {
		return fmt.Errorf("ping database: %w", err)
	}

	options := &sqldef.Options{
		DesiredDDLs: desiredDDLs,
		DryRun:      cfg.DryRun,
	}

	sqlParser := database.NewParser(parser.ParserModeMysql)
	sqldef.Run(schema.GeneratorModeMysql, db, sqlParser, options)

	return nil
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

func parsePort(raw string) (int, error) {
	port, err := strconv.Atoi(raw)
	if err != nil || port <= 0 {
		return 0, fmt.Errorf("invalid SQLDEF_MYSQL_PORT: %q", raw)
	}
	return port, nil
}
