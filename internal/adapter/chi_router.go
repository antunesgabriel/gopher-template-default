package adapter

import (
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

type ChiRouter struct {
	mux *chi.Mux
}

func NewChiRouter() *ChiRouter {
	cr := ChiRouter{
		mux: chi.NewRouter(),
	}

	origin := os.Getenv("APP_FRONT_ORIGIN")

	// TODO: adapater this
	cr.mux.Use(middleware.Logger)
	cr.mux.Use(middleware.Heartbeat("/ping"))
	cr.mux.Use(middleware.AllowContentType("application/json","text/xml"))
	cr.mux.Use(
		cors.Handler(cors.Options{
		AllowedOrigins:   []string{origin},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
  	}))

	return &cr
}

func (c *ChiRouter) Use(middleware func(http.Handler) http.Handler) {
	c.mux.Use(middleware)
}

func (c *ChiRouter) Get(path string, handleFn func(http.ResponseWriter, *http.Request)) {
	c.mux.Get(path, handleFn)
}

func (c *ChiRouter) Put(path string, handleFn func(http.ResponseWriter, *http.Request)) {
	c.mux.Put(path, handleFn)
}

func (c *ChiRouter) Post(path string, handleFn func(http.ResponseWriter, *http.Request)) {
	c.mux.Post(path, handleFn)
}

func (c *ChiRouter) Patch(path string, handleFn func(http.ResponseWriter, *http.Request)) {
	c.mux.Patch(path, handleFn)
}

func (c *ChiRouter) Delete(path string, handleFn func(http.ResponseWriter, *http.Request)) {
	c.mux.Delete(path, handleFn)
}

func (c *ChiRouter) Handler() http.Handler {
	return c.mux
}
