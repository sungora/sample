package main

import (
	"os"

	"github.com/sungora/app/connect"
	"github.com/sungora/app/core"
	"github.com/sungora/app/lg"
	"github.com/sungora/app/servhttp"
	"github.com/sungora/app/workflow"

	"sample/internal"
	"sample/pkg/sample"
)

func main() {

	path := "sample.yaml"
	cfg := new(Config)

	err := core.LoadConfigYaml(path, cfg)


	core.LoadConfigYaml()

	// инициализация компонентов
	core.Init()

	if 1 == core.Init() {
		os.Exit(1)
	}

	// инициализация
	if 1 == internal.Init() {
		os.Exit(1)
	}

	// инициализация модуля
	if 1 == sample.Init() {
		os.Exit(1)
	}

	// запуск приложения
	os.Exit(core.Start())
}

type Config struct {
	Core     core.Config     `yaml:"Core"`
	Lg       lg.Config       `yaml:"Lg"`
	Workflow workflow.Config `yaml:"Workflow"`
	Http     servhttp.Config `yaml:"Http"`
	Connect  connect.Config  `yaml:"Connect"`
}
