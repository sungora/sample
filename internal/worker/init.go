package worker

import (
	"github.com/sungora/app/core"
	"github.com/sungora/app/workflow"

	"sample/internal/worker/workfour"
	"sample/internal/worker/workone"
	"sample/internal/worker/worktwo"
)

// init регистрация компонента в приложении
func init() {
	component = new(componentTyp)
	core.ComponentReg(component)
}

// компонент
type componentTyp struct {
}

var (
	component *componentTyp // компонент
)

// Init инициализация компонента в приложении
func (comp *componentTyp) Init(cfg *core.ConfigRoot) (err error) {

	workflow.TaskAddCron("SampleTaskOne", &workone.SampleTaskOne{})
	workflow.TaskAddCron("SampleTaskTwo", &worktwo.SampleTaskTwo{})
	workflow.TaskAddCron("SampleTaskFour", &workfour.SampleTaskFour{})

	return
}

// Start запуск компонента в работу
func (comp *componentTyp) Start() (err error) {
	return
}

// Stop завершение работы компонента
func (comp *componentTyp) Stop() (err error) {
	return
}
