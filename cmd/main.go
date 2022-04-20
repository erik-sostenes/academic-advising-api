package main

import (
	"log"

	"github.com/itsoeh/academy-advising-api/cmd/config"
)

func main() {

	if err := config.Run(); err != nil {
		log.Println(err.Error())
	}

}
