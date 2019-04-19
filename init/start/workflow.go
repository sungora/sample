package start

import (
	"github.com/sungora/app/workflow"

	"github.com/sungora/sample/internal/users"
)

func workers() {
	workflow.TaskAddCron(&users.One{})
	workflow.TaskAddCron(&users.Two{})
	workflow.TaskAddCron(&users.Four{})
}
