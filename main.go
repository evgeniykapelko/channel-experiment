package main

import (
	"fmt"
	"strings"
)

func shout(ping <-chan string, pong chan<- string) {
	for {
		s, ok := <-ping

		if !ok {
			// do something
		}

		pong <- fmt.Sprintf("%s!!!", strings.ToLower(s))
	}
}

func main() {
	ping := make(chan string)
	pong := make(chan string)

	go shout(ping, pong)

	fmt.Println("Type something and press ENTER (enter Q to quit)")

	for {
		fmt.Print("-> ")

		var input string

		_, _ = fmt.Scanln(&input)

		if input == strings.ToUpper("q") {
			break
		}

		ping <- input
		fmt.Print("\nping <- input")

		response := <-pong
		fmt.Print("\nresponse := <-pong")

		fmt.Println("\nResponse:", response)
	}

	fmt.Println("All done. Closing channel.")
	close(ping)
	close(pong)
}
