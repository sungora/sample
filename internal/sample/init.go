package sample

import (
	"github.com/sungora/app/lg"
	"github.com/sungora/app/servhttp"
	"github.com/sungora/app/workflow"

	"sample/internal/sample/config"
	"sample/internal/sample/worker/workfour"
	"sample/internal/sample/worker/workone"
	"sample/internal/sample/worker/worktwo"
)

// Init инициализация модуля
func Init(cfg *config.Config) (err error) {

	// config
	config.Cfg = cfg

	// router
	servhttp.MountRoutes("/", RoutesPage)
	servhttp.MountRoutes("/api/v1", RoutesApiV1)

	// workers
	workflow.TaskAddCron(&workone.SampleTaskOne{})
	workflow.TaskAddCron(&worktwo.SampleTaskTwo{})
	workflow.TaskAddCron(&workfour.SampleTaskFour{})

	// log
	lg.SetMessages(map[int]string{
		1000: "Message format Fmt from 1000",
		1001: "Message format Fmt from 1001",
		1002: "Message format Fmt from 1002",
		1003: "Message format Fmt from 1003",
		1004: "Message format Fmt from 1004",
		1005: "Message format Fmt from 1005",
	})

	return
}
