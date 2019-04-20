package start

import (
	"flag"
	"fmt"
	"os"

	"github.com/go-chi/chi"
	"github.com/sungora/app"
	"github.com/sungora/app/connect"
	"github.com/sungora/app/lg"
	"github.com/sungora/app/servhttp"
	"github.com/sungora/app/workflow"

	"github.com/sungora/sample/init/core"
)

func Start() (code int) {
	var (
		err       error
		route     *chi.Mux
		component app.Componenter
	)

	// Флаги
	flagConfigPath := flag.String("c", "config.yaml", "used for set path to config file")
	flag.Parse()

	// загрузка конфигурации
	if err = app.ConfigLoad(*flagConfigPath, core.Cfg); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		return 1
	}
	app.ConfigSetDefault(&core.Cfg.App)

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
	if component, route, err = servhttp.Init(&core.Cfg.Http); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		return 1
	}
	app.ComponentAdd(component)

	// ROUTES
	routes(route)

	// START запуск и остановка приложения
	if err = app.StartLock(); err != nil {
		return 1
	}
	return
}
