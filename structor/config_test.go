package structor

import (
	"path/filepath"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	fpath := filepath.Join(ROOT_PATH, "kipi.yml")
	var config Config
	if err := LoadConfig(fpath, &config); err != nil {
		t.Errorf("Failed to read yaml file => %s", err)
	}
	t.Log(config.Data.ItemDbFile)
	t.Log(config.Data.JapanTextFile)
}
