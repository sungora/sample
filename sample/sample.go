package main

import (
	"os"

	_ "sample/config"
	_ "sample/controller"
	_ "sample/model"
	_ "sample/worker"

	"github.com/sungora/app/core"
)

func main() {

	os.Exit(core.Start())

}
