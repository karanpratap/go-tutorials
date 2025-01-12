package main

import (
    "fmt"
    "log"

    "example.com/greetings"
)

func main() {
    log.SetPrefix("greetings: ")
    log.SetFlags(0)

	names := []string{"Prince", "Royal_courtesan", "Emily"}

    // Get a greeting message and print it.
    messages, err := greetings.Hellos(names)
    if err != nil {
        log.Fatal(err)
    }

    // no error
    fmt.Println(messages)
}