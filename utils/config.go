package utils

import (
	"fmt"
	"io/ioutil"
	"path"

	"gopkg.in/yaml.v2"

	"github.com/desmos-labs/changeset/types"
)

// GetConfigFilePath returns the path to the configuration file, or an error if something goes wrong
func GetConfigFilePath() (string, error) {
	baseFolderPath, err := getBaseFolderPath()
	if err != nil {
		return "", err
	}

	return path.Join(baseFolderPath, "config.yaml"), nil
}

// WriteConfig writes the given configuration instance inside the default file
func WriteConfig(config *types.Config) error {
	err := config.Validate()
	if err != nil {
		return err
	}

	configFilePath, err := GetConfigFilePath()
	if err != nil {
		return err
	}

	bytes, err := yaml.Marshal(config)
	if err != nil {
		return fmt.Errorf("error while marshalling config: %s", err)
	}

	err = ioutil.WriteFile(configFilePath, bytes, 0666)
	if err != nil {
		return fmt.Errorf("error while writing config: %s", err)
	}

	return nil
}

// ReadConfig reads and parses the configuration file
func ReadConfig() (*types.Config, error) {
	filePath, err := GetConfigFilePath()
	if err != nil {
		return nil, err
	}

	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("error while reading config file: %s", err)
	}

	var config types.Config
	err = yaml.Unmarshal(bytes, &config)
	if err != nil {
		return nil, fmt.Errorf("error while parsing config: %s", err)
	}

	err = config.Validate()
	if err != nil {
		return nil, err
	}

	return &config, nil
}
