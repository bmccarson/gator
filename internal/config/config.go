package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Db_url            string
	Current_user_name string
}

const (
	configFile = "/.gatorconfig.json"
)

// read the config json file in the home dir and return a config struct
func Read() (Config, error) {
	c := Config{}

	filePath, err := getConfigPath()
	if err != nil {
		return c, err
	}

	configJSON, err := os.ReadFile(filePath)
	if err != nil {
		return c, fmt.Errorf("Couldnt read the file: %s", err)
	}

	err = json.Unmarshal(configJSON, &c)

	return c, nil
}

// set the user field of the config struct
func (c *Config) SetUser(user string) error {
	c.Current_user_name = user

	err := write(*c)
	if err != nil {
		return err
	}

	return nil
}

func getConfigPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("Couldnt get home dir: %s", err)
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
		return fmt.Errorf("Couldnt open the file: %s", err)
	}
	defer file.Close()

	json, err := json.Marshal(cfg)

	_, err = file.Write(json)
	if err != nil {
		return fmt.Errorf("Couldnt write to the file: %s", err)
	}

	return nil

}
