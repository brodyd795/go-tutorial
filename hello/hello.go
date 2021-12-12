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
	message, err := greetings.Hello("")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
