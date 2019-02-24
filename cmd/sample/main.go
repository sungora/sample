package main

import (
	"os"

	"github.com/sungora/app/core"

	_ "sample/internal/config"
	_ "sample/internal/handlers"
	_ "sample/internal/model"
	_ "sample/internal/worker"
)

func main() {

	os.Exit(core.Start())

}
