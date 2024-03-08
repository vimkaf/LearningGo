package main

import (
	"errors"
	"fmt"
	"log"
)

func sayHello(name string) (string, error) {

	if name == "" || len(name) < 0 {
		return " ", errors.New("Missing name parameter")
	}

	message := fmt.Sprintf("Hello, %s", name)

	return message, nil

}

func main() {

	message, err := sayHello("James")

	if err != nil {
		log.Fatal("Fatal error occurred")
	}

	fmt.Println(message)

}
