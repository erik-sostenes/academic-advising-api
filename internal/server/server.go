package server

import (
	"log"

	"github.com/itsoeh/academy-advising-api/internal/handlers"
	"github.com/itsoeh/academy-advising-api/internal/model"
	"github.com/itsoeh/academy-advising-api/internal/services"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const defaultPort = ":8080"

type server struct {
	port string
	engine *echo.Echo
	services services.AdvisoryManager
}

// NewServer to start the server
func NewServer(port string, services services.AdvisoryManager) server {
	stream := make(chan *model.ChannelIsAccepted)

	s :=  server{
		port: port,
		engine: echo.New(),
		services: services,
	}

	s.SetAllEndpoints(stream)

	return s
}

// Run will start running the program on the defined port
func (s *server) Run() error {
	log.Printf("Initialize server on the port %v", s.port)

	return  s.engine.Start(s.port)
}

// SetAllEndpoints contains all endpoints
func (s *server) SetAllEndpoints(stream chan *model.ChannelIsAccepted) {
	h := handlers.NewHandlers()
	
	// Add middlewares 
	s.engine.Use(middleware.Logger(), middleware.Recover(), middleware.CORS())

	route := s.engine.Group("/v1/itsoeh/academic-advising-api")

	route.POST("/create", h.CreateAdvisory(s.services))
	route.PUT("/update/:is_accepted/:advisory_id", h.UpdateAdvisory(s.services, stream))
	route.GET("/handshake", h.Notify(stream))
}
