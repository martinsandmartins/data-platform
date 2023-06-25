package main

import (
	"fmt"
	"time"
)

func main() {
	// Create channels for sending events and receiving responses
	eventCh := make(chan string)
	responseCh := make(chan string)

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
			responseCh <- response
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
		response := <-responseCh
		fmt.Println("Received response:", response)
	}

	// Wait for a few seconds before exiting
	time.Sleep(2 * time.Second)
}
