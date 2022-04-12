package server

import (
	"github.com/itsoeh/academic-advising-api/internal/handlers"
	"github.com/labstack/echo/v4"
)

type server struct {
	handlers.HandlersAdvisory
}

func NewServer() *server {
	return &server{
		handlers.NewHandlers(),
	}
}

func (s *server) AllEnpoints(c *echo.Echo) {
	route := c.Group("/v1/itsoeh/academy-advising-api")
	route.POST("/create", s.HandlerCreateAdvisory)
	route.PUT("/update/:is_acepted/:advisory_id", s.HandlerUpdateAdvisory)
}
