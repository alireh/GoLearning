package basic

import "fmt"

func TestPointer() {
	fmt.Println("----------------------------------------Start Pointer----------------------------------------")

	var a int = 10
	var p *int

	p = &a // p now holds the address of a

	fmt.Println("Value of a:", a)
	fmt.Println("Address of a:", &a)
	fmt.Println("Value of p (address of a):", p)
	fmt.Println("Value at address p:", *p) // dereferencing p to get the value of a

	// Changing the value of a using the pointer
	*p = 20
	fmt.Println("New value of a:", a)
	fmt.Println("----------------------------------------End Pointer  ----------------------------------------")
}
