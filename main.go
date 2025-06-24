package main

import "fmt"

func main() {
	// Create a channel to communicate between goroutines
	ch := make(chan string)

	// Start a goroutine that sends a message to the channel
	go func() {
		ch <- "Hello from the goroutine!"
	}()

	// Receive the message from the channel
	message := <-ch

	// Print the received message
	fmt.Println(message)
}