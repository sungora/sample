package main

import (
	"os"

	"github.com/sungora/app/core"

	"sample/internal/sample"
)

func main() {
	// инициализация компонентов
	if 0 != core.Init() {
		os.Exit(1)
	}

	// инициализация модуля
	sample.Init()

	// запуск приложения
	os.Exit(core.Start())
}
