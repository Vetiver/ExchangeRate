package config

import (
	"os"
)

type Config struct {
	ProfilerAddr string
	HTTPAddr     string
	BaseUrlCBR   string
}

func NewConfig() *Config {
	return &Config{
		ProfilerAddr: fromEnv("PROFILER_ADDR", "0.0.0.0:6060"),
		HTTPAddr:     fromEnv("HTTP_ADDR", "0.0.0.0:8888"),
		BaseUrlCBR:    fromEnv("BaseUrlCBR", "http://www.cbr.ru"),
	}
}

func fromEnv(key, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}
