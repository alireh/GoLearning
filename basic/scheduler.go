package basic

import (
	"fmt"
	"runtime"
	"sync"
)

func TestScheduler() {
	fmt.Println("----------------------------------------Start Scheduler----------------------------------------")

	numJobs := 10
	numWorkers := runtime.NumCPU()

	jobs := make(chan int, numJobs)
	var wg sync.WaitGroup

	// Start workers
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker1(i, jobs, &wg)
	}

	// Enqueue jobs
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// Wait for all workers to finish
	wg.Wait()

	fmt.Println("----------------------------------------End Scheduler  ----------------------------------------")
}

func worker1(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, job)
		// Simulate some work
		for i := 0; i < 100000000; i++ {
			_ = i * 2
		}
		fmt.Printf("Worker %d finished job %d\n", id, job)
	}
}
