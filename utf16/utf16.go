package utf16

import (
	"bytes"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
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

func ReadUTF16File(fpath string) ([]byte, error) {
	raw, err := ioutil.ReadFile(fpath)
	if err != nil {
		return nil, err
	}

	win16be := unicode.UTF16(unicode.BigEndian, unicode.IgnoreBOM)
	utf16bom := unicode.BOMOverride(win16be.NewDecoder())

	unicodeReader := transform.NewReader(bytes.NewReader(raw), utf16bom)
	decoded, err := ioutil.ReadAll(unicodeReader)
	return decoded, err
}
