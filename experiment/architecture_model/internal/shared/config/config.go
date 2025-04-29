package config

import (
	"fmt"
	"os"
)

type Config struct {
	DBUser      string
	DBPass      string
	DBName      string
	DBHost      string // ローカル: "127.0.0.1"; 本番: インスタンス接続名 (PROJECT:REGION:INSTANCE)
	DBPort      string // ローカル環境用
	UseCloudSQL bool
}

func NewConfig() Config {
	useCloudSQL := os.Getenv("USE_CLOUD_SQL") == "true"
	cfg := Config{
		DBUser:      os.Getenv("DB_USER"),
		DBPass:      os.Getenv("DB_PASSWORD"),
		DBName:      os.Getenv("DB_NAME"),
		DBPort:      os.Getenv("DB_PORT"),
		UseCloudSQL: useCloudSQL,
	}
	if useCloudSQL {
		cfg.DBHost = fmt.Sprintf("%s:%s:%s", os.Getenv("GCP_PROJECT_ID"), os.Getenv("DB_REGION"), os.Getenv("DB_INSTANCE"))
	} else {
		cfg.DBHost = os.Getenv("DB_HOST")
	}
	return cfg
}
