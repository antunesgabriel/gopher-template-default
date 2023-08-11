package api

import (
	"fmt"
	"github.com/antunesgabriel/gopher-template-default/internal/config"
	"github.com/antunesgabriel/gopher-template-default/internal/delivery/api/controller"
	"log"
	"net/http"
)

type Server struct {
	router                    Router
	createLocalUserController *controller.CreateLocalUserController
	checkHealthController     *controller.CheckHealthController
	authLocalController       *controller.AuthLocalController
	env                       *config.Env
}

func NewServer(
	env *config.Env,
	router Router,
	createLocalUserController *controller.CreateLocalUserController,
	checkHealthController *controller.CheckHealthController,
	authLocalController *controller.AuthLocalController,

) *Server {
	s := Server{
		router:                    router,
		createLocalUserController: createLocalUserController,
		checkHealthController:     checkHealthController,
		authLocalController:       authLocalController,
		env:                       env,
	}

	return &s
}

func (it *Server) Load() *Server {
	it.router.Get("/health", it.checkHealthController.Handle)

	log.Println("✅ /health is loaded")

	it.router.Post("/signup/local", it.createLocalUserController.Handle)
	log.Println("✅ /signup/local is loaded")

	it.router.Post("/auth/local", it.authLocalController.Handle)
	log.Println("✅ /auth/local is loaded")

	it.router.ProtectedGroup("/private", func(r RouteGroup) {
		// TODO: implement protected routes
		r.Get("/", func(writer http.ResponseWriter, request *http.Request) {
			writer.WriteHeader(http.StatusOK)
			writer.Write([]byte("hello world"))
		})
	})

	return it
}

func (it *Server) Run() error {
	handler := it.router.Handler()

	log.Printf("🚀 Server Starting on %s\n", it.env.AppURL)

	return http.ListenAndServe(fmt.Sprintf(":%d", it.env.Port), handler)
}
