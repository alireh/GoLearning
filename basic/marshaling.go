package basic

import (
	"encoding/json"
	"fmt"
	"log"
)

// Define a struct to hold our data
type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Code  string `json:"-"`               // //برای اینکه درjson نیاید //
	Phone string `json:"phone,omitempty"` //در صورتیکه مقدار نداشته باشد در خروجی نمی‌آید
	Email string `json:"email"`
}

func TestMarshaling() {
	fmt.Println("----------------------------------------Start Marshaling----------------------------------------")
	person := Person{
		Name:  "John Doe",
		Age:   30,
		Phone: "0912",
		Email: "johndoe@example.com",
	}

	// Marshal the Person struct to JSON
	personJSON, err := json.Marshal(person)
	if err != nil {
		log.Fatalf("Error marshaling JSON: %v", err)
	}

	// Print the JSON output
	fmt.Println("Marshaled JSON:", string(personJSON))

	// Example JSON data
	jsonData := `{"name":"Jane Doe","age":25,"email":"janedoe@example.com"}`

	// Unmarshal the JSON data into a Person struct
	var person2 Person
	err = json.Unmarshal([]byte(jsonData), &person2)
	if err != nil {
		log.Fatalf("Error unmarshaling JSON: %v", err)
	}

	// Print the unmarshaled struct
	fmt.Printf("Unmarshaled struct: %+v\n", person2)
	fmt.Println("----------------------------------------End Marshaling  ----------------------------------------")
}
