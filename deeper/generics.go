package deeper

import (
	"fmt"
)

func TestGenerics() {
	fmt.Println("----------------------------------------Start Generics----------------------------------------")

	ints := []int{1, 2, 3, 4, 5}
	floats := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
	strings := []string{"apple", "banana", "cherry"}

	fmt.Println("Max int:", Max(ints))
	fmt.Println("Max float:", Max(floats))
	fmt.Println("Max string:", Max(strings))
	fmt.Println("----------------------------------------End Generics  ----------------------------------------")
}

// Comparable is a type constraint that matches any type that supports comparison operators.
type Comparable interface {
	~int | ~float64 | ~string
}

// Max returns the maximum value from a slice of comparable elements.
func Max[T Comparable](slice []T) T {
	if len(slice) == 0 {
		var zero T
		return zero
	}
	max := slice[0]
	for _, v := range slice[1:] {
		if v > max {
			max = v
		}
	}
	return max
}
