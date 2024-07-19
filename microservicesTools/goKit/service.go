package microservicesTools

import (
	"context"
)

// HelloService describes the service.
type HelloService interface {
	SayHello(ctx context.Context, name string) (string, error)
}

// helloService is a concrete implementation of HelloService.
type helloService struct{}

func (s helloService) SayHello(ctx context.Context, name string) (string, error) {
	return "Hello, " + name, nil
}
