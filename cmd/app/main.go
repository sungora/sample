package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/go-chi/chi"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sungora/app"
	"github.com/sungora/app/connect"
	"github.com/sungora/app/servhttp"
	"github.com/sungora/app/workflow"
	"gopkg.in/webnice/migrate.v1/goose"

	"github.com/sungora/sample/internal"
	"github.com/sungora/sample/internal/config"
)

// @title Sample Service API
// @description Описание сервиса
// @version 1.0
// @contact.name API Support
// @contact.email test@test.ru
// @termsOfService http://swagger.io/terms/
//
// @host {{.Host}}
// @BasePath /api/v1
// @schemes http
//
// @tag.name General
// @tag.description Общие запросы и авторизация
//
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	os.Exit(Start())
}

func Start() (code int) {
	var (
		err       error
		router    *chi.Mux
		component app.Componenter
	)

	// FLAGS & CONFIGURATION
	flagConfigPath := flag.String("c", "config.yaml", "used for set path to config file")
	flagMigrate := flag.Bool("migrate", true, "used for disable migration mode")
	flagLogDB := flag.Bool("log-db", false, "used for disable loading routes on start up")
	flag.Parse()
	// загрузка конфигурации
	Cfg := &config.Config{}
	if err = app.ConfigLoad(*flagConfigPath, Cfg); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		return 1
	}
	app.ConfigSetDefault(&Cfg.App)

	// COMPONENTS
	// connect
	if component, err = connect.Init(&Cfg.Connect, *flagLogDB); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		return 1
	}
	app.ComponentAdd(component)
	// workflow
	if component, err = workflow.Init(&Cfg.Workflow); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		return 1
	}
	app.ComponentAdd(component)
	// servhttp
	if component, router, err = servhttp.Init(&Cfg.Http); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		return 1
	}
	app.ComponentAdd(component)

	// MIGRATION DB
	if *flagMigrate == true {
		if err = goose.Up(connect.GetDB().DB(), Cfg.Connect.Postgresql.Migration); err != nil {
			_, _ = fmt.Fprintln(os.Stderr, err)
			return 1
		}
	}

	// ROUTES
	internal.Routes(router, Cfg)

	// APP MODULES INIT

	// START запуск и остановка приложения
	if err = app.StartLock(); err != nil {
		return 1
	}
	return
}
