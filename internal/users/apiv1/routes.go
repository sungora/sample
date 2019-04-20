package apiv1

import (
	"net/http"

	"github.com/go-chi/chi"

	"github.com/sungora/sample/internal/users/middlep"
)

// Routes роутинг api запросов
// /api/v1/users
func Routes() http.Handler {
	r := chi.NewRouter()
	r.Use(middlep.SampleOne)
	// sample routes for "users" resource
	r.Route("/user", func(r chi.Router) {
		r.Use(middlep.SampleTwo)
		r.Get("/", UserGets)        // array /user
		r.Post("/", UserPost)       // POST /user
		r.Options("/", UserOptions) // OPTIONS /user
		r.Route("/{ID}", func(r chi.Router) {
			r.Get("/", UserGet)       // GET /user/123
			r.Put("/", UserPut)       // PUT /user/123
			r.Delete("/", UserDelete) // DELETE /user/123
		})

	})
	// sample routes for "groups" resource
	r.Route("/group", func(r chi.Router) {
		r.Use(middlep.SampleFour)
		r.Get("/", GroupGets)        // array /group
		r.Post("/", GroupPost)       // POST /group
		r.Options("/", GroupOptions) // OPTIONS /group
		r.Route("/{ID}", func(r chi.Router) {
			r.Get("/", GroupGet)       // GET /group/123
			r.Put("/", GroupPut)       // PUT /group/123
			r.Delete("/", GroupDelete) // DELETE /group/123
		})

	})
	return r
}
