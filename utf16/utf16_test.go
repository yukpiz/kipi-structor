package utf16

import (
	"os"
	"testing"
)

const fpath string = "/home/yukpiz/.go/extend/src/github.com/yukpiz/kipi-structor/data/itemdb.japan.txt"

func TestNewUTF16Scanner(t *testing.T) {
	f, err := os.Open(fpath)
	if err != nil {
		t.Errorf("Can not open text file: %s => %s", fpath, err)
	}
	s := NewUTF16Scanner(f)
	if s == nil {
		t.Errorf("Scanner is nil.")
	}
}

func TestReadUTF16File(t *testing.T) {
	b, err := ReadUTF16File(fpath)
	if err != nil {
		t.Errorf("Failed to read text file.")
	}
	if len(b) == 0 {
		t.Errorf("Content is empty.")
	}
}
