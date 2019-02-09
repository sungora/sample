package model

import (
	"sample/controller/api"
	"sample/controller/page"

	"github.com/sungora/app/core"
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

	core.Route.Set("/", page.NewControlSample)
	core.Route.Set("/api/model", api.NewControlModel)

	// sample group route
	core.Route.Path("/api").Path("/v1").
		Set("/page1", api.NewControlModel).
		Set("/page2", api.NewControlModel).
		Set("/page3", api.NewControlModel)
	core.Route.Path("/api").Path("/v2").
		Set("/page1", api.NewControlModel).
		Set("/page2", api.NewControlModel).
		Set("/page3", api.NewControlModel)
	core.Route.Path("/api").Path("/v2").Path("/page").
		Set("/page1", api.NewControlModel).
		Set("/page2", api.NewControlModel).
		Set("/page3", api.NewControlModel)
	core.Route.Path("/api").Path("/v2").Path("/page").Path("/page2").
		Set("/page1", api.NewControlModel).
		Set("/page2", api.NewControlModel).
		Set("/page3", api.NewControlModel)

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
