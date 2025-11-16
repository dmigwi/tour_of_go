package greetings

import (
	"errors"
	"fmt"
	"math/rand/v2"
)

// Hello returns greetings for the named person
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}

	message := fmt.Sprintf(randFormat(), name)
	return message, nil
}

// Hellos accepts an array of names and returns a map of greetings associated
// with each of the names person.
func Hellos(names []string) (map[string]string, error) {
	responses := make(map[string]string)

	for _, name := range names {
		response, err := Hello(name)
		if err != nil {
			return nil, err
		}

		responses[name] = response
	}
	return responses, nil
}

// randFormat returns a random greetings format
func randFormat() string {
	// A slice of the format strings supported
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	return formats[rand.IntN(len(formats))]
}
