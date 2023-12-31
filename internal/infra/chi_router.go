package infra

import (
	"github.com/antunesgabriel/gopher-template-default/internal/config"
	"github.com/antunesgabriel/gopher-template-default/internal/delivery/api"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/jwtauth/v5"
	"net/http"
)

type ChiRouter struct {
	mux       *chi.Mux
	tokenAuth *jwtauth.JWTAuth
	env       *config.Env
}

func NewChiRouter(env *config.Env) *ChiRouter {
	tokenAuth := jwtauth.New("HS256", []byte(env.JWTSignKey), nil)

	cr := ChiRouter{
		mux:       chi.NewRouter(),
		tokenAuth: tokenAuth,
		env:       env,
	}

	origin := env.ClientURL

	cr.mux.Use(middleware.Logger)
	cr.mux.Use(middleware.AllowContentType("application/json", "multipart/form-data"))
	cr.mux.Use(
		cors.Handler(cors.Options{
			AllowedOrigins:   []string{origin},
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
			AllowedHeaders:   []string{"User-Agent", "Content-Type", "Accept", "Accept-Encoding", "Accept-Language", "Cache-Control", "Connection", "DNT", "Host", "Origin", "Pragma", "Referer"},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: true,
			MaxAge:           300,
		}))

	return &cr
}

func (it *ChiRouter) Use(middleware func(http.Handler) http.Handler) {
	it.mux.Use(middleware)
}

func (it *ChiRouter) Get(path string, handleFn http.HandlerFunc) {
	it.mux.Get(path, handleFn)
}

func (it *ChiRouter) Put(path string, handleFn http.HandlerFunc) {
	it.mux.Put(path, handleFn)
}

func (it *ChiRouter) Post(path string, handleFn http.HandlerFunc) {
	it.mux.Post(path, handleFn)
}

func (it *ChiRouter) Patch(path string, handleFn http.HandlerFunc) {
	it.mux.Patch(path, handleFn)
}

func (it *ChiRouter) Delete(path string, handleFn http.HandlerFunc) {
	it.mux.Delete(path, handleFn)
}

func (it *ChiRouter) Handler() http.Handler {
	return it.mux
}

func (it *ChiRouter) ProtectedGroup(path string, handleFn func(r api.RouteGroup)) {
	it.mux.Route(path, func(r chi.Router) {
		r.Use(jwtauth.Verifier(it.tokenAuth))

		r.Use(jwtauth.Authenticator)

		handleFn(r)
	})
}
