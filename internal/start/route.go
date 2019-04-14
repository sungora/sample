package start

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"github.com/sungora/app/servhttp/middles"
	"github.com/sungora/sample/internal/apiv1"
	"github.com/sungora/sample/internal/core"
	"github.com/sungora/sample/internal/middlep"
	"github.com/sungora/sample/internal/page"
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
		r.Use(middlep.SamplePing)
		r.Get("/api/ping", page.Ping)                                      // sample more routes
		r.Get("/api/info", page.Info)                                      // sample more routes
		r.Get("/test/{testID}/order/{orderID}/page/{pageID}", page.Sample) // sample more routes
	})

	route.Mount("/api/v1", apiv1.Routes())
}
