package main

import "fmt"

// Hello returns a greeting for the named person.
func main() {
    // Return a greeting that embeds the name in a message.
    message := fmt.Sprintf("Hi, %v. Welcome!", "Sam")
    fmt.Println(message)
}
