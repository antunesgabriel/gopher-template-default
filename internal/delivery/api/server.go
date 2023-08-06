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
}

func NewServer(
	router Router,
	CreateLocalUserController *controller.CreateLocalUserController,
	CheckHealthController *controller.CheckHealthController,
) *Server {
	s := Server{
		router:                    router,
		createLocalUserController: CreateLocalUserController,
		checkHealthController:     CheckHealthController,
	}

	return &s
}

func (s *Server) Load() *Server {
	s.router.Get("/health", s.checkHealthController.Handle)

	log.Println("✅ /health is loaded")

	s.router.Post("/users", s.createLocalUserController.Handle)
	log.Println("✅ /users is loaded")

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
