package server

import (

	"github.com/itsoeh/academy-advising-api/internal/handlers"
	"github.com/itsoeh/academy-advising-api/internal/model"
	"github.com/itsoeh/academy-advising-api/internal/services"
	"github.com/labstack/echo/v4"
)

const defaultPort = ":8080"

type server struct {
	port string
	engine *echo.Echo
	services services.AdvisoryManager
}

func NewServer(port string, services services.AdvisoryManager) server {
	stream := make(chan *model.ChannelIsAccepted)
	defer close(stream)

	s :=  server{
		port: port,
		engine: echo.New(),
		services: services,
	}

	s.SetAllEndpoints(stream)

	return s
}

func (s *server) Run() error {
	return s.engine.Start(s.port)
}

func (s *server) SetAllEndpoints(stream chan *model.ChannelIsAccepted) {
	h := handlers.NewHandlers()
	
	route := s.engine.Group("/v1/itsoeh/academy-advising-api")
	route.POST("/create", h.CreateAdvisory(s.services))
	route.PUT("/update/:is_accepted/:advisory_id", h.UpdateAdvisory(s.services, stream))
	route.GET("/sse", h.Notify(stream))
}
