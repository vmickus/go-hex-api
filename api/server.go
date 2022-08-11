package api

import (
	"encoding/json"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/vmickus/go-hexagonal-api/internal/service"
)

type Server struct {
	Router      *echo.Echo
	PostService *service.PostService
}

func NewServer() *Server {
	e := echo.New()
	return &Server{
		Router: e,
	}
}

func (s *Server) WithPostService(svc *service.PostService) *Server {
	s.PostService = svc
	return s
}

func (s *Server) LogRoutes() error {
	data, err := json.MarshalIndent(s.Router.Routes(), "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(data))

	return nil
}
