package deeper

import (
	"fmt"
	"time"
)

func TestChannel() {
	fmt.Println("----------------------------------------Start Channel----------------------------------------")
	ch := make(chan int)

	// Start a goroutine to send data to the channel
	go sendData(ch)

	// Receive data from the channel
	receiveData(ch)
	fmt.Println("----------------------------------------End Channel  ----------------------------------------")
}

func sendData(ch chan<- int) {
	// Send numbers 1 to 5 to the channel
	for i := 1; i <= 5; i++ {
		fmt.Println("Sending:", i)
		ch <- i                     // Send value to channel
		time.Sleep(1 * time.Second) // Simulate some work
	}
	close(ch) // Close the channel after sending all values
}

func receiveData(ch <-chan int) {
	// Receive data from the channel until it's closed
	for {
		// Attempt to receive from channel
		// If channel is closed and empty, ok will be false
		if val, ok := <-ch; ok {
			fmt.Println("Received:", val)
		} else {
			fmt.Println("Channel closed, exiting receiveData goroutine")
			return
		}
	}
}
