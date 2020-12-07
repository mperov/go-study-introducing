package main

import (
    "os"
    "fmt"
    "log"

    "greetings"
)

func Usage() {
    fmt.Println("./hello <your-name>")
}

func main() {
    log.SetPrefix("greetings: ")
    log.SetFlags(0)
    var name string = ""
    if len(os.Args) < 2 {
        Usage()
        os.Exit(1)
    } else {
        name = os.Args[1]
    }

    // Get a greeting message and print it.
    message, err := greetings.Hello(name)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(message)
}
