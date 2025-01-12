package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path"
)

const configFilename = ".feedlyconfig.json"

type Config struct {
	DBURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() (*Config, error) {
	filePath, err := getConfigFilePath()
	if err != nil {
		return &Config{}, fmt.Errorf("getConfigFilePath: %w", err)
	}

	file, err := os.Open(filePath)
	if err != nil {
		return &Config{}, fmt.Errorf("open: %w", err)
	}
	data, err := io.ReadAll(file)
	if err != nil {
		return &Config{}, fmt.Errorf("readAll: %w", err)
	}

	var cfg Config
	err = json.Unmarshal(data, &cfg)
	if err != nil {
		return &Config{}, fmt.Errorf("unmarshal: %w", err)
	}

	return &cfg, nil
}

func (c *Config) SetUser(userName string) error {
	c.CurrentUserName = userName

	err := write(*c)
	if err != nil {
		return fmt.Errorf("write: %w", err)
	}

	return nil
}

func write(cfg Config) error {
	filePath, err := getConfigFilePath()
	if err != nil {
		return fmt.Errorf("getConfigFilePath: %w", err)
	}

	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("openFile: %w", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(cfg)
	if err != nil {
		return fmt.Errorf("encode: %w", err)
	}

	return nil
}

func getConfigFilePath() (string, error) {
	homePath, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("userHomeDir: %w", err)
	}

	return path.Join(homePath, configFilename), nil
}
