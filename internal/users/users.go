package users

import (
	"net/http"

	"github.com/go-chi/chi"

	"github.com/sungora/sample/internal/users/apiv1"
	"github.com/sungora/sample/internal/users/middlep"
)

type Config struct {
	Name     string  `yaml:"Name"`
	IsAccess bool    `yaml:"IsAccess"`
	Balance  float64 `yaml:"Balance"`
}

// RoutesApiV1 роутинг api запросов
// /api/v1/users
func RoutesApiV1() http.Handler {
	r := chi.NewRouter()
	r.Use(middlep.SampleOne)
	// sample routes for "users" resource
	r.Route("/user", func(r chi.Router) {
		r.Use(middlep.SampleTwo)
		r.Get("/", apiv1.UserGets)        // array /user
		r.Post("/", apiv1.UserPost)       // POST /user
		r.Options("/", apiv1.UserOptions) // OPTIONS /user
		r.Route("/{ID}", func(r chi.Router) {
			r.Get("/", apiv1.UserGet)       // GET /user/123
			r.Put("/", apiv1.UserPut)       // PUT /user/123
			r.Delete("/", apiv1.UserDelete) // DELETE /user/123
		})

	})
	// sample routes for "groups" resource
	r.Route("/group", func(r chi.Router) {
		r.Use(middlep.SampleFour)
		r.Get("/", apiv1.GroupGets)        // array /group
		r.Post("/", apiv1.GroupPost)       // POST /group
		r.Options("/", apiv1.GroupOptions) // OPTIONS /group
		r.Route("/{ID}", func(r chi.Router) {
			r.Get("/", apiv1.GroupGet)       // GET /group/123
			r.Put("/", apiv1.GroupPut)       // PUT /group/123
			r.Delete("/", apiv1.GroupDelete) // DELETE /group/123
		})

	})
	return r
}
