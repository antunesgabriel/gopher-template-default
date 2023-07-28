package adapter

import "github.com/go-chi/chi/v5"

type chiRouter struct {
	mux *chi.Mux
}

func NewChiRouter() *chiRouter {
	cr := chiRouter{
		mux: chi.NewRouter(),
	}

	return &cr
}
