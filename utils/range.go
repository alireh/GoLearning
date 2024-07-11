package utils

import (
	"fmt"
)

func TestRange() {
	fmt.Println("----------------------------------------Start Range----------------------------------------")
	letters := []string{"a", "b", "c"}

	//With index and value

	fmt.Println("Both Index and Value of []string{\"a\", \"b\", \"c\"}")
	fmt.Println("***************************************1")

	for i, letter := range letters {

		fmt.Printf("Index: %d Value:%s\n", i, letter)

	}

	//Only value

	fmt.Println("\nOnly value")
	fmt.Println("***************************************2")

	for _, letter := range letters {
		fmt.Printf("Value: %s\n", letter)
	}

	fmt.Println("***************************************3")
	sample := map[string]string{
		"key 1": "Value 1",
		"Key 2": "Value 2",
	}

	//Iterating over all keys and values
	fmt.Println("Both Key and Value")
	for k, v := range sample {
		fmt.Printf("key :%s value: %s\n", k, v)
	}

	//Iterating over only keys
	fmt.Println("\nOnly keys")
	for k := range sample {
		fmt.Printf("key :%s\n", k)
	}

	//Iterating over only values
	fmt.Println("\nOnly values")
	for _, v := range sample {
		fmt.Printf("value :%s\n", v)
	}

	fmt.Println("***************************************4")
	str := "abz"

	//With index and value

	fmt.Println("Both Index and Value")

	for i, letter := range str {

		fmt.Printf("Start Index: %d Value:%s\n", i, string(letter))

	}

	//Only value

	fmt.Println("\nOnly value")

	for _, letter := range str {

		fmt.Printf("Value:%s\n", string(letter))

	}

	//Only index

	fmt.Println("\nOnly Index")

	for i := range str {

		fmt.Printf("Start Index: %d\n", i)

	}

	fmt.Println("----------------------------------------End Range----------------------------------------")
}
