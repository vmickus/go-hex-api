package api

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreatePost() echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Println("CreatePost")
		return c.NoContent(http.StatusCreated)
	}
}
