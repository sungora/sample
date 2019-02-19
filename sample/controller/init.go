package controller

import (
	"github.com/BurntSushi/toml"
	"github.com/sungora/app/server"
	"os"
	"sample/controller/api"
	"sample/controller/page"

	"github.com/sungora/app/core"
)

// init регистрация компонента в приложении
func init() {
	component = new(componentTyp)
	core.ComponentReg(component)
}

var (
	config    *configFile   // конфигурация
	component *componentTyp // компонент
)

// компонент
type componentTyp struct {
	serverHTTP *server.HandlerHttp
}

// Init инициализация компонента в приложении
func (comp *componentTyp) Init(cfg *core.ConfigRoot) (err error) {

	sep := string(os.PathSeparator)
	config = new(configFile)

	// читаем конфигурацию
	path := cfg.DirConfig + sep + cfg.ServiceName + ".toml"
	if _, err = toml.DecodeFile(path, config); err != nil {
		return
	}

	var handler = server.NewHandlerHttp(config.Server)

	handler.Path("/").Use(core.MidlSession).
		GET(page.SampleAll).
		POST(page.SampleAll).
		PUT(page.SampleAll).
		DELETE(page.SampleAll)
	handler.Path("/api/v1/model").
		GET(api.ModelGET).
		POST(api.ModelPOST).
		PUT(api.ModelPUT).
		DELETE(api.ModelDELETE)
	handler.Path("/api/v2/model").
		GET(api.ModelGET).
		POST(api.ModelPOST).
		PUT(api.ModelPUT).
		DELETE(api.ModelDELETE)

	comp.serverHTTP = handler

	// core.Route.Set("/", page.NewControlSample)
	// core.Route.Set("/api/model", api.NewControlModel)
	//
	// // sample group route
	// core.Route.Path("/api").Path("/v1").
	// 	Set("/page1", api.NewControlModel).
	// 	Set("/page2", api.NewControlModel).
	// 	Set("/page3", api.NewControlModel)
	// core.Route.Path("/api").Path("/v2").
	// 	Set("/page1", api.NewControlModel).
	// 	Set("/page2", api.NewControlModel).
	// 	Set("/page3", api.NewControlModel)
	// core.Route.Path("/api").Path("/v2").Path("/page").
	// 	Set("/page1", api.NewControlModel).
	// 	Set("/page2", api.NewControlModel).
	// 	Set("/page3", api.NewControlModel)
	// core.Route.Path("/api").Path("/v2").Path("/page").Path("/page2").
	// 	Set("/page1", api.NewControlModel).
	// 	Set("/page2", api.NewControlModel).
	// 	Set("/page3", api.NewControlModel)

	return
}

// Start запуск компонента в работу
func (comp *componentTyp) Start() (err error) {
	err = comp.serverHTTP.Start()
	return
}

// Stop завершение работы компонента
func (comp *componentTyp) Stop() (err error) {
	err = comp.serverHTTP.Stop()
	return
}
