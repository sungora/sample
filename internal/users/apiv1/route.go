package apiv1

import (
	"net/http"

	"github.com/go-chi/chi"

	users2 "github.com/sungora/sample/internal/users"
)

// Routes роутинг api запросов /api/v1
func Routes() http.Handler {
	r := chi.NewRouter()
	r.Use(users2.SampleOne)
	// sample routes for "users" resource
	r.Route("/user", func(r chi.Router) {
		r.Use(users2.SampleTwo)
		r.Get("/", users2.Gets)        // array /user
		r.Post("/", users2.Post)       // POST /user
		r.Options("/", users2.Options) // OPTIONS /user
		r.Route("/{ID}", func(r chi.Router) {
			r.Get("/", users2.Get)       // GET /user/123
			r.Put("/", users2.Put)       // PUT /user/123
			r.Delete("/", users2.Delete) // DELETE /user/123
		})

	})
	// sample routes for "groups" resource
	r.Route("/group", func(r chi.Router) {
		r.Use(users2.SampleFour)
		r.Get("/", users2.Gets)        // array /group
		r.Post("/", users2.Post)       // POST /group
		r.Options("/", users2.Options) // OPTIONS /group
		r.Route("/{ID}", func(r chi.Router) {
			r.Get("/", users2.Get)       // GET /group/123
			r.Put("/", users2.Put)       // PUT /group/123
			r.Delete("/", users2.Delete) // DELETE /group/123
		})

	})
	return r
}
