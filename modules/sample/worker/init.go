package worker

import (
	"github.com/sungora/app/core"
	"github.com/sungora/app/workflow"

	"sample/internal/sample/worker/workfour"
	"sample/internal/sample/worker/workone"
	"sample/internal/sample/worker/worktwo"
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
