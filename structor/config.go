package structor

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	Data DataConfig `yaml:"data"`
}

type DataConfig struct {
	ItemDbFile    string `yaml:"ITEM_DB_FILE"`
	JapanTextFile string `yaml:"JAPAN_TEXT_FILE"`
}

func LoadConfig(path string, config *Config) error {
	bytes, err := read(path)
	if err != nil {
		return err
	}
	if err := yaml.Unmarshal(bytes, config); err != nil {
		return err
	}
	return nil
}

func read(path string) ([]byte, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}
