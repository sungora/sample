package main

import (
	"os"

	"sample/internal/sample"

	"github.com/sungora/app/core"
)

func main() {

	if 0 != core.Init() {
		os.Exit(1)
	}

	sample.Init()

	os.Exit(core.Start())

}
