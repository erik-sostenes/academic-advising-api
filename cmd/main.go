package main

import (
	"log"
	"os"
	"strings"

	"github.com/itsoeh/academy-advising-api/internal/repository"
	"github.com/itsoeh/academy-advising-api/internal/server"
	"github.com/labstack/echo/v4"
)

const defaultPort = ":8080"

func main() {
	port := os.Getenv("PORT")

	if strings.TrimSpace(port) == "" {
		port = defaultPort
	}

	_, err := repository.LoadSqlConnection()
	if err != nil {
		log.Fatal(err)
	}

	start := echo.New()
	server.NewServer().SetAllEndpoints(start)

	log.Printf("Starting server in localhost %v", port)
	start.Logger.Fatal(start.Start(port))
}
