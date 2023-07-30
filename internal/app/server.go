package app

import (
	"fmt"
	"net/http"

	"gihub.com/antunesgabriel/gopher-template-default/internal/app/module/user"
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
	s.router.Post("/user", s.userController.StoreLocal)

	return s
}

func (s *Server) Run(port string) error {
	handler := s.router.Handler()

	return http.ListenAndServe(fmt.Sprintf(":%s", port), handler)
}
