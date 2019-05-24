package internal

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/sungora/app/request"
	"github.com/sungora/app/servhttp/middles"
	"github.com/swaggo/http-swagger"

	"github.com/sungora/sample/docs"
	"github.com/sungora/sample/internal/api"
	"github.com/sungora/sample/internal/config"
	"github.com/sungora/sample/internal/middlep"
)

func Routes(router *chi.Mux, cfg *config.Config) {

	router.Use(middles.TimeoutContext(cfg.Http.WriteTimeout))
	router.Use(middles.Cors(cfg.Http.Cors).Handler)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Logger)

	// general
	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%d", cfg.App.Domain, cfg.Http.Port)
	router.Get("/api/v1/*", httpSwagger.Handler()) // swagger
	router.Get("/api/v1/ping", handlerPing)        // check server work

	// users
	router.Mount("/api/v1/users", RouteUsers())
}

// handlerPing проверка доступности сервера
// @Tags General
// @Router /ping [get]
// @Summary проверка доступности сервера
// @Success 200 {string} string
// @Failure 500 {string} string
func handlerPing(w http.ResponseWriter, r *http.Request) {
	request.NewIn(w, r).Json("pong")
}

// RouteUsers роутинг api запросов
// /api/v1/users
func RouteUsers() http.Handler {
	r := chi.NewRouter()
	r.Use(middlep.SampleOne)
	// sample routes for "users" resource
	r.Route("/user", func(r chi.Router) {
		r.Use(middlep.SampleTwo)
		r.Get("/", api.UserGets)        // array /user
		r.Post("/", api.UserPost)       // POST /user
		r.Options("/", api.UserOptions) // OPTIONS /user
		r.Route("/{ID}", func(r chi.Router) {
			r.Get("/", api.UserGet)       // GET /user/123
			r.Put("/", api.UserPut)       // PUT /user/123
			r.Delete("/", api.UserDelete) // DELETE /user/123
		})

	})
	// sample routes for "groups" resource
	r.Route("/group", func(r chi.Router) {
		r.Use(middlep.SampleFour)
		r.Get("/", api.GroupGets)        // array /group
		r.Post("/", api.GroupPost)       // POST /group
		r.Options("/", api.GroupOptions) // OPTIONS /group
		r.Route("/{ID}", func(r chi.Router) {
			r.Get("/", api.GroupGet)       // GET /group/123
			r.Put("/", api.GroupPut)       // PUT /group/123
			r.Delete("/", api.GroupDelete) // DELETE /group/123
		})

	})
	return r
}
