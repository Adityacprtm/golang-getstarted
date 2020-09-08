package main

// import "fmt"
// import "rsc.io/quote"
import (
	"fmt"
	"rsc.io/quote"
	"example.com/greetings"
	"log"
)

func main() {

	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// Request a greeting message.
	message, err := greetings.Hello("")
	// If an error was returned, print it to the console and
	// exit the program.
        if err != nil {
		// log.Fatal(err)
		log.Println(err)
	}

	// If no error was returned, print the returned message
        // to the console.
	fmt.Println(message)


	fmt.Println(quote.Glass())
	fmt.Println(quote.Go())
	fmt.Println(quote.Hello())
	fmt.Println(quote.Opt())
}
