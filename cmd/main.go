package main

import (
	"log"

	"github.com/itsoeh/academy-advising-api/internal/repository"
	"github.com/itsoeh/academy-advising-api/internal/server"
	"github.com/labstack/echo/v4"
)

func main() {
	_, err := repository.LoadSqlConnection()
	if err != nil {
		log.Fatal(err)
	}

	start := echo.New()
	server.NewServer().AllEnpoints(start)

	log.Println("Starting server in localhost :8080 😀")
	start.Logger.Fatal(start.Start(":8080"))
}
