package app

import "net/http"

type Router interface {
	Use(middleware func(http.Handler) http.Handler)
	Get(path string, handleFn func(http.ResponseWriter, *http.Request))
	Put(path string, handleFn func(http.ResponseWriter, *http.Request))
	Post(path string, handleFn func(http.ResponseWriter, *http.Request))
	Patch(path string, handleFn func(http.ResponseWriter, *http.Request))
	Delete(path string, handleFn func(http.ResponseWriter, *http.Request))
	Handler() http.Handler
}
