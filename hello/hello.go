package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

// import "fmt"
// import "rsc.io/quote/v4"

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// A slice of names.
	names := []string{"Brody", "Rachel"}

	message, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
