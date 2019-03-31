package sample

import (
	"github.com/sungora/app/workflow"

	"github.com/sungora/sample/internal/sample/worker/workfour"
	"github.com/sungora/sample/internal/sample/worker/workone"
	"github.com/sungora/sample/internal/sample/worker/worktwo"
)

func workers() {
	workflow.TaskAddCron(&workone.SampleTaskOne{})
	workflow.TaskAddCron(&worktwo.SampleTaskTwo{})
	workflow.TaskAddCron(&workfour.SampleTaskFour{})
}
