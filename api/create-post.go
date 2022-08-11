package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/vmickus/go-hexagonal-api/internal/port"
)

func CreatePost(postService port.PostService) echo.HandlerFunc {
	type Request struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	}

	return func(c echo.Context) error {
		log.Println("CreatePost")

		req := new(Request)
		if err := json.NewDecoder(c.Request().Body).Decode(req); err != nil {
			log.Print(err)
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}

		post, err := postService.CreatePost(req.Title, req.Description)
		if err != nil {
			log.Print(err)
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}

		return c.JSON(http.StatusCreated, post)
	}
}
