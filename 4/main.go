package main

import (
	"fmt"
	"math/rand"
)

func getRandomGreeting() string {

	greetings := []string{
		"Hello %v, Nice to meet you",
		"Whatagwan %v, It's a nice day to smoke weed",
		"Hi there, Nice to meet you %v",
	}

	return greetings[rand.Intn(len(greetings))]

}

func main() {

	greeting := getRandomGreeting()

	message := fmt.Sprintf(greeting, "John")

	fmt.Println(message)

}
