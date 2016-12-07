package structor

import (
	"bufio"
	"fmt"
	"github.com/yukpiz/kipi-structor/model"
	"github.com/yukpiz/kipi-structor/utf16"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

//TODO: Uses $GOPATH
const ROOT_PATH string = "/home/yukpiz/.go/extend/src/github.com/yukpiz/kipi-structor"

var MemoryData model.Data

func Execute() error {
	fmt.Println("Execute kipi-structor ===> (•ө•)♡")
	MemoryData = model.Data{}
	//Load yaml configuration.
	var config Config
	fpath := filepath.Join(ROOT_PATH, "kipi.yml")
	if err := LoadConfig(fpath, &config); err != nil {
		return err
	}

	//Read japanese text file.
	fpath = filepath.Join(ROOT_PATH, config.Data.JapanTextFile)
	f, err := os.Open(fpath)
	if err != nil {
		return err
	}

	if err := TextOnMemory(f); err != nil {
		return err
	}
	return nil
}

func TextOnMemory(f *os.File) error {
	s := utf16.NewUTF16Scanner(f)
	scanner := bufio.NewScanner(s)
	line := 1
	for scanner.Scan() {
		seps := strings.SplitN(scanner.Text(), "\t", 2)
		if len(seps) != 2 {
			fmt.Sprintf("LINE: %d => ERROR: Split line strings. %v", line, seps)
		}

		id, err := strconv.Atoi(seps[0])
		if err != nil {
			return err
		}
		name := string(seps[1])
		MemoryData.Items = append(MemoryData.Items, model.Item{Id: id, JpName: name})
		line += 1
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
