package logging

import (
	"fmt"
	"log"
	"os"
)

func TestLog() {
	fmt.Println("----------------------------------------Start Log----------------------------------------")
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Set the output of the standard logger to the file
	log.SetOutput(file)

	// Log some messages with different severity levels
	log.Println("This is an info message")
	log.Println("This is another info message")
	log.Fatalf("This is a fatal error message")
	fmt.Println("----------------------------------------End Log  ----------------------------------------")
}
