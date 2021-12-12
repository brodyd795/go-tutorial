package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
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

// init sets initial values for variables used in the function.
func init() {
	rand.Seed(time.Now().UnixNano())
}

// randomFormat returns one of a set of greeting messages. The returned
// message is selected at random.
func randomFormat() string {
	// A slice of message formats.
	formats := []string{
		"Hi, %v. Welcome!",
		"Hail, %v! Well met!",
		"Great to see you, %v!",
	}

	// Return a randomly selected message format by specifying
	// a random index for the slice of formats.
	return formats[rand.Intn(len(formats))]
}

// notes:
// nice compilation errors; used {} instead of (), and it directed me to the exact line of the error
// nice type errors; said what types needed to be returned as I was in the middle of adding a new error type
// takes a few seconds to adjust types; annoying
// strange that I have to return a string AND an error; why not one or the other?
// nice intellisense out of the box with "logs" module
// formatOnSave removes unused imports, even if I'm about to use them
// interesting use of capital vs. lowercase function names to indicate exports
// nice that arrays can be of dynamic size (Slices), and it's easy to declare/instantiate
// neet use of init function that runs automatically
// would be nice if the docs pointed out lines that changed in the tutorials.
