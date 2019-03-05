package main

import (
	"os"

	"sample/internal"

	"sample/pkg/sample"

	"github.com/sungora/app/core"
)

func main() {
	// инициализация компонентов
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
