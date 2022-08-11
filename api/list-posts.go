package api

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func ListPosts() echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Println("ListPosts")
		return c.NoContent(http.StatusOK)
	}
}
