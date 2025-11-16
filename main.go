package main

import (
	"fmt"
	"log"

	"github.com/dmigwi/tour_of_go/greetings"
)

func main() {
	// Sets properties of the predefined logger.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"Gladys", "Samantha", "Dawid"}

	// Requests a greetings message.
	message, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}

	// If no error print message to the console
	fmt.Println(message)
}
