package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Db_url            *string `json:"db_url"`
	Current_user_name *string `json:"current_user_name"`
}

const configFileName = "/.gatorconfig.json"

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return homeDir + configFileName, nil
}

func Read() (*Config, error) {
	var result *Config

	configFilePath, err := getConfigFilePath()
	if err != nil {
		return result, err
	}

	rawFile, err := os.ReadFile(configFilePath)
	if err != nil {
		return result, err
	}

	json.Unmarshal(rawFile, &result)
	return result, nil
}

func (c Config) SetUser(userName string) error {
	c.Current_user_name = &userName
	rawJson, err := json.Marshal(c)
	if err != nil {
		return err
	}

	filePath, err := getConfigFilePath()
	if err != nil {
		return err
	}
	err = os.WriteFile(filePath, rawJson, 0666)
	return err
}
