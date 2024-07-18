package realtimeCommunication

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/olahol/melody"
)

//installation command
//go get github.com/olahol/melody

func TestMelody() {
	fmt.Println("----------------------------------------Start Melody----------------------------------------")

	// Create a new Gin router
	router := gin.Default()

	// Create a new Melody instance
	m := melody.New()

	// Handle HTTP GET requests to the /ws endpoint
	router.GET("/ws", func(c *gin.Context) {
		m.HandleRequest(c.Writer, c.Request)
	})

	// Handle new connections
	m.HandleConnect(func(s *melody.Session) {
		log.Println("New connection established")
	})

	// Handle disconnections
	m.HandleDisconnect(func(s *melody.Session) {
		log.Println("Connection closed")
	})

	// Handle incoming messages
	m.HandleMessage(func(s *melody.Session, msg []byte) {
		log.Printf("Received message: %s\n", string(msg))
		m.Broadcast(msg)
	})

	// Start the HTTP server on port 8080
	router.Run(":8080")
	fmt.Println("----------------------------------------End Melody  ----------------------------------------")
}
