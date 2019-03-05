package page

import (
	"bytes"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/sungora/app/core"
	"github.com/sungora/app/servhttp/middleware"

	"sample/pkg/sample/model"
)

// Main главная страница
func Main(w http.ResponseWriter, r *http.Request) {
	var rw = r.Context().Value(middleware.KeyRW).(*core.RW)
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
		Variables     = make(map[string]interface{})
		Functions     = make(map[string]interface{})
		TplController = core.Config.DirWww + "/controllers/page/sample.html"
		TplLayout     = core.Config.DirWww + "/layout/index.html"
	)
	Variables["Header"] = "Head Control"
	Variables["User"] = u

	// шаблон контроллера
	var buf bytes.Buffer
	if buf, err = core.TplCompilation(TplController, Functions, Variables); err != nil {
		rw.ResponseHtml("<H1>Internal Server Error</H1>", 500)
		return
	}
	if TplLayout == "" {
		rw.ResponseHtml(buf.String(), 200)
		return
	}
	// шаблон макета
	Variables["Content"] = buf.String()
	if buf, err = core.TplCompilation(TplLayout, Functions, Variables); err != nil {
		rw.ResponseHtml("<H1>Internal Server Error</H1>", 500)
		return
	}
	rw.ResponseHtml(buf.String(), 200)
	return

}

// Api страница api
func Api(w http.ResponseWriter, r *http.Request) {
	var rw = r.Context().Value(middleware.KeyRW).(*core.RW)
	rw.ResponseHtml("IndexApi", 200)
}

// Sample Пример многоуровневого роутинга и GET параметры
func Sample(w http.ResponseWriter, r *http.Request) {
	var rw = r.Context().Value(middleware.KeyRW).(*core.RW)
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
