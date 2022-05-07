package main

import (
	"log"

	"github.com/itsoeh/academy-advising-api/internal/dependency"
)

func main() {

	// run the program
	if err := dependency.Run(); err !=  nil {
		log.Println(err.Error())
	}

}
