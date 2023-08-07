package api

import (
	"fmt"
	"github.com/antunesgabriel/gopher-template-default/internal/delivery/api/controller"
	"log"
	"net/http"
	"os"
)

type Server struct {
	router                    Router
	createLocalUserController *controller.CreateLocalUserController
	checkHealthController     *controller.CheckHealthController
	authLocalController       *controller.AuthLocalController
}

func NewServer(
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
	}

	return &s
}

func (s *Server) Load() *Server {
	s.router.Get("/health", s.checkHealthController.Handle)

	log.Println("✅ /health is loaded")

	s.router.Post("/users", s.createLocalUserController.Handle)
	log.Println("✅ /users is loaded")

	s.router.Post("/auth/local", s.authLocalController.Handle)
	log.Println("✅ /auth/local is loaded")

	s.router.ProtectedGroup("/private", func(r RouteGroup) {
		// TODO: implement protected routes
		r.Get("/", func(writer http.ResponseWriter, request *http.Request) {
			writer.WriteHeader(http.StatusOK)
			writer.Write([]byte("hello world"))
		})
	})

	return s
}

func (s *Server) Run(port string) error {
	handler := s.router.Handler()

	log.Printf("🚀 Server Starting on %s\n", os.Getenv("APP_URL"))

	return http.ListenAndServe(fmt.Sprintf(":%s", port), handler)
}
