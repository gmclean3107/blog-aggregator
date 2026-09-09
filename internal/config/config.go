package config

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	Db_url            string `json:"db_url"`
	Current_user_name string `json:"current_user_name"`
}

func Read() (Config, error) {
	path := getConfigFilePath()

	if path == "" {
		return Config{}, errors.New("error getting filepath")
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return Config{}, err
	}

	cfg := Config{}

	err = json.Unmarshal(data, &cfg)
	if err != nil {
		return Config{}, err
	}

	return cfg, nil
}

func (cfg *Config) SetUser(username string) error {
	cfg.Current_user_name = username

	path := getConfigFilePath()

	if path == "" {
		return errors.New("error getting filepath")
	}

	config, err := json.Marshal(cfg)
	if err != nil {
		return err
	}

	err = os.WriteFile(path, config, 0644)
	if err != nil {
		return err
	}

	return nil
}

func getConfigFilePath() string {
	homeDir, err := os.UserHomeDir()

	if err != nil {
		return ""
	}

	return filepath.Join(homeDir, configFileName)
}
