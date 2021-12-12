package greetings

import (
	"errors"
	"fmt"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}

	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}

// notes:
// nice compilation errors; used {} instead of (), and it directed me to the exact line of the error
// nice type errors; said what types needed to be returned as I was in the middle of adding a new error type
// takes a few seconds to adjust types; annoying
// strange that I have to return a string AND an error; why not one or the other?
// nice intellisense out of the box with "logs" module