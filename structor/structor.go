package structor

import (
	"fmt"
	"os"
	"path/filepath"
)

//TODO: Uses $GOPATH
const ROOT_PATH string = "/home/yukpiz/.go/extend/src/github.com/yukpiz/kipi-structor"

func Execute() {
	var config Config
	fpath := filepath.Join(ROOT_PATH, "kipi.yml")
	if err := LoadConfig(fpath, &config); err != nil {
		panic(err)
	}
	fmt.Println(config)
}

func TextOnMemory(f *os.File) error {
	return nil
}
