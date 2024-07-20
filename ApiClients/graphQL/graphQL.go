package graphQL

import (
	"fmt"
    "log"
    "net/http"

    "github.com/graph-gophers/graphql-go"
    "github.com/graph-gophers/graphql-go/relay"

)
//installation commands
// go get github.com/graph-gophers/graphql-go
// go get github.com/graph-gophers/graphql-go/relay

//for test
//in Postman call post api :
	//url:http://localhost:8090/query
	//body:{ "query": "{ hello }" }

func TestGraphqlGo() {
	fmt.Println("----------------------------------------Start GraphqlGo----------------------------------------")
	
    // Read the schema from the file
    schemaString := `
    type Query {
        hello: String!
    }`

    // Parse the schema
    schema := graphql.MustParseSchema(schemaString, &Resolver{})

    // Set up the HTTP handler
    http.Handle("/query", &relay.Handler{Schema: schema})

    // Start the server
    log.Println("GraphQL server started on http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8090", nil))
	fmt.Println("----------------------------------------End GraphqlGo  ----------------------------------------")
}

// Resolver is the struct where the resolver methods are defined
type Resolver struct{}

// Hello resolves the "hello" field from the schema
func (r *Resolver) Hello() string {
    return "Hello, world!"
}
