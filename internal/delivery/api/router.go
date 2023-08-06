package api

import "net/http"

type RouteGroup interface {
	Get(path string, h http.HandlerFunc)
	Put(path string, h http.HandlerFunc)
	Post(path string, h http.HandlerFunc)
	Patch(path string, h http.HandlerFunc)
	Delete(path string, h http.HandlerFunc)
}

type Router interface {
	Use(middleware func(http.Handler) http.Handler)
	Get(path string, h http.HandlerFunc)
	Put(path string, h http.HandlerFunc)
	Post(path string, h http.HandlerFunc)
	Patch(path string, h http.HandlerFunc)
	Delete(path string, h http.HandlerFunc)
	Handler() http.Handler
	ProtectedGroup(path string, handleFn func(r RouteGroup))
}
