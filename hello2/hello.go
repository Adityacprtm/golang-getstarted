package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// fmt.Println("Hello, World!")
	// fmt.Println(quote.Go())

	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// A slice of names.
	names := []string{"Gladys", "Samantha", "Darrin"}
	// A single name
	// name := "Aditya Pratama"

	// Request greeting messages for the names.
	messages, err := greetings.Hellos(names)
	// Request a greeting message.
	// message, err := greetings.Hello(name)
	// If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, print the returned message
	// to the console.
	// fmt.Println(message)
	fmt.Println(messages)
	for _, message := range messages {
		fmt.Println(message)
	}
}
