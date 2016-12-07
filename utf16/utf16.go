package utf16

import (
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"os"
)

type UTF16Scanner interface {
	Read(p []byte) (n int, err error)
}

func NewUTF16Scanner(f *os.File) UTF16Scanner {
	win16be := unicode.UTF16(unicode.BigEndian, unicode.IgnoreBOM)
	utf16bom := unicode.BOMOverride(win16be.NewDecoder())
	unicodeReader := transform.NewReader(f, utf16bom)
	return unicodeReader
}
