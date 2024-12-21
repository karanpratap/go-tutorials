package main

import (
    "fmt"
    "log"

    "example.com/greetings"
)

func main() {
    log.SetPrefix("greetings: ")
    log.SetFlags(0)

    // Get a greeting message and print it.
    message, err := greetings.Hello("Prince")
    if err != nil {
        log.Fatal(err)
    }

    // no error
    fmt.Println(message)
}