package structor

import (
	"bufio"
	"encoding/xml"
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

var MemoryData *model.Data

func Execute() error {
	fmt.Println("Execute kipi-structor ===> (•ө•)♡")
	MemoryData = new(model.Data)
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

	//Read item xml file.
	fpath = filepath.Join(ROOT_PATH, config.Data.ItemDbFile)
	buf, err := utf16.ReadUTF16File(fpath)
	if err != nil {
		return err
	}

	if err := XmlOnMemory(buf); err != nil {
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
		text := string(seps[1])
		MemoryData.Texts = append(MemoryData.Texts, model.Text{TextId: id, JpText: text})
		line += 1
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func XmlOnMemory(buf []byte) error {
	seps := strings.Split(string(buf), "\n")
	for _, line := range seps {
		if strings.Index(line, "Mabi_Item") < 0 {
			continue
		}

		item := new(model.Item)
		if err := xml.Unmarshal([]byte(line), item); err != nil {
			return err
		}
		item.Convert()

		text0 := MemoryData.FindFromId(item.JpNameId)
		item.JpName = text0.JpText
		text1 := MemoryData.FindFromId(item.JpDescId)
		item.JpDesc = text1.JpText

		MemoryData.Items = append(MemoryData.Items, *item)
	}

	return nil
}
