package microservicesTools

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

// Endpoints holds all Go kit endpoints for the HelloService.
type Endpoints struct {
	SayHelloEndpoint endpoint.Endpoint
}

// MakeSayHelloEndpoint creates the SayHello endpoint.
func MakeSayHelloEndpoint(svc HelloService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SayHelloRequest)
		v, err := svc.SayHello(ctx, req.Name)
		if err != nil {
			return SayHelloResponse{v, err.Error()}, nil
		}
		return SayHelloResponse{v, ""}, nil
	}
}

type SayHelloRequest struct {
	Name string
}

type SayHelloResponse struct {
	Message string `json:"message"`
	Err     string `json:"err,omitempty"`
}
