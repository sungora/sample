package core

import (
	"github.com/sungora/app"
	"github.com/sungora/app/connect"
	"github.com/sungora/app/lg"
	"github.com/sungora/app/servhttp"
	"github.com/sungora/app/workflow"

	"github.com/sungora/sample/internal/users"
)

type Config struct {
	App      app.Config      `yaml:"App"`
	Lg       lg.Config       `yaml:"Lg"`
	Workflow workflow.Config `yaml:"Workflow"`
	Http     servhttp.Config `yaml:"Http"`
	Connect  connect.Config  `yaml:"Connect"`
	Users    users.Config    `yaml:"Users"`
}

var Cfg = &Config{}
