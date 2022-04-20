package config

import (
	"os"
	"strings"

	"github.com/itsoeh/academy-advising-api/internal/repository"
	"github.com/itsoeh/academy-advising-api/internal/server"
	"github.com/itsoeh/academy-advising-api/internal/services"
)

const defaultPort = ":8080"

func Run() error {
	port := os.Getenv("PORT")

	if strings.TrimSpace(port) == "" {
		port = defaultPort
	}

	_, err := repository.LoadSqlConnection()
	if err != nil {
		return err
	}
	
	r := repository.NewAdvisoryStorage()
	s := services.NewAdvisoryManager(r)

	start := server.NewServer(port, s)
	
	return start.Run()
}
