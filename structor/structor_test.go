package structor

import (
	"os"
	"path/filepath"
	"testing"
)

func TestExecute(t *testing.T) {
	if err := Execute(); err != nil {
		t.Errorf("Failed to Execute(): %s", err)
	}
}

func TestTextOnMemory(t *testing.T) {
	fpath := filepath.Join(ROOT_PATH, "data/itemdb.japan.txt")
	f, err := os.Open(fpath)
	if err != nil {
		t.Errorf("Can not open text file: %s => %s", fpath, err)
	}
	err = TextOnMemory(f)
	if err != nil {
		t.Errorf("Failed to TextOnMemory(): %s", err)
	}
}
