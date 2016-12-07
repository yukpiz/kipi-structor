package utf16

import (
	"os"
	"testing"
)

func TestNewUTF16Scanner(t *testing.T) {
	fpath := "/home/yukpiz/.go/extend/src/github.com/yukpiz/kipi-structor/data/itemdb.japan.txt"
	f, err := os.Open(fpath)
	if err != nil {
		t.Errorf("Can not open text file: %s => %s", fpath, err)
	}
	s := NewUTF16Scanner(f)
	if s == nil {
		t.Errorf("Scanner is nil.")
	}
}
