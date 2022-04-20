package main

import (
	"log"

	"github.com/itsoeh/academy-advising-api/cmd/config"
)

func main() {
	// run the program
	if err := config.Run(); err !=  nil {
		log.Println(err.Error())
	}

}
