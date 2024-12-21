package greetings

import (
	"fmt"
	"errors"
	"math/rand"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}

    // Return a greeting that embeds the name in a message.
    message := fmt.Sprintf(randomFormat(), name)
    return message, nil
}

func randomFormat() string {
	// slice of message formats
	formats := []string{
		"Hello %v. Welcome!",
		"All welcome the mighty %v!",
		"May the lord be with you %v",
	}

	return formats[rand.Intn(len(formats))]
}
