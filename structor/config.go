package structor

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

const Root string = "/home/yukpiz/.go/extend/src/github.com/yukpiz/kipi-structor"

type Config struct {
	Data DataConfig `yaml:"data"`
}

type DataConfig struct {
	ItemDbFile    string `yaml:"ITEM_DB_FILE"`
	JapanTextFile string `yaml:"JAPAN_TEXT_FILE"`
}

func LoadConfig(path string, config *Config) error {
	bytes := read(path)
	if err := yaml.Unmarshal(bytes, config); err != nil {
		return err
	}
	return nil
}

func read(path string) []byte {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return bytes
}