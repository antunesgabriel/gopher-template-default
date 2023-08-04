package presentation

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/antunesgabriel/gopher-template-default/internal/presentation/controller"
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

	s.router.Post("/user", s.createLocalUserController.Handle)
	log.Println("✅ /user is loaded")

	return s
}

func (s *Server) Run(port string) error {
	handler := s.router.Handler()

	log.Printf("🚀 Server Starting on %s\n", os.Getenv("APP_URL"))

	return http.ListenAndServe(fmt.Sprintf(":%s", port), handler)
}
