package sample

import (
	"net/http"

	"github.com/go-chi/chi"

	"github.com/sungora/sample/internal/sample/apiv1/groups"
	"github.com/sungora/sample/internal/sample/apiv1/users"
	"github.com/sungora/sample/internal/sample/middleware"
	"github.com/sungora/sample/internal/sample/page"
)

// ApiV1 роутинг api запросов /api/v1/sample
func ApiV1() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.SampleOne)
	r.HandleFunc("/", page.ApiV1)
	// sample routes for "users" resource
	r.Route("/user", func(r chi.Router) {
		r.Use(middleware.SampleTwo)
		r.Get("/", users.Gets)        // array /user
		r.Get("/search", users.Gets)  // array search /user
		r.Post("/", users.Post)       // POST /user
		r.Options("/", users.Options) // OPTIONS /user
		r.Route("/{ID}", func(r chi.Router) {
			r.Get("/", users.Get)       // GET /user/123
			r.Put("/", users.Put)       // PUT /user/123
			r.Delete("/", users.Delete) // DELETE /user/123
		})

	})
	// sample routes for "groups" resource
	r.Route("/group", func(r chi.Router) {
		r.Use(middleware.SampleFour)
		r.Get("/", groups.Gets)        // array /group
		r.Get("/search", groups.Gets)  // array search /group
		r.Post("/", groups.Post)       // POST /group
		r.Options("/", groups.Options) // OPTIONS /group
		r.Route("/{ID}", func(r chi.Router) {
			r.Get("/", groups.Get)       // GET /group/123
			r.Put("/", groups.Put)       // PUT /group/123
			r.Delete("/", groups.Delete) // DELETE /group/123
		})

	})
	return r
}
