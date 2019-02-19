package controller

import "github.com/sungora/app/server"

// конфигурация подгружаемая из файла конфигурации
type configFile struct {
	Server server.ConfigTyp
}
