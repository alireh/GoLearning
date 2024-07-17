package deeper

import (
	"fmt"
	"time"
)

func TestBuffer() {
	fmt.Println("----------------------------------------Start Buffer----------------------------------------")

	const numJobs = 5
	jobs := make(chan int, numJobs)    // Buffered channel for jobs
	results := make(chan int, numJobs) // Buffered channel for results

	// Start three workers
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// Send jobs to workers
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs) // Close jobs channel to indicate no more jobs will be sent

	// Collect results from workers
	for a := 1; a <= numJobs; a++ {
		<-results // Receive a result from results channel
	}

	fmt.Println("----------------------------------------End Buffer  ----------------------------------------")
}

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, job)
		time.Sleep(time.Second) // Simulate time-consuming task
		fmt.Printf("Worker %d finished job %d\n", id, job)
		results <- job * 2 // Send result back to results channel
	}
}
