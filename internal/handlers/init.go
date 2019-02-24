package handlers

import (
	"sample/internal/handlers/apiv1"

	"github.com/BurntSushi/toml"
	"github.com/go-chi/chi"

	"github.com/sungora/app/core"
	"github.com/sungora/app/server"

	"os"
)

// init регистрация компонента в приложении
func init() {
	component = new(componentTyp)
	core.ComponentReg(component)
}

var (
	config    *configTyp    // конфигурация
	component *componentTyp // компонент
)

// компонент
type componentTyp struct {
	serverHTTP *server.HandlerHttp
}

// конфигурация подгружаемая из файла конфигурации
type configTyp struct {
	Server server.ConfigTyp
}

// Init инициализация компонента в приложении
func (comp *componentTyp) Init(cfg *core.ConfigRoot) (err error) {

	sep := string(os.PathSeparator)
	config = new(configTyp)

	// читаем конфигурацию
	path := cfg.DirConfig + sep + cfg.ServiceName + ".toml"
	if _, err = toml.DecodeFile(path, config); err != nil {
		return
	}

	var route = chi.NewRouter();
	route.HandleFunc("/", PageMain)
	route.HandleFunc("/api", PageApi)

	route.Mount("/api/v1", apiv1.Routes())

	comp.serverHTTP = server.NewHandlerHttp(config.Server, route)

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
