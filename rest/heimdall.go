package rest

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gojek/heimdall"
	"github.com/gojek/heimdall/v7/httpclient"
)

//installation command
//go get github.com/gojek/heimdall/v7
//go get github.com/gojek/heimdall/v7/httpclient@v7.0.3

func TestHeimdall() {
	fmt.Println("----------------------------------------Start Heimdall----------------------------------------")

	// Create a new HTTP client with retries and backoff
	client := httpclient.NewClient(
		httpclient.WithHTTPTimeout(10*time.Second), // Timeout for HTTP requests
		httpclient.WithRetrier(heimdall.NewRetrier(
			heimdall.NewExponentialBackoff(100*time.Millisecond, 2*time.Second, 1.5, 2*time.Millisecond),
		)),
	)

	// Example endpoint URL
	url := "https://jsonplaceholder.typicode.com/posts/1"

	// Create an HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Perform the HTTP request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	// Print the response status and body
	fmt.Println("Response Status:", resp.Status)
	fmt.Println("Response Body:")
	if resp.Body != nil {
		// Read the response body (just printing the first 500 characters)
		buf := make([]byte, 500)
		resp.Body.Read(buf)
		fmt.Println(string(buf))
	}
	fmt.Println("----------------------------------------End Heimdall  ----------------------------------------")
}
