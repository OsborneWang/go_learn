package config

import (
	"fmt"
	"os"
	"strconv"
)

// Config 用于集中管理服务运行所需的关键配置。
type Config struct {
	ServerPort string

	DBHost     string
	DBPort     int
	DBUser     string
	DBPassword string
	DBName     string
	DBCharset  string
	DBLoc      string

	MailHost     string
	MailPort     int
	MailUsername string
	MailPassword string
	MailFrom     string
}

// Load 从环境变量中构建 Config，提供合理默认值，便于本地调试。
func Load() (*Config, error) {
	cfg := &Config{
		ServerPort: getEnv("APP_PORT", "8080"),

		DBHost:     getEnv("DB_HOST", "127.0.0.1"),
		DBPort:     getEnvAsInt("DB_PORT", 3306),
		DBUser:     getEnv("DB_USER", "root"),
		DBPassword: getEnv("DB_PASSWORD", ""),
		DBName:     getEnv("DB_NAME", "go_demo"),
		DBCharset:  getEnv("DB_CHARSET", "utf8mb4"),
		DBLoc:      getEnv("DB_LOC", "Local"),

		MailHost:     getEnv("MAIL_HOST", "smtp.example.com"),
		MailPort:     getEnvAsInt("MAIL_PORT", 587),
		MailUsername: getEnv("MAIL_USERNAME", "demo@example.com"),
		MailPassword: getEnv("MAIL_PASSWORD", "change_me"),
		MailFrom:     getEnv("MAIL_FROM", "Demo Service <demo@example.com>"),
	}

	if cfg.DBName == "" {
		return nil, fmt.Errorf("DB_NAME 不能为空")
	}

	return cfg, nil
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok && value != "" {
		return value
	}
	return fallback
}

func getEnvAsInt(key string, fallback int) int {
	if value, ok := os.LookupEnv(key); ok {
		if parsed, err := strconv.Atoi(value); err == nil {
			return parsed
		}
	}
	return fallback
}
