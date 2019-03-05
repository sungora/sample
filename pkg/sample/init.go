package sample

import (
	"github.com/sungora/app/lg"
	"github.com/sungora/app/servhttp"
	"github.com/sungora/app/workflow"

	"sample/pkg/sample/apiv1"
	"sample/pkg/sample/config"
	"sample/pkg/sample/page"
	"sample/pkg/sample/worker/workfour"
	"sample/pkg/sample/worker/workone"
	"sample/pkg/sample/worker/worktwo"
)

const ModuleName string = "sample"

// Init инициализация модуля
func Init() (code int) {

	// config
	if 0 < config.Init(ModuleName) {
		return 1
	}

	// router
	servhttp.MountRoutes("/", page.Routes)
	servhttp.MountRoutes("/api/v1", apiv1.Routes)

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
