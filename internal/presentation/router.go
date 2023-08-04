package presentation

import "net/http"

type Router interface {
	Use(middleware func(http.Handler) http.Handler)
	Get(path string, handleFn func(w http.ResponseWriter, r *http.Request))
	Put(path string, handleFn func(w http.ResponseWriter, r *http.Request))
	Post(path string, handleFn func(w http.ResponseWriter, r *http.Request))
	Patch(path string, handleFn func(w http.ResponseWriter, r *http.Request))
	Delete(path string, handleFn func(w http.ResponseWriter, r *http.Request))
	Handler() http.Handler
}
