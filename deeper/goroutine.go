package deeper

import (
	"fmt"
	"time"
)

func TestGoroutine() {
	fmt.Println("----------------------------------------Start Goroutine----------------------------------------")
	go printNumbers() // start the printNumbers function as a goroutine

	// Keep the main function alive long enough to see the goroutine's output
	time.Sleep(3 * time.Second)
	fmt.Println("Main function finished")
	fmt.Println("----------------------------------------End Goroutine  ----------------------------------------")
}

func printNumbers() {
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
		time.Sleep(500 * time.Millisecond) // sleep for 500 milliseconds
	}
}
