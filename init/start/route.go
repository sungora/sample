package start

import (
	"bytes"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/sungora/app/lg"
	"github.com/sungora/app/request"
	"github.com/sungora/app/servhttp"
	"github.com/sungora/app/tpl"

	"github.com/sungora/app/servhttp/middles"

	"github.com/sungora/sample/init/core"
	"github.com/sungora/sample/internal/page"
	"github.com/sungora/sample/internal/users"
)

func routes(route *chi.Mux) {

	route.Use(middles.TimeoutContext(core.Cfg.Http.WriteTimeout))
	route.Use(middleware.Recoverer)
	route.Use(middleware.Logger)
	route.NotFound(middles.NotFound)

	// Group 1
	route.Group(func(r chi.Router) {
		r.HandleFunc("/", page.Main)
		r.HandleFunc("/api", page.Main)
	})

	// Group 2
	route.Group(func(r chi.Router) {
		r.Use(users.SamplePing)
		r.Get("/api/ping", page.Ping)                                      // sample more routes
		r.Get("/api/info", page.Info)                                      // sample more routes
		r.Get("/test/{testID}/order/{orderID}/page/{pageID}", page.Sample) // sample more routes
	})

	route.Mount("/api/v1", users.Routes())
}


// Handler главная страница
func HandlerMain(w http.ResponseWriter, r *http.Request) {
	var err error

	var count int = 10
	// работа с моделью
	u := users.NewUser(0)
	invoiceID := uint64(8697)
	u.InvoiceID = &invoiceID
	name := "Вася пупкин"
	u.Nam = &name
	u.Age = count

	var (
		Variables = make(map[string]interface{})
		Functions = make(map[string]interface{})
	)
	Variables["Header"] = "Head Control"
	Variables["User"] = u

	// шаблон контроллера
	var buf bytes.Buffer
	if buf, err = tpl.ParseFile(core.Cfg.App.DirWork+"/controllers/page/sample.html", Functions, Variables); err != nil {
		request.NewIn(w, r).Html("<H1>Internal Server Error</H1>"+err.Error(), 500)
		return
	}
	// шаблон макета
	Variables["Content"] = buf.String()
	if buf, err = tpl.ParseFile(core.Cfg.App.DirWork+"/layout/index.html", Functions, Variables); err != nil {
		request.NewIn(w, r).Html("<H1>Internal Server Error</H1>"+err.Error(), 500)
		return
	}
	request.NewIn(w, r).Html(buf.String())
	return
}


// Ping проверка работы приложения /ping
func Ping(w http.ResponseWriter, r *http.Request) {
	request.NewIn(w, r).Json("pong")
}

