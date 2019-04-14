package start

import (
	"github.com/sungora/app/workflow"

	"github.com/sungora/sample/internal/worker"
)

func workers() {
	workflow.TaskAddCron(&worker.One{})
	workflow.TaskAddCron(&worker.Two{})
	workflow.TaskAddCron(&worker.Four{})
}
