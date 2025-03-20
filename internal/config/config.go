package config

import (
	"app/internal/utils"
	"fmt"
)

const AppSecret = "APP_SECRET"
const DatabaseDsn = "DATABASE_DSN"

type Loader struct {
}

type Config struct {
	AppSecret string
	Database
}

type Database struct {
	DSN string
}

func (loader *Loader) createConfig() *Config {
	return &Config{
		utils.GetEnvStr(AppSecret, ""),
		Database{
			DSN: utils.GetEnvStr(DatabaseDsn, ""),
		},
	}
}

func (loader *Loader) LoadConfig() (*Config, error) {
	var cfg = loader.createConfig()
	if err := loader.validate(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

func (loader *Loader) validate(cfg *Config) error {
	if cfg.AppSecret == "" {
		return loader.createNotNullEnvError(AppSecret)
	}

	if cfg.Database.DSN == "" {
		return loader.createNotNullEnvError(DatabaseDsn)
	}

	return nil
}

func (loader *Loader) createNotNullEnvError(envName string) error {
	return fmt.Errorf("env variable %v cannot be null", envName)
}
