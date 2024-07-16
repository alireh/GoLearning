package basic

import (
	"fmt"
	"sync"
	"time"
)

// SafeCounter is a struct to hold a map and a mutex to safely access it concurrently
type SafeCounter struct {
	mu    sync.Mutex
	count map[string]int
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	// Lock and unlock to ensure exclusive access to the map
	c.count[key]++
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	// Lock and unlock to ensure exclusive access to the map
	return c.count[key]
}

func TestMutex() {
	fmt.Println("----------------------------------------Start Mutex----------------------------------------")
	counter := SafeCounter{count: make(map[string]int)}

	// We'll launch 10 goroutines to increment the counters concurrently
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				counter.Inc("somekey")
			}
		}()
	}

	// Let the goroutines work for a while
	time.Sleep(time.Second)

	// Get the final count
	fmt.Println("Final count:", counter.Value("somekey"))
	fmt.Println("----------------------------------------End Mutex  ----------------------------------------")
}
