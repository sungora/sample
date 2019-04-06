package sample

import (
	"flag"
	"fmt"
	"os"
	"time"

	middlewareChi "github.com/go-chi/chi/middleware"
	"github.com/sungora/app"
	"github.com/sungora/app/connect"
	"github.com/sungora/app/lg"
	"github.com/sungora/app/servhttp"
	"github.com/sungora/app/servhttp/middlew"
	"github.com/sungora/app/workflow"

	"github.com/sungora/sample/internal/core"
	"github.com/sungora/sample/internal/sample/middleware"
)

func Start() (code int) {
	var (
		err             error
		component       app.Componenter
		componentServer *servhttp.Component
	)

	// Флаги
	flagConfigPath := flag.String("c", "config/sample.yaml", "used for set path to config file")
	flag.Parse()

	// загрузка конфигурации
	if err = app.LoadConfig(*flagConfigPath, core.Cfg); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		return 1
	}

	// COMPONENTS
	// logs
	if component, err = lg.Init(&core.Cfg.Lg); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		return 1
	}
	app.ComponentAdd(component)
	// connect
	if component, err = connect.Init(&core.Cfg.Connect); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		return 1
	}
	app.ComponentAdd(component)
	// workflow
	if component, err = workflow.Init(&core.Cfg.Workflow); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		return 1
	}
	app.ComponentAdd(component)
	// servhttp
	if componentServer, err = servhttp.Init(&core.Cfg.Http); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		return 1
	}
	app.ComponentAdd(componentServer)

	// APPLICATION
	// routes
	r := componentServer.GetRoute()
	r.NotFound(middlew.NotFound)
	r.Use(middlew.TimeoutContext(time.Second * time.Duration(core.Cfg.Http.WriteTimeout-1)))
	r.Use(middlewareChi.Recoverer)
	r.Use(middlewareChi.Logger)
	r.Use(middleware.SampleRoot)
	routes(r)
	// workers
	workers()
	// logs
	logs()

	// START запуск и остановка приложения
	if err = app.StartLock(&core.Cfg.App); err != nil {
		return 1
	}
	return
}
