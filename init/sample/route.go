package sample

import (
	"github.com/go-chi/chi"

	"github.com/sungora/sample/internal/sample"
	"github.com/sungora/sample/internal/sample/middleware"
	"github.com/sungora/sample/internal/sample/page"
)

func routes(r *chi.Mux) {

	// Group 1
	r.Group(func(r chi.Router) {
		r.HandleFunc("/", page.Main)
		r.HandleFunc("/api", page.Api)
		r.HandleFunc("/api/v1", page.ApiV1)
	})

	// Group 2
	r.Group(func(r chi.Router) {
		r.Use(middleware.SamplePing)
		r.Get("/ping", page.Ping)                                          // sample more routes
		r.Get("/version", page.Version)                                    // sample more routes
		r.Get("/test/{testID}/order/{orderID}/page/{pageID}", page.Sample) // sample more routes
	})

	r.Mount("/api/v1/sample", sample.ApiV1())
}
