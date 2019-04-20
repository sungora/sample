package start

import (
	"bytes"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/sungora/app/request"
	"github.com/sungora/app/servhttp/middles"
	"github.com/sungora/app/tpl"

	"github.com/sungora/sample/init/core"
	usersApiv1 "github.com/sungora/sample/internal/users/apiv1"
)

func routes(route *chi.Mux) {

	route.Use(middles.TimeoutContext(core.Cfg.Http.WriteTimeout))
	route.Use(middleware.Recoverer)
	route.Use(middleware.Logger)
	route.NotFound(middles.NotFound)

	// Group 1
	route.Group(func(r chi.Router) {
		r.HandleFunc("/", handlerMain)
		r.HandleFunc("/api", handlerPing)
	})

	route.Mount("/api/v1/users", usersApiv1.Routes())
}

// handlerMain главная страница
func handlerMain(w http.ResponseWriter, r *http.Request) {
	var err error

	var count int = 10
	// работа с моделью
	u := struct {
		ID         uint64     ``
		InvoiceID  *uint64    ``
		Nam        *string    ``
		Age        int        ``
		Credit     float64    ``
		IsOnline   bool       `gorm:"not null;default:1;"`
		Status     string     `gorm:"type:enum('Актив','Пассив','Универсал');not null;default:'Пассив';"`
		Hobby      *string    `gorm:"type:set('music','sport','reading', 'stamps', 'travel');"`
		SampleJson *string    `gorm:"type:json;"`
		Address    *string    `gorm:"type:text;"`
		CreatedAt  time.Time  ``
		UpdatedAt  time.Time  ``
		DeletedAt  *time.Time ``
	}{
		ID: 34,
	}

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

// handlerPing проверка работы приложения /ping
func handlerPing(w http.ResponseWriter, r *http.Request) {
	request.NewIn(w, r).Json("pong")
}
