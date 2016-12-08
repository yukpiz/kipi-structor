package structor

import (
	"github.com/yukpiz/kipi-structor/utf16"
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

func TestXmlOnMemory(t *testing.T) {
	fpath := filepath.Join(ROOT_PATH, "data/itemdb.xml")
	buf, err := utf16.ReadUTF16File(fpath)
	if err != nil {
		t.Errorf("Can not open xml file: %s => %s", fpath, err)
	}
	if err := XmlOnMemory(buf); err != nil {
		t.Errorf("Failed to XmlOnMemory(): %s", err)
	}
}
