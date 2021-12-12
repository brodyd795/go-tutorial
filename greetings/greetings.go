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
	// to fail the test
	// message := fmt.Sprintf(randomFormat())
	return message, nil
}

// Hellos returns a map that associates each of the named people
// with a greeting message.
func Hellos(names []string) (map[string]string, error) {
	// A map to associate names with messages.
	messages := make(map[string]string)

	// Loop through the received slice of names, calling
	// the Hello function to get a message for each name.
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}

		// In the map, associate the retrieved message with
		// the name.
		messages[name] = message
	}

	return messages, nil
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
// docs aren't consistent about showing changed lines in tutorials
// slice initialization syntax is unexpected. Why {} instead of []?

// testing
// love that testing capability comes out of the box!
// interesting use of TestFunctionName syntax for declaring tests, and filename_test.go for test files
// was surprised to see that the debugging to "just worked". Maybe there was something in the recommended extensions that I installed? Cool anyway
// test output is a bit hard to read
// interesting to see that the regexp package had a built-in MustCompile method

// compilation was nice and fast :)
// can I direct the executable to a /build directory or something? I need to gitignore it, but I'd have to do that manually for every package...
