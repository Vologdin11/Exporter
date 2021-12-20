package config

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Tables []Table `yaml:"tables"`
}

type Table struct {
	Name          string `yaml:"name"`
	Value_index   int    `yaml:"value_index"`
	Label_indexes []int  `yaml:"label_indexes"`
}

func GetConfig() (Config, error) {
	yamlConfig, err := os.Open("config.yml")
	if err != nil {
		return Config{}, err
	}
	defer yamlConfig.Close()
	bytes, err := ioutil.ReadAll(yamlConfig)
	if err != nil {
		return Config{}, err
	}
	config := Config{}
	err = yaml.Unmarshal(bytes, &config)
	if err != nil {
		return Config{}, err
	}
	return config, nil
}
