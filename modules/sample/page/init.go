package page

import (
	"net/http"

	"github.com/go-chi/chi"
)

func Routes() http.Handler {
	r := chi.NewRouter()

	r.HandleFunc("/", Main)
	r.HandleFunc("/api", Api)
	r.Get("/test/{testID}/order/{orderID}/page/{pageID}", Sample) // sample more routes

	return r
}
