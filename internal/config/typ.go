package config

import (
	"github.com/sungora/app"
	"github.com/sungora/app/connect"
	"github.com/sungora/app/servhttp"
	"github.com/sungora/app/workflow"
)

type Config struct {
	App      app.Config      `yaml:"App"`
	Workflow workflow.Config `yaml:"Workflow"`
	Http     servhttp.Config `yaml:"Http"`
	Connect  connect.Config  `yaml:"Connect"`
	Users    Users           `yaml:"Users"`
}

type Users struct {
	Name     string  `yaml:"Name"`
	IsAccess bool    `yaml:"IsAccess"`
	Balance  float64 `yaml:"Balance"`
}

var Cfg *Config
