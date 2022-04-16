package server

import (
	"github.com/itsoeh/academy-advising-api/internal/handlers"
	"github.com/labstack/echo/v4"
)

type server struct {
	handlers.Advisory
	handlers.Notifier
}

func NewServer() *server {
	return &server{
		handlers.NewHandlers(),
		handlers.NewNotifier(),
	}
}

func (s *server) SetAllEndpoints(c *echo.Echo) {
	route := c.Group("/v1/itsoeh/academy-advising-api")
	route.POST("/create", s.CreateAdvisory)
	route.PUT("/update/:is_accepted/:advisory_id", s.UpdateAdvisory)
	c.GET("/sse", s.Notify)
}
