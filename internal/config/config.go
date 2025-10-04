// Package config holds all configuration data
package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	DBURL           string
	CurrentUserName string
}

const (
	configFile = "/.gatorconfig.json"
)

// Read the config json file in the home dir and return a config struct
func Read() (Config, error) {
	c := Config{}

	filePath, err := getConfigPath()
	if err != nil {
		return c, err
	}

	configJSON, err := os.ReadFile(filePath)
	if err != nil {
		return c, fmt.Errorf("couldnt read the file: %s", err)
	}

	err = json.Unmarshal(configJSON, &c)
	if err != nil {
		return c, err
	}

	return c, nil
}

// SetUser sets the CurrentUserName in the Config
func (c *Config) SetUser(user string) error {
	c.CurrentUserName = user

	err := write(*c)
	if err != nil {
		return err
	}

	return nil
}

func getConfigPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("couldnt get home dir: %s", err)
	}

	filepath := homeDir + configFile

	return filepath, nil
}

func write(cfg Config) error {
	filePath, err := getConfigPath()
	if err != nil {
		return err
	}

	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("couldnt open the file: %s", err)
	}
	defer file.Close()

	json, err := json.Marshal(cfg)
	if err != nil {
		return err
	}

	_, err = file.Write(json)
	if err != nil {
		return fmt.Errorf("couldnt write to the file: %s", err)
	}

	return nil
}
