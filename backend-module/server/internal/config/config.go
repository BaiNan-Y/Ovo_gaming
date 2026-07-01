package config

import "os"

type Config struct {
	HTTPAddr     string
	DatabaseURL  string
	JWTSecret    string
	AppEnv       string
	AdminWebBase  string
	WeChatAppID   string
	WeChatAppKey  string
}

func Load() Config {
	return Config{
		HTTPAddr:    getEnv("HTTP_ADDR", ":8080"),
		DatabaseURL: getEnv("DATABASE_URL", ""),
		JWTSecret:   getEnv("JWT_SECRET", "change-me"),
		AppEnv:      getEnv("APP_ENV", "dev"),
		AdminWebBase: getEnv("ADMIN_WEB_BASE", "http://localhost:5173"),
		WeChatAppID:  getEnv("WECHAT_APP_ID", ""),
		WeChatAppKey: getEnv("WECHAT_APP_KEY", ""),
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
