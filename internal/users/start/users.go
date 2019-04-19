package start

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/sungora/app/lg"
	"github.com/sungora/app/servhttp"
	"github.com/sungora/app/workflow"

	"github.com/sungora/sample/internal/users"
	"github.com/sungora/sample/internal/users/apiv1"
	"github.com/sungora/sample/internal/users/middlep"
	"github.com/sungora/sample/internal/users/worker"
)

func Init(cfg *users.Config) {

	// config
	users.Cfg = cfg

	lg.Dumper(users.Cfg)

	// workers
	workflow.TaskAddCron(&worker.One{})
	workflow.TaskAddCron(&worker.Two{})
	workflow.TaskAddCron(&worker.Four{})

	// роутинг
	servhttp.GetRoute().Mount("/api/v1/user", Routes())
}

// Routes роутинг api запросов /api/v1
func Routes() http.Handler {
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
