package adapter

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type ChiRouter struct {
	mux *chi.Mux
}

func NewChiRouter() *ChiRouter {
	cr := ChiRouter{
		mux: chi.NewRouter(),
	}

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
