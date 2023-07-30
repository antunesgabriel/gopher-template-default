package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/antunesgabriel/gopher-template-default/internal/app/module/user"
)

type Server struct {
	router         Router
	userController *user.UserController
}

func NewServer(router Router, userController *user.UserController) *Server {
	s := Server{
		router:         router,
		userController: userController,
	}

	return &s
}

func (s *Server) Load() *Server {
	// TODO: make db health check module

	log.Println("✅ /ping is loaded")

	s.router.Post("/user", s.userController.StoreLocal)
	s.router.Get("/user", s.userController.Me)
	log.Println("✅ /user is loaded")

	return s
}

func (s *Server) Run(port string) error {
	handler := s.router.Handler()

	log.Printf("🚀 Server Starting on %s\n", os.Getenv("APP_URL"))

	return http.ListenAndServe(fmt.Sprintf(":%s", port), handler)
}
