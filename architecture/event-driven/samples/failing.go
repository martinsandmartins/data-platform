package main

import (
	"fmt"
	"time"
)

func main() {
	// Create a channel to receive events
	eventCh := make(chan string)

	// Start a goroutine to handle events
	go func() {
		for {
			// Wait for an event to be received on the channel
			event := <-eventCh

			// Simulate event processing
			fmt.Println("Processing event:", event)
			time.Sleep(1 * time.Second)

			// Send a response event
			response := "Response to " + event
			eventCh <- response
		}
	}()

	// Send multiple events to the channel
	events := []string{"Event 1", "Event 2", "Event 3"}

	for _, event := range events {
		eventCh <- event
		fmt.Println("Sent event:", event)
	}

	// Wait for responses to the events
	for _, _ = range events {
		response := <-eventCh
		fmt.Println("Received response:", response)
	}

	// Wait for a few seconds before exiting
	time.Sleep(2 * time.Second)
}
