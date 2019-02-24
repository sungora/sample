package page

import (
	"bytes"
	"net/http"

	"github.com/go-chi/chi"

	"sample/internal/model/users"

	"github.com/sungora/app/core"

	"sample/internal/middle"
)

// MainPage главная страница
func Index(w http.ResponseWriter, r *http.Request) {
	var rw = r.Context().Value(middle.KEY_RW).(*core.RW)
	var err error

	var count int = 10
	// работа с моделью
	u := users.New(0)
	invoiceID := uint64(8697)
	u.InvoiceID = &invoiceID
	name := "Вася пупкин"
	u.Nam = &name
	u.Age = count

	rw.Variables["Header"] = "Head Control"
	rw.Variables["User"] = u
	rw.TplController += "/page/sample.html"
	rw.TplLayout += "/index.html"

	// шаблон контроллера
	var buf bytes.Buffer
	if buf, err = core.TplCompilation(rw.TplController, rw.Functions, rw.Variables); err != nil {
		rw.ResponseHtml("<H1>Internal Server Error</H1>", 500)
		return
	}
	if rw.TplLayout == "" {
		rw.ResponseHtml(buf.String(), 200)
		return
	}
	// шаблон макета
	rw.Variables["Content"] = buf.String()
	if buf, err = core.TplCompilation(rw.TplLayout, rw.Functions, rw.Variables); err != nil {
		rw.ResponseHtml("<H1>Internal Server Error</H1>", 500)
		return
	}
	rw.ResponseHtml(buf.String(), 200)
	return

}

// IndexApi страница api
func IndexApi(w http.ResponseWriter, r *http.Request) {
	var rw = r.Context().Value(middle.KEY_RW).(*core.RW)
	rw.ResponseHtml("IndexApi", 200)
}

// Sample Пример многоуровневого роутинга и GET параметры
func Sample(w http.ResponseWriter, r *http.Request) {
	var rw = r.Context().Value(middle.KEY_RW).(*core.RW)
	testID := chi.URLParam(r, "testID")
	orderID := chi.URLParam(r, "orderID")
	pageID := chi.URLParam(r, "pageID")
	response := []string{
		testID,
		orderID,
		pageID,
	}
	rw.ResponseJsonApi200(response, 0, "OK")
}
