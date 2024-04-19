package config

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

type Enviroment = string

const (
	DEV     Enviroment = "dev.env"
	TESTING Enviroment = "testing.env"
)

type Config struct {
	Port     string
	Host     string
	User     string
	Dbname   string
	Password string
	Sslmode  string
}

func LoadConfig(mode Enviroment) (Config, error) {
	configPath := filepath.Join("internal", "config", mode)
	if err := godotenv.Load(configPath); err != nil {
		return Config{}, err
	}
	config := Config{}
	configFields := map[string]*string{
		"PORT":              &config.Port,
		"POSTGRES_HOST":     &config.Host,
		"POSTGRES_USER":     &config.User,
		"POSTGRES_DBNAME":   &config.Dbname,
		"POSTGRES_PASSWORD": &config.Password,
		"POSTGRES_SSLMODE":  &config.Sslmode,
	}

	for key, value := range configFields {
		envVal, exists := os.LookupEnv(key)
		if !exists {
			return Config{}, errors.New("environment variable " + key + " not set")
		}
		*value = envVal
	}

	return config, nil
}
