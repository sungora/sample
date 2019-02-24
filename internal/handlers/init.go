package handlers

import (
	"os"
	"time"

	"sample/internal/handlers/page"

	"sample/internal/middle"

	"sample/internal/handlers/apiv1"

	"github.com/BurntSushi/toml"
	"github.com/go-chi/chi"

	"github.com/sungora/app/core"
	"github.com/sungora/app/server"
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

	var r = chi.NewRouter();
	r.NotFound(middle.NotFound)
	r.Use(middle.Main(time.Second*time.Duration(config.Server.WriteTimeout) - time.Millisecond))

	r.Mount("/", page.Routes())
	r.Mount("/api/v1", apiv1.Routes())

	comp.serverHTTP = server.NewHandlerHttp(config.Server, r)

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
