package page

import (
	"bytes"
	"net/http"

	"github.com/go-chi/chi"

	"github.com/sungora/app/lg"
	"github.com/sungora/app/request"
	"github.com/sungora/app/servhttp"
	"github.com/sungora/app/tpl"

	"github.com/sungora/sample/internal/core"
	"github.com/sungora/sample/internal/model"
)

// Main главная страница /
func Main(w http.ResponseWriter, r *http.Request) {
	var err error

	var count int = 10
	// работа с моделью
	u := model.NewUser(0)
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
	if buf, err = tpl.ParseFile("assets/controllers/page/sample.html", Functions, Variables); err != nil {
		request.NewIn(w, r).Html("<H1>Internal Server Error</H1>"+err.Error(), 500)
		return
	}
	// шаблон макета
	Variables["Content"] = buf.String()
	if buf, err = tpl.ParseFile("assets/layout/index.html", Functions, Variables); err != nil {
		request.NewIn(w, r).Html("<H1>Internal Server Error</H1>"+err.Error(), 500)
		return
	}
	request.NewIn(w, r).Html(buf.String())
	return
}

// Sample Пример многоуровневого роутинга и GET параметры
func Sample(w http.ResponseWriter, r *http.Request) {
	testID := chi.URLParam(r, "testID")
	orderID := chi.URLParam(r, "orderID")
	pageID := chi.URLParam(r, "pageID")
	response := []string{
		testID,
		orderID,
		pageID,
	}

	ctx := chi.NewRouteContext()
	servhttp.GetRoute().Match(ctx, r.Method, r.URL.Path)
	lg.Dumper(ctx.RoutePattern())

	request.NewIn(w, r).JsonOk(response, 0, "OK")
}

// Ping проверка работы приложения /ping
func Ping(w http.ResponseWriter, r *http.Request) {
	request.NewIn(w, r).Json("pong")
}

// Info информация о приложении
func Info(w http.ResponseWriter, r *http.Request) {
	request.NewIn(w, r).Json(core.Cfg)
}
