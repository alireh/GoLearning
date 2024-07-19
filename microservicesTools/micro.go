package microservicesTools

import (
	"context"
	"fmt"

	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
)

// Installation command
// go get -u go-micro.dev/v4
func TestMicro() {
	fmt.Println("----------------------------------------Start Micro----------------------------------------")

	// Create a new service
	service := micro.NewService(
		micro.Name("hello"),
	)

	// Initialize the service
	service.Init()

	// Register the handler
	micro.RegisterHandler(service.Server(), new(HelloService))

	// Run the service
	if err := service.Run(); err != nil {
		logger.Fatal(err)
	}
	fmt.Println("----------------------------------------End Micro  ----------------------------------------")
}

type HelloService struct{}

func (h *HelloService) Hello(ctx context.Context, req *HelloRequest, rsp *HelloResponse) error {
	rsp.Greeting = "Hello, " + req.Name
	return nil
}

type HelloRequest struct {
	Name string `json:"name"`
}

type HelloResponse struct {
	Greeting string `json:"greeting"`
}
