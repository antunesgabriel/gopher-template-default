package api

import (
	"fmt"
	controller2 "github.com/antunesgabriel/gopher-template-default/internal/interfaces/api/controller"
	"log"
	"net/http"
	"os"
)

type Server struct {
	router                    Router
	createLocalUserController *controller2.CreateLocalUserController
	checkHealthController     *controller2.CheckHealthController
}

func NewServer(
	router Router,
	CreateLocalUserController *controller2.CreateLocalUserController,
	CheckHealthController *controller2.CheckHealthController,
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
