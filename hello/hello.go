package main

import (
	"fmt"
	"log"

	"github.com/dilip-kashyp/greetings"
)

func main() {
	log.SetPrefix("greetings ")
	log.SetFlags(0)

	names := []string{
		"Dilip",
		"Ram",
		"Kumar",
		"Rahul",
	}

	message, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
