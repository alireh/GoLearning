package basic

import (
	"fmt"
	"math/rand"
	"time"
)

func TestSelect() {
	fmt.Println("----------------------------------------Start Select----------------------------------------")

	rand.Seed(time.Now().UnixNano())

	dataChannel := make(chan int)
	doneA := make(chan bool)
	doneB := make(chan bool)

	go producer(dataChannel)

	go consumerA(dataChannel, doneA)
	go consumerB(dataChannel, doneB)

	<-doneA
	<-doneB

	fmt.Println("Main: All consumers finished")
	fmt.Println("----------------------------------------End Select  ----------------------------------------")
}

func producer(out chan<- int) {
	defer close(out)
	for i := 0; i < 5; i++ {
		num := rand.Intn(100)
		fmt.Println("Producing:", num)
		out <- num
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	}
}

func consumerA(in <-chan int, done chan<- bool) {
	defer func() { done <- true }()
	for {
		select {
		case num, ok := <-in:
			if !ok {
				fmt.Println("Consumer A: Channel closed")
				return
			}
			fmt.Println("Consumer A:", num)
		case <-time.After(500 * time.Millisecond):
			fmt.Println("Consumer A: Timed out")
			return
		}
	}
}

func consumerB(in <-chan int, done chan<- bool) {
	defer func() { done <- true }()
	for {
		select {
		case num, ok := <-in:
			if !ok {
				fmt.Println("Consumer B: Channel closed")
				return
			}
			fmt.Println("Consumer B:", num)
		case <-time.After(800 * time.Millisecond):
			fmt.Println("Consumer B: Timed out")
			return
		}
	}
}
