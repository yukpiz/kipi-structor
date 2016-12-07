package main

import (
	"github.com/yukpiz/kipi-structor/structor"
)

func main() {
	if err := structor.Execute(); err != nil {
		panic(err)
	}
	return
}
