package app

import (
	"fmt"
	"net/http"

	"gihub.com/antunesgabriel/gopher-template-default/internal/app/module/user"
)

type server struct {
	router Router
}

func NewServer(router Router) *server {
	s := server{
		router: router,
	}

	return &s
}

func (s *server) Load(userController *user.UserController) *server {
	s.router.Post("/user", userController.StoreLocal)

	return s
}

func (s *server) Run(port string) error {
	handler := s.router.Handler()

	return http.ListenAndServe(fmt.Sprintf(":%s", port), handler)
}
