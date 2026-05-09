package config

import (
	"encoding/json"
	"os"
)

const (
	configFileName = "/.gatorconfig.json"
)

type Config struct {
	Db_URL       string `json:"db_url"`
	Current_user string `json:"current_user_name"`
}

func getConfigFilePath() (string, error) {
	homePath, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	filepath := homePath + configFileName
	return filepath, nil
}

func Read() (Config, error) {
	var config Config
	filepath, err := getConfigFilePath()
	data, err := os.ReadFile(filepath)
	if err != nil {
		return Config{}, err
	}
	if err := json.Unmarshal(data, &config); err != nil {
		return Config{}, err
	}
	return config, nil
}

func (cfg Config) write() error {
	data, err := json.Marshal(cfg)
	if err != nil {
		return err
	}
	filePath, err := getConfigFilePath()
	if err != nil {
		return err
	}
	if err = os.WriteFile(filePath, data, 0755); err != nil {
		return err
	}
	return nil
}

func (cfg Config) SetUser(name string) error {
	cfg.Current_user = name
	if err := cfg.write(); err != nil {
		return err
	}
	return nil
}
