package config

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/sungora/app/core"
)

// конфигурация
var config *configTyp

// Init инициализация конфигурации
func Init(suffix string) (code int) {
	sep := string(os.PathSeparator)
	config = new(configTyp)
	var err error

	// читаем конфигурацию
	path := core.Config.DirConfig + sep + core.Config.ServiceName + "_" + suffix + ".toml"
	if _, err = toml.DecodeFile(path, config); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return 1
	}

	// Назанчение конфигураций
	SampleCustomConfig = config.SampleCustomConfig

	return
}

var SampleCustomConfig *sampleCustomConfig
