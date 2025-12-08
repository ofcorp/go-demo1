package config

import "os"

type Config struct {
	Key	 string
}

func NewConfig() *Config {
	key := os.Getenv("KEY")
	if key == "" {
		panic("не передан параметр KEY в переменные окружения")
	}
	return &Config{
		Key: key,
	}
}