package api

import (
	"encoding/json"
	"fmt"

	"github.com/labstack/echo/v4"
)

type Server struct {
	Router *echo.Echo
}

func NewServer() *Server {
	e := echo.New()
	return &Server{
		Router: e,
	}
}

func (s *Server) LogRoutes() error {
	data, err := json.MarshalIndent(s.Router.Routes(), "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(data))

	return nil
}
