package microservicesTools

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
)

//installation command
// go get github.com/go-kit/kit

func TestGoKit() {
	fmt.Println("----------------------------------------Start GoKit----------------------------------------")

	svc := helloService{}

	sayHelloHandler := httptransport.NewServer(
		MakeSayHelloEndpoint(svc),
		decodeSayHelloRequest,
		encodeResponse,
	)

	http.Handle("/sayhello", sayHelloHandler)
	http.ListenAndServe(":8080", nil)

	fmt.Println("----------------------------------------End GoKit  ----------------------------------------")
}

func decodeSayHelloRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request SayHelloRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
