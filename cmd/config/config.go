package config

import (
	"os"
	"strings"

	"github.com/itsoeh/academy-advising-api/internal/repository"
	"github.com/itsoeh/academy-advising-api/internal/server"
	"github.com/itsoeh/academy-advising-api/internal/services"
)

const defaultPort = ":8080"

// Run method that is responsible for injecting dependencies
func Run() error {
	port := os.Getenv("PORT")

	if strings.TrimSpace(port) == "" {
		port = defaultPort
	}
	
	_, err := repository.LoadSqlConnection()
	if err != nil {
		return err
	}

	// inject dependencies
	DB := repository.NewDB()
	r := repository.NewAdvisoryStorage(DB)
	s := services.NewAdvisoryManager(r)
	
	// initialize the server
	start := server.NewServer(port, s)
	
	return start.Run()
}
