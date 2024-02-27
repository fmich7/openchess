package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/rekjef/openchess/pkg/utils"
)

type Env struct {
	logger *utils.Logger
}

func NewEnv(logger *utils.Logger) *Env {
	return &Env{logger: logger}
}

func (e *Env) LoadENV(mode string) error {
	switch mode {
	case "dev":
		return e.loadEnvFile("dev.env")
	case "prod":
		return e.loadEnvFile("test.env")
	default:
		return fmt.Errorf("unsupported mode: %s", mode)
	}
}

func (e *Env) loadEnvFile(filename string) error {
	configPath := filepath.Join("internal", "config", filename)
	if err := godotenv.Load(configPath); err != nil {
		return err
	}

	return nil
}

func (e *Env) GetEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		e.logger.Error.Fatalf("ENV KEY=%s is not assigned", key)
	}
	return value
}
