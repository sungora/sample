package sample

import (
	"flag"
	"fmt"
	"os"
	"time"

	middlewareChi "github.com/go-chi/chi/middleware"
	"github.com/sungora/app"
	"github.com/sungora/app/connect"
	"github.com/sungora/app/core"
	"github.com/sungora/app/lg"
	"github.com/sungora/app/servhttp"
	"github.com/sungora/app/servhttp/middleware"
	"github.com/sungora/app/workflow"

	"sample/internal/sample"
	"sample/internal/sample/config"
)

type Config struct {
	Core     core.Config     `yaml:"Core"`
	Lg       lg.Config       `yaml:"Lg"`
	Workflow workflow.Config `yaml:"Workflow"`
	Http     servhttp.Config `yaml:"Http"`
	Connect  connect.Config  `yaml:"Connect"`
	Sample   config.Config   `yaml:"Calculator"`
}

const (
	version = "1.0.0"
)

func Init() (code int) {
	var (
		err       error
		component app.Componenter
	)

	configPath := flag.String("c", "config/sample.yaml", "used for set path to config file")
	flag.Parse()

	// загрузка конфигурации
	cfg := &Config{}
	if err = core.LoadConfig(*configPath, cfg); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	// COMPONENTS
	// core
	if component, err = core.Init(&cfg.Core, version); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	app.ComponentAdd(component)
	// logs
	if component, err = lg.Init(&cfg.Lg, cfg.Core.ServiceName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	app.ComponentAdd(component)
	// connect
	if component, err = connect.Init(&cfg.Connect); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	app.ComponentAdd(component)
	// workflow
	if component, err = workflow.Init(&cfg.Workflow); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	app.ComponentAdd(component)
	// servhttp
	if component, err = servhttp.Init(&cfg.Http); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	app.ComponentAdd(component)

	// APPLICATION
	servhttp.MiddlewareRoot(middleware.TimeoutContext(time.Second * time.Duration(cfg.Http.WriteTimeout-1)))
	servhttp.MiddlewareRoot(middlewareChi.Recoverer)
	servhttp.MiddlewareRoot(middlewareChi.Logger)
	servhttp.NotFound(middleware.NotFound)

	// MODULES
	if err = sample.Init(&cfg.Sample); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	// запуск и остановка приложения
	var isStart = int8(2)
	return app.Start(&isStart)
}
