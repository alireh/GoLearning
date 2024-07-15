// برای مدیریت و کنترل چرخه عمر عملیات و مدیریت لغو
package basic

import (
	"context"
	"fmt"
	"time"
)

func TestContext() {
	fmt.Println("----------------------------------------Start Context----------------------------------------")
	ctx, cancel := context.WithCancel(context.Background())

	// Ensure cancel is called to release resources associated with context
	defer cancel()

	// Run a goroutine that performs some work
	go doWork(ctx)

	// Wait for a few seconds
	time.Sleep(3 * time.Second)

	// Cancel the context to signal the worker to stop
	cancel()

	// Wait a bit more to see the output
	time.Sleep(1 * time.Second)
	fmt.Println("----------------------------------------End Context  ----------------------------------------")
}

func doWork(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Work was canceled.")
			return
		default:
			// Simulate work
			fmt.Println("Doing some work...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}
